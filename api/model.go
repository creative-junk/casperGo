//Object Models
package api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Business struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId         string        `json:"user_id"`
	BusinessName   string        `json:"business_name"`
	BusinessEmail  string        `json:"business_email"`
	PhoneNumber    string        `json:"phone_number"`
	ContactPerson  string        `json:"contact_person"`
	Address        string        `json:"address"`
	City           string        `json:"city"`
	Country        string        `json:"country"`
	Currency       string        `json:"currency"`
	LastEstimateId int           `json:"estimate_id"`
	LastInvoiceId  int           `json:"invoice_id"`
	LastSaleId     int           `json:"sale_id"`
}
type Customer struct {
	ID              bson.ObjectId    `json:"id" bson:"_id,omitempty"`
	UserId          string           `json:"user_id"`
	Business        string           `json:"business_id"`
	CustomerName    string           `json:"customer_name"`
	CompanyName     string           `json:"company_name"`
	MobileNumber    string           `json:"mobile_number"`
	EmailAddress    string           `json:"email_address"`
	Address         string           `json:"address"`
	City            string           `json:"city"`
	Country         string           `json:"country"`
	PaymentTerms    string           `json:"payment_terms"`
	Notes           []Note           `json:"notes"`
	Attachments     []Attachment     `json:"attachment"`
	CreatedAt       time.Time        `json:"created_at"`
	OpenInvoices    string           `json:"open_invoices" bson:"-"`
	OverdueInvoices string           `json:"overdue_invoices" bson:"-"`
	OpenEstimates   string           `json:"open_estimates" bson:"-"`
	InvoiceSummary  []InvoiceSummary `json:"invoice_summary" bson:"-"`
}
type Estimate struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId             string        `bson:"user_id"`
	Business           string        `json:"business_id"`
	CustomerId         string        `json:"customer_id"`
	EstimateId         int           `json:"estimate_id"`
	ProjectDescription string        `json:"project_description"`
	EstimateItems      []Item        `json:"estimate_items"`
	Taxes              []Tax         `json:"taxes"`
	CreatedAt          time.Time     `json:"created_at"`
	ExpiresAt          time.Time     `json:"expires_at"`
	Discount           string        `json:"discount"`
	ShippingFees       string        `json:"shipping_fees"`
	EstimateTotal      string        `json:"estimate_total"`
	EstimateStatus     string        `json:"estimate_status"`
	Notes              []Note        `json:"notes"`
	Attachments        []Attachment  `json:"attachment"`
}
type Invoice struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId             string        `json:"user_id"`
	Business           string        `json:"business_id"`
	CustomerId         string        `json:"customer_id"`
	ProjectDescription string        `json:"project_description"`
	InvoiceId          int           `json:"invoice_id"`
	InvoiceItems       []Item        `json:"invoice_items"`
	CreatedAt          time.Time     `json:"created_at"`
	ExpiresAt          time.Time     `json:"expires_at"`
	Discount           string        `json:"discount"`
	ShippingFees       string        `json:"shipping_fees"`
	PaymentTerms       string        `json:"payment_terms"`
	InvoiceTotal       string        `json:"invoice_total"`
	InvoiceStatus      string        `json:"invoice_status"`
	Notes              []Note        `json:"notes"`
	Attachments        []Attachment  `json:"attachment"`
}

type Sale struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId             string        `json:"user_id"`
	Business           string        `json:"business_id"`
	CustomerId         string        `json:"customer_id"`
	ProjectDescription string        `json:"project_description"`
	SaleId             int           `json:"sale_id"`
	SaleItems          []Item        `json:"sale_items"`
	CreatedAt          time.Time     `json:"created_at"`
	ExpiresAt          time.Time     `json:"expires_at"`
	Discount           string        `json:"discount"`
	ShippingFees       string        `json:"shipping_fees"`
	SaleTotal          string        `json:"sale_total"`
	Notes              []Note        `json:"notes"`
	Attachments        []Attachment  `json:"attachment"`
}
type Payment struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId      string        `bson:"user_id"`
	Business    string        `json:"business_id"`
	InvoiceId   string        `json:"invoice_id"`
	CreatedAt   time.Time     `json:"created_at"`
	Amount      string        `json:"amount"`
	PaymentMode string        `json:"payment_mode"`
	Notes       []Note        `json:"notes"`
	Attachments []Attachment  `json:"attachment"`
}
type Expense struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId      string        `bson:"user_id"`
	Business    string        `json:"business"`
	ExpenseDate time.Time     `json:"expense_date"`
	ExpenseName string        `json:"expense_name"`
	PaymentMode string        `json:"payment_mode"`
	Notes       []Note        `json:"notes"`
	Attachments []Attachment  `json:"attachment"`
}
type Category struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId       string        `json:"user_id"`
	Business     string        `json:"business_id"`
	CategoryName string        `json:"category_name"`
}
type Item struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId          string        `json:"user_id"`
	Business        string        `json:"business_id"`
	ItemName        string        `json:"item_name"`
	ItemPrice       string        `json:"item_price"`
	ItemDescription string        `json:"item_description"`
}
type Tax struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId   string        `json:"user_id"`
	Business string        `json:"business_id"`
	TaxName  string        `json:"tax_name"`
	TaxRate  int           `json:"tax_rate"`
}
type Notification struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId           string        `json:"user_id"`
	Business         string        `json:"business_id"`
	Notification     string        `json:"notification"`
	NotificationDate time.Time     `json:"notification_date"`
}
type User struct {
	ID string
}
type JwtToken struct {
	Token string `json:"token"`
}
type Exception struct {
	Message string `json:"message"`
}
type InvoiceSummary struct {
	Id             string    `json:"_id"`
	InvoiceId      string    `json:"invoice_id"`
	InvoiceAmount  string    `json:"invoice_amount"`
	CreationDate   time.Time `json:"created_at"`
	ExpirationDate time.Time `json:"expiration_date"`
}
type Note struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Note     string        `json:"note"`
	UserId   string        `json:"user_id"`
	Business string        `json:"business_id"`
}
type Attachment struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Url      string        `json:"url"`
	UserId   string        `json:"user_id"`
	Business string        `json:"business_id"`
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
