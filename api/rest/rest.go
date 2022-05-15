package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default() // the gin engine
	r.GET("/relativepath/to/url", func(c *gin.Context) {
		//take action
	})
	//get products
	r.GET("/products", func(c *gin.Context) {
		//return a list of all products to the client
	})
	//get promos
	r.GET("/promos", func(c *gin.Context) {
		//return a list of all promotions to the client
	})
	//post user sign in
	r.POST("/users/signin", func(c *gin.Context) {
		//sign in a user
	})
	//add user
	r.POST("/users", func(c *gin.Context) {
		//add a user
	})
	r.POST("/user/:id/signout", func(c *gin.Context) {
		//sign out a user with the provided id
	})

	//get user orders
	r.GET("/user/:id/orders", func(c *gin.Context) {
		//get all orders belonging to the provided user id
	})

	//post purchase charge
	r.POST("/users/charge", func(c *gin.Context) {
		//charge credit card for user
	})

	return errors.New("")
}
