package dblayer

import (
	"errors"
	"github.com/toluwase1/playstore/models"
)

var ErrINVALIDPASSWORD = errors.New("Invalid password")

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(id int) (product models.Product, error error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}
