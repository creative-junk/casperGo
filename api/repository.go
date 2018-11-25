//Interacts with the database
package api

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Repository struct {
}

const DB_SERVER = "mongodb://hernandez:gKZa7YS7xn444wfR@ds011820.mlab.com:11820/heroku_mg3t2zhk"
const DB_NAME = "heroku_mg3t2zhk"
const INVOICE_COLLECTION = "invoice"
const CUSTOMER_COLLECTION = "customer"
const ESTIMATE_COLLECTION = "estimate"
const EXPENSE_COLLECTION = "expense"
const SALE_COLLECTION = "sale"
const PAYMENT_COLLECTION = "payment"
const ITEM_COLLECTION = "item"
const TAX_COLLECTION = "tax"
const NOTIFICATION_COLLECTION = "notification"
const BUSINESS_COLLECTION = "business"

//Add a Business
func (r Repository) setupBusiness(business Business) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(BUSINESS_COLLECTION).Insert(business)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added a Business")
	return true
}

//Update a Business
func (r Repository) updateBusiness(business Business) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(BUSINESS_COLLECTION).UpdateId(business.ID, business)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Get All Invoices
func (r Repository) getInvoices(userId string) Invoices {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(INVOICE_COLLECTION)
	results := Invoices{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

//Add an Invoice
func (r Repository) addInvoice(invoice Invoice) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(INVOICE_COLLECTION).Insert(invoice)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added a new Invoice")
	return true
}

//Update an Invoice
func (r Repository) modifyInvoice(invoice Invoice) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(INVOICE_COLLECTION).UpdateId(invoice.ID, invoice)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Fetch an Invoice
func (r Repository) fetchInvoice(id string) Invoice {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(INVOICE_COLLECTION)
	var result Invoice

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete an Invoice
func (r Repository) deleteInvoice(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(INVOICE_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get All Customers
func (r Repository) getCustomers(userId string) Customers {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(CUSTOMER_COLLECTION)
	results := Customers{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//Add a Customer
func (r Repository) addCustomer(customer Customer) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(CUSTOMER_COLLECTION).Insert(customer)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a New Customer")
	return true
}

//Update a Customer
func (r Repository) modifyCustomer(customer Customer) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(CUSTOMER_COLLECTION).UpdateId(customer.ID, customer)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Fetch a Customer
func (r Repository) fetchCustomer(id string) Customer {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(CUSTOMER_COLLECTION)
	var result Customer

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete a Customer
func (r Repository) deleteCustomer(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(CUSTOMER_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get All Estimates
func (r Repository) getEstimates(userId string) Estimates {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Database:", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(ESTIMATE_COLLECTION)
	results := Estimates{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//Add an Estimate
func (r Repository) AddEstimate(estimate Estimate) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(ESTIMATE_COLLECTION).Insert(estimate)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a New Estimate")
	return true
}

//Update an Estimate
func (r Repository) modifyEstimate(estimate Estimate) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(ESTIMATE_COLLECTION).UpdateId(estimate.ID, estimate)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Fetch an Estimate
func (r Repository) fetchEstimate(id string) Estimate {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(ESTIMATE_COLLECTION)
	var result Estimate

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete an Estimate
func (r Repository) deleteEstimate(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(ESTIMATE_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get All Sales Receipts
func (r Repository) getSales(userId string) Sales {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(SALE_COLLECTION)
	results := Sales{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

//Add an Invoice
func (r Repository) addSale(sale Sale) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(SALE_COLLECTION).Insert(sale)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added a new Sale")
	return true
}

//Update a Sale
func (r Repository) modifySale(sale Sale) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(SALE_COLLECTION).UpdateId(sale.ID, sale)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Fetch a Sale
func (r Repository) fetchSale(id string) Sale {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(SALE_COLLECTION)
	var result Sale

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete a Sales Receipt
func (r Repository) deleteSale(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(SALE_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get All Expenses
func (r Repository) getExpenses(userId string) Expenses {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(EXPENSE_COLLECTION)
	results := Expenses{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}
	return results
}

//Add an Expense
func (r Repository) AddExpense(expense Expense) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(EXPENSE_COLLECTION).Insert(expense)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a new Expense")
	return true
}

// Modify an Expense
func (r Repository) modifyExpense(expense Expense) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(ESTIMATE_COLLECTION).UpdateId(expense.ID, expense)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Get an Expense
func (r Repository) fetchExpense(id string) Expense {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(EXPENSE_COLLECTION)
	var result Expense

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete an Expense
func (r Repository) deleteExpense(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(EXPENSE_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get Payments
func (r Repository) getPayments(userId string) Payments {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to the Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(PAYMENT_COLLECTION)
	results := Payments{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//Add a Payment
func (r Repository) AddPayment(payment Payment) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(PAYMENT_COLLECTION).Insert(payment)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a new Payment")
	return true
}

//Modify Payment
func (r Repository) modifyPayment(payment Payment) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(PAYMENT_COLLECTION).UpdateId(payment.ID, payment)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Fetch a Payment
func (r Repository) fetchPayment(id string) Payment {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(PAYMENT_COLLECTION)
	var result Payment

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete a Payment
func (r Repository) deletePayment(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(PAYMENT_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get Items
func (r Repository) getItems(userId string) Items {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a connection to the Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(ITEM_COLLECTION)
	results := Items{}
	//TODO Filter according to LoggedIn user
	iter := c.Find(bson.M{"userId": userId}).Sort("").Limit(100).Iter()
	if err := iter.All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

//Add an Item
func (r Repository) AddItem(item Item) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(ITEM_COLLECTION).Insert(item)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added an Item")
	return true
}

//Update Item
func (r Repository) modifyItem(item Item) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(ITEM_COLLECTION).UpdateId(item.ID, item)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Fetch an Item
func (r Repository) fetchItem(id string) Item {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(ITEM_COLLECTION)
	var result Item

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete an Item
func (r Repository) deleteItem(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(ITEM_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get Taxes
func (r Repository) getTaxes(userId string) Taxes {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a connection to the Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(TAX_COLLECTION)
	results := Taxes{}

	if err := c.Find(bson.M{"userId": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}
	return results
}

//Add a Tax
func (r Repository) AddTax(tax Tax) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(TAX_COLLECTION).Insert(tax)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a Tax")
	return true
}

//Modify a Tax
func (r Repository) modifyTax(tax Tax) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	err = session.DB(DB_NAME).C(TAX_COLLECTION).UpdateId(tax.ID, tax)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Fetch a Tax
func (r Repository) fetchTax(id string) Tax {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a Database Connection: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(TAX_COLLECTION)
	var result Tax

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result: ", err)
	}
	return result
}

//Delete a Tax
func (r Repository) deleteTax(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(TAX_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

//Get Notifications
func (r Repository) getNotifications(userId string) Notifications {
	session, err := mgo.Dial(DB_SERVER)

	if err != nil {
		fmt.Println("Failed to establish a connection to the Database: ", err)
	}

	defer session.Close()

	c := session.DB(DB_NAME).C(NOTIFICATION_COLLECTION)
	results := Notifications{}

	if err := c.Find(bson.M{"_id": userId}).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}
	return results
}

//Add a Notification
func (r Repository) AddNotification(notification Notification) bool {
	session, err := mgo.Dial(DB_SERVER)
	defer session.Close()

	session.DB(DB_NAME).C(NOTIFICATION_COLLECTION).Insert(notification)

	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Added a Notification")
	return true
}

//Delete a Notification
func (r Repository) deleteNotification(id string) string {
	session, err := mgo.Dial(DB_SERVER)

	defer session.Close()

	if err = session.DB(DB_NAME).C(NOTIFICATION_COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}
