package logger

import (
	"context"
	"os"

	"github.com/thienkb1123/go-clean-arch/config"
	"github.com/thienkb1123/go-clean-arch/pkg/converter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface
type Logger interface {
	Debug(ctx context.Context, args ...any)
	Debugf(ctx context.Context, template string, args ...any)
	Info(ctx context.Context, args ...any)
	Infof(ctx context.Context, template string, args ...any)
	Warn(ctx context.Context, args ...any)
	Warnf(ctx context.Context, template string, args ...any)
	Error(ctx context.Context, args ...any)
	Errorf(ctx context.Context, template string, args ...any)
	DPanic(ctx context.Context, args ...any)
	DPanicf(ctx context.Context, template string, args ...any)
	Fatal(ctx context.Context, args ...any)
	Fatalf(ctx context.Context, template string, args ...any)
	GetSugaredLogger() *zap.SugaredLogger
	LoggerWithCtx(ctx context.Context, logger *zap.SugaredLogger) context.Context
	WithFields(ctx context.Context, fields Fields) context.Context
}

type Fields = map[string]any

// Logger
type apiLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) *apiLogger {
	return &apiLogger{cfg: cfg}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// loggerKey holds the context key used for loggers.
type loggerKey struct{}

// GetSugaredLogger get instance zap SugaredLogger
func (l *apiLogger) GetSugaredLogger() *zap.SugaredLogger {
	return l.sugarLogger
}

// LoggerWithCtx returns a new context derived from ctx that
// is associated with the given logger.
func (l *apiLogger) LoggerWithCtx(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// WithFields returns a new context derived from ctx
// that has a logger that always logs the given fields.
func (l *apiLogger) WithFields(ctx context.Context, fields Fields) context.Context {
	logger := l.Ctx(ctx).With(converter.MapStringToSlice(fields)...)
	return l.LoggerWithCtx(ctx, logger)
}

// Ctx returns the logger associated with the given
// context. If there is no logger, it will return Default.
func (l *apiLogger) Ctx(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		panic("nil context passed to Logger")
	}
	if logger, _ := ctx.Value(loggerKey{}).(*zap.SugaredLogger); logger != nil {
		return logger
	}

	return l.sugarLogger
}

// Logger methods
func (l *apiLogger) Debug(ctx context.Context, args ...any) {
	l.Ctx(ctx).Debug(args...)
}

func (l *apiLogger) Debugf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Debugf(template, args...)
}

func (l *apiLogger) Info(ctx context.Context, args ...any) {
	l.Ctx(ctx).Info(args...)
}

func (l *apiLogger) Infof(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Infof(template, args...)
}

func (l *apiLogger) Warn(ctx context.Context, args ...any) {
	l.Ctx(ctx).Warn(args...)
}

func (l *apiLogger) Warnf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Warnf(template, args...)
}

func (l *apiLogger) Error(ctx context.Context, args ...any) {
	l.Ctx(ctx).Error(args...)
}

func (l *apiLogger) Errorf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Errorf(template, args...)
}

func (l *apiLogger) DPanic(ctx context.Context, args ...any) {
	l.Ctx(ctx).DPanic(args...)
}

func (l *apiLogger) DPanicf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).DPanicf(template, args...)
}

func (l *apiLogger) Panic(ctx context.Context, args ...any) {
	l.Ctx(ctx).Panic(args...)
}

func (l *apiLogger) Panicf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Panicf(template, args...)
}

func (l *apiLogger) Fatal(ctx context.Context, args ...any) {
	l.Ctx(ctx).Fatal(args...)
}

func (l *apiLogger) Fatalf(ctx context.Context, template string, args ...any) {
	l.Ctx(ctx).Fatalf(template, args...)
}
