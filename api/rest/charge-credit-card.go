package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/toluwase1/playstore/models"
	"log"
	"net/http"
)

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	request := struct {
		models.Order
		Remember    bool   `json:"rememberCard"`
		UseExisting bool   `json:"useExisting"`
		Token       string `json:"token"`
	}{}

	err := c.ShouldBindJSON(&request)
	log.Printf("request: %+v \n", request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}
	// Set your secret key: remember to change this to your live secret key in production
	// Keys can be obtained from: https://dashboard.stripe.com/account/apikeys
	// They key below is just for testing
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	//test cards available at:	https://stripe.com/docs/testing#cards
	//setting charge parameters

	chargeP := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(request.Price)),
		Currency:    stripe.String("usd"),
		Description: stripe.String("play-store charge..."),
	}
	stripeCustomerID := ""
	//Either remembercard or use exeisting should be enabled but not both
	if request.UseExisting {
		//use existing
		log.Println("Getting credit card id...")
		stripeCustomerID, err = h.db.GetCreditCardCID(request.CustomerID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cp := &stripe.CustomerParams{}
		cp.SetSource(request.Token)
		customer, err := customer.New(cp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stripeCustomerID = customer.ID
		if request.Remember {
			//save card!!
			err = h.db.SaveCreditCardForCustomer(request.CustomerID, stripeCustomerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}
	//we should check if the customer already ordered the same item or not but for simplicity, let's assume it's a new order
	chargeP.Customer = stripe.String(stripeCustomerID)
	_, err = charge.New(chargeP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.AddOrder(request.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
