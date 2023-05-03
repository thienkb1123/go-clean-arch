package news

import (
	"github.com/gin-gonic/gin"
)

// News HTTP Handlers interface
type Handlers interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetByID(c *gin.Context)
	Delete(c *gin.Context)
	GetNews(c *gin.Context)
}
