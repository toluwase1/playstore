package dblayer

import (
	"github.com/jinzhu/gorm"
	"github.com/toluwase1/playstore/models"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

//GetAllProducts raw sql query:= select * from products
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

//GetPromos select * from products where promotion IS NOT NULL
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

//GetCustomerByName select * from customers where firstname='..' and lastname='..'
func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(
		&models.Customer{
			FirstName: firstname,
			LastName:  lastname}).
		Find(&customer).Error
}

//GetCustomerByID select * from customers where id='..'
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, error error) {
	return customer, db.First(&customer, id).Error
}

//GetProduct select * from products where id='..'
func (db *DBORM) GetProduct(id int) (product models.Product, error error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	//password covered
	customer.Pass = ""
	return customer, err
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	//Obtain a *gorm.DB object representing our customer's row
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	//Retrieve the data for the customer with the passed email
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	//Compare the saved hashed password with the password provided by the user trying to sign in
	if !checkPassword(customer.Pass, pass) {
		//If failed, returns an error
		return customer, ErrINVALIDPASSWORD
	}
	//set customer pass to empty because we don't need to share this information again
	customer.Pass = ""
	//update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	//return the new customer row
	return customer, result.Find(&customer).Error
}
func (db *DBORM) SignOutUserById(id int) error {
	//Create a customer Go struct with the provided if
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//Update the customer row to reflect the fact that the customer is not logged in
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

//GetCustomerOrdersByID
//SELECT * FROM `orders` join customers on customers.id = customer_id join
//products on products.id = product_id WHERE (customer_id='1')
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, error error) {

	return orders, db.Table("orders").Select("*").
		Joins("join customers on customers.id = customer_id").
		Joins("join products on products.id = product_id").
		Where("customer_id=?", id).Scan(&orders).Error
}

/*
1. Join the two tables
	SELECT *
	FROM orders (acts as left table)
	INNER JOIN customers (acts as right table)

2. State how to Connect the two tables
	ON orders.customer_id = customers.customer_id
*/

//Add the order to the orders table
func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

//Get the id representing the credit card from the database
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	cusomterWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return cusomterWithCCID.CCID, db.First(&cusomterWithCCID, id).Error
}

//Save the credit card information for the customer
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}
