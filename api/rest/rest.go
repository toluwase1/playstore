package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	//Get gin's default engine
	r := gin.Default()
	//Define a handler
	h, _ := NewHandler()
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)
	//post user sign in
	r.POST("/users/signin", h.SignIn)
	//add a user
	r.POST("/users", h.AddUser)
	//post user sign out
	r.POST("/user/:id/signout", h.SignOut)
	//get user orders
	r.GET("/user/:id/orders", h.GetOrders)
	//post purchase charge
	r.POST("/users/charge", h.Charge)
	//run the server
	return r.Run(address)
}
