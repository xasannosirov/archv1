package errors

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, Error{
		Message: msg,
	})

	log.Println(msg)
}
