package rest

import "github.com/gin-gonic/gin"

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}
}
