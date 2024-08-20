package errorHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrDB                  = fmt.Errorf("database error")
	ErrGenerateShortUrl    = fmt.Errorf("generate short url error")
	ErrNotFoundUrl         = fmt.Errorf("not found url")
)

// обрабатываем кастомные ошибки
func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch err.Err {
		case ErrNotFoundUrl:
			c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFoundUrl.Error()})
			return
		case ErrInternalServerError:
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError.Error()})
			return
		case ErrDB:
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrDB.Error()})
			return
		case ErrGenerateShortUrl:
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrGenerateShortUrl.Error()})
			return
		default:
			c.JSON(418, gin.H{"error": "I'm a teapot"})
			return
		}
	}
}
