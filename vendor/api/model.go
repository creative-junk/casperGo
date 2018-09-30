//Object Models
package api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Business struct {
	ID            bson.ObjectId `bson:"_id"`
	UserId        bson.ObjectId `bson:"user_id"`
	BusinessName  string        `json:"business_name"`
	BusinessEmail string        `json:"business_email"`
	PhoneNumber   string        `json:"phone_number"`
	ContactPerson string        `json:"contact_person"`
	Address       string        `json:"address"`
	City          string        `json:"city"`
	Country       string        `json:"country"`
	Currency      string        `json:"currency"`
}
type Customer struct {
	ID           string `bson:"_id"`
	UserId       string `bson:"user_id"`
	CustomerName string `json:"customer_name"`
	CompanyName  string `json:"company_name"`
	MobileNumber string `json:"mobile_number"`
	EmailAddress string `json:"email_address"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PaymentTerms string `json:"payment_terms"`
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
	ShippingFees       uint64    `json:"shipping_fees"`
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
	ShippingFees       uint64    `json:"shipping_fees"`
	PaymentTerms       string    `json:"payment_terms"`
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
	ShippingFees       uint64    `json:"shipping_fees"`
}
type Payment struct {
	ID          string    `bson:"_id"`
	UserId      string    `bson:"user_id"`
	InvoiceId   string    `json:"invoice_id"`
	CreatedAt   time.Time `json:"created_at"`
	Amount      uint64    `json:"amount"`
	PaymentMode string    `json:"payment_mode"`
}
type Expense struct {
	ID          string    `bson:"_id"`
	UserId      string    `bson:"user_id"`
	ExpenseDate time.Time `json:"expense_date"`
	ExpenseName string    `json:"expense_name"`
	PaymentMode string    `json:"payment_mode"`
	Note        string    `json:"note"`
}
type Category struct {
	ID           string `bson:"_id"`
	UserId       string `bson:"user_id"`
	CategoryName string `json:"category_name"`
}
type Item struct {
	ID              string `bson:"_id"`
	UserId          string `bson:"user_id"`
	ItemName        string `json:"item_name"`
	ItemPrice       uint64 `json:"item_price"`
	ItemDescription string `json:"item_description"`
}
type Tax struct {
	ID      string `bson:"_id"`
	UserId  string `bson:"user_id"`
	TaxName string `json:"tax_name"`
	TaxRate int    `json:"tax_rate"`
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
type Sales []Sale 
