package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/thienkb1123/go-clean-arch/internal/models"
	"github.com/thienkb1123/go-clean-arch/pkg/errors"
	"github.com/thienkb1123/go-clean-arch/pkg/logger"
	"github.com/thienkb1123/go-clean-arch/pkg/sanitize"
)

// Get request id from gin context
func GetRequestID(c *fiber.Ctx) string {
	return c.GetRespHeader(fiber.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from echo context
func GetCtxWithReqID(c *fiber.Ctx) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// Get context  with request id
func GetRequestCtx(c *fiber.Ctx) context.Context {
	return context.WithValue(c.Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// UserCtxKey is a key used for the User object in the context
type UserCtxKey struct{}

// Get user ip address
func GetIPAddress(c *fiber.Ctx) string {
	return c.IP()
}

// Get user from context
func GetUserFromCtx(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(UserCtxKey{}).(*models.User)
	if !ok {
		return nil, errors.Unauthorized
	}

	return user, nil
}

// Error response with logging error for echo context
func LogResponseError(ctx *fiber.Ctx, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// Read request body and validate
func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}

// Read sanitize and validate request
func SanitizeRequest(ctx *gin.Context, request interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	defer ctx.Request.Body.Close()

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request.Context(), request)
}
