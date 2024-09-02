package middleware

import (
	"net/http"
	"sitemate-challenge-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			utils.SendResponse(c, http.StatusInternalServerError, "Something Went Wrong", c.Errors)
		}
	}
}
