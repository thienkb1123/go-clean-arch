package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/thienkb1123/go-clean-arch/config"
	"github.com/thienkb1123/go-clean-arch/internal/models"
	"github.com/thienkb1123/go-clean-arch/pkg/errors"
	"github.com/thienkb1123/go-clean-arch/pkg/utils"
	"go.uber.org/zap"
)

// JWT way of auth using cookie or Authorization header
func (mw *MiddlewareManager) AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		ctx := c.Request.Context()
		mw.logger.Infof(ctx, "auth middleware header %s", tokenString)
		if err := mw.validateJWTToken(tokenString, c, mw.cfg); err != nil {
			mw.logger.Error(ctx, "middleware validateJWTToken", zap.String("headerJWT", err.Error()))
			c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedError(errors.Unauthorized))
			c.Abort()
		}
		c.Next()
	}
}

func (mw *MiddlewareManager) validateJWTToken(tokenString string, c *gin.Context, cfg *config.Config) error {
	if tokenString == "" {
		return errors.InvalidJWTToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(cfg.Server.JwtSecretKey)
		return secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.InvalidJWTToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(string)
		if !ok {
			return errors.InvalidJWTClaims
		}

		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return err
		}

		user := &models.User{
			UserID: userUUID,
		}

		ctx := context.WithValue(c.Request.Context(), utils.UserCtxKey{}, user)
		c.Request.WithContext(ctx)
	}
	return nil
}
