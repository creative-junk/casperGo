//Object Models
package api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Business struct {
	ID            bson.ObjectId `bson:"_id"`
	UserId        bson.ObjectId `bson:"user_id"`
	BusinessName  string        `json:"businessName"`
	BusinessEmail string        `json:"businessEmail"`
	PhoneNumber   string        `json:"phoneNumber"`
	ContactPerson string        `json:"contactPerson"`
	Address       string        `json:"address"`
	City          string        `json:"city"`
	Country       string        `json:"country"`
	Currency      string        `json:"currency"`
}
type Customer struct {
	ID           string `bson:"_id"`
	UserId       string `bson:"user_id"`
	CustomerName string `json:"customerName"`
	CompanyName  string `json:"companyName"`
	MobileNumber string `json:"mobileNumber"`
	EmailAddress string `json:"emailAddress"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PaymentTerms string `json:"paymentTerms"`
	Notes        string `json:"notes"`
}
type Estimate struct {
	ID                 string    `bson:"_id"`
	UserId             string    `bson:"user_id"`
	CustomerId         string    `json:"customer_id"`
	ProjectDescription string    `json:"project_description"`
	EstimateItems      []Item    `json:"estimate_items"`
	Taxes              []Tax     `json:"taxes"`
	CreatedAt          time.Time `json:"created_at"`
	ExpiresAt          time.Time `json:"expires_at"`
	Discount           string    `json:"discount"`
	ShippingFees       uint64    `json:"shippingFees"`
}
type Invoice struct {
	ID                 string    `bson:"_id"`
	UserId             string    `bson:"user_id"`
	CustomerId         string    `json:"customer_id"`
	ProjectDescription string    `json:"project_description"`
	InvoiceId          string    `json:"invoice_id"`
	InvoiceItems       []Item    `json:"invoice_items"`
	CreatedAt          time.Time `json:"created_at"`
	ExpiresAt          time.Time `json:"expires_at"`
	Discount           string    `json:"discount"`
	ShippingFees       uint64    `json:"shippingFees"`
	PaymentTerms       string    `json:"paymentTerms"`
}

type Sale struct {
	ID                 string    `bson:"_id"`
	UserId             string    `bson:"user_id"`
	CustomerId         string    `json:"customer_id"`
	ProjectDescription string    `json:"project_description"`
	SaleId             string    `json:"sale_id"`
	SaleItems          []Item    `json:"sale_items"`
	CreatedAt          time.Time `json:"created_at"`
	ExpiresAt          time.Time `json:"expires_at"`
	Discount           string    `json:"discount"`
	ShippingFees       uint64    `json:"shippingFees"`
}
type Payment struct {
	ID          string    `bson:"_id"`
	UserId      string    `bson:"user_id"`
	InvoiceId   string    `json:"invoice_id"`
	CreatedAt   time.Time `json:"created_at"`
	Amount      uint64    `json:"amount"`
	PaymentMode string    `json:"paymentMode"`
}
type Expense struct {
	ID          string    `bson:"_id"`
	UserId      string    `bson:"user_id"`
	ExpenseDate time.Time `json:"expenseDate"`
	ExpenseName string    `json:"expenseName"`
	PaymentMode string    `json:"paymentMode"`
	Note        string    `json:"note"`
}
type Category struct {
	ID           string `bson:"_id"`
	UserId       string `bson:"user_id"`
	CategoryName string `json:"categoryName"`
}
type Item struct {
	ID              string `bson:"_id"`
	UserId          string `bson:"user_id"`
	ItemName        string `json:"itemName"`
	ItemPrice       uint64 `json:"itemPrice"`
	ItemDescription string `json:"itemDescription"`
}
type Tax struct {
	ID      string `bson:"_id"`
	UserId  string `bson:"user_id"`
	TaxName string `json:"taxName"`
	TaxRate int    `json:"taxRate"`
}
type Notification struct {
	ID               string    `bson:"_id"`
	UserId           string    `bson:"user_id"`
	Notification     string    `json:"notification"`
	NotificationDate time.Time `json:"notification_date"`
}
type JwtToken struct {
	Token string `json:"token"`
}
type Exception struct {
	Message string `json:"message"`
}

type Invoices []Invoice
type Estimates []Estimate
type Items []Item
type Expenses []Expense
type Customers []Customer
type Payments []Payment
type Taxes []Tax
type Notifications []Notification
