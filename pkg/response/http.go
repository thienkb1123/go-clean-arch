package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thienkb1123/go-clean-arch/pkg/errors"
)

const (
	CodeOK = 0
)

const (
	MessageOK = "Success"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"`
}

func WithOK(c *gin.Context, data any) {
	WithCode(c, http.StatusOK, data)
}

func WithNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func WithCode(c *gin.Context, code int, data any) {
	c.JSON(code, Response{
		Code:    CodeOK,
		Message: MessageOK,
		Result:  data,
	})
}

func WithError(c *gin.Context, err error) {
	c.JSON(errors.HTTPErrorResponse(err))
}
