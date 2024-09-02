package ping

import (
	"net/http"
	"sitemate-challenge-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Pong(c *gin.Context) {
	utils.SendResponse(c, http.StatusOK, "pong", nil)
}
