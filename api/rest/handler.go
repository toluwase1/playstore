package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/playstore/dblayer"
)

type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

// NewHandler
//To follow good design principles,
//we need a constructor for Handler
func NewHandler() (*Handler, error) {
	//This creates a new pointer to the Handler object
	return new(Handler), nil
}