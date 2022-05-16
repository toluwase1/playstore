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
// we ensure that our constructor connects to our db
func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("mysql", "root:toluwase@/play-store")
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
