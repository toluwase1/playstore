package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPIWithHandler(address string, h HandlerInterface) error {
	//Get gin's default engine
	r := gin.Default()
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)

	//post user sign in
	r.POST("/user/signin", h.SignIn)
	//post user sign out

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}
	return r.Run(address)
}

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
