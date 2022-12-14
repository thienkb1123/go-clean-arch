package mysql

import (
	"database/sql"
	"time"

	"github.com/thienkb1123/go-clean-arch/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Return new MySQL db instance
func New(c *config.MySQLConfig) (*gorm.DB, error) {
	db, err := sql.Open("mysql", c.URI)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(c.MaxIdleConns) * time.Second)
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second)

	logMode := logger.Silent
	if c.Debug {
		logMode = logger.Error
	}

	return gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
}
