package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/playstore/models"
	"net/http"
)

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}
