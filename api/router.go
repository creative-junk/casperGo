//Defines Routes and Endpoints
package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var controller = &Controller{repository: Repository{}}

//Define a Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//The available Routes for this API
type Routes []Route

var routes = Routes{
	//Business Routes
	Route{
		"SetupBusiness",
		"POST",
		"/v0/business/setup",
		Authenticate(controller.NewBusiness),
	},
	Route{
		"UpdateBusiness",
		"PUT",
		"/v0/business/update",
		Authenticate(controller.ModifyBusiness),
	},
	Route{
		"GetSubscription",
		"GET",
		"/v0/business/subscription",
		Authenticate(controller.ConfirmSubscription),
	},
	//Customer Routes
	Route{
		"ListCustomers",
		"GET",
		"/v0/customers",
		Authenticate(controller.ListCustomer),
	},
	Route{
		"NewCustomer",
		"POST",
		"/v0/customer/new",
		Authenticate(controller.AddCustomer),
	},
	Route{
		"GetCustomer",
		"GET",
		"/v0/customer/{id}",
		Authenticate(controller.FetchCustomer),
	},
	Route{
		"UpdateCustomer",
		"PUT",
		"/v0/customer/update/{id}",
		Authenticate(controller.ModifyCustomer),
	},
	Route{
		"DeleteCustomer",
		"DELETE",
		"/v0/customer/delete/{id}",
		Authenticate(controller.DeleteCustomer),
	},
	//Estimate Routes
	Route{
		"ListEstimates",
		"GET",
		"/v0/estimate",
		Authenticate(controller.ListEstimate),
	},
	Route{
		"NewEstimate",
		"POST",
		"/v0/estimate/new",
		Authenticate(controller.AddEstimate),
	},
	Route{
		"GetEstimate",
		"GET",
		"/v0/estimate/{id}",
		Authenticate(controller.FetchEstimate),
	},
	Route{
		"UpdateEstimate",
		"PUT",
		"/v0/estimate/update/{id}",
		Authenticate(controller.ModifyEstimate),
	},
	Route{
		"DeleteEstimate",
		"DELETE",
		"/v0/estimate/delete/{id}",
		Authenticate(controller.DeleteEstimate),
	},
	Route{
		"EmailEstimate",
		"GET",
		"/v0/estimate/email/{id}",
		Authenticate(controller.SendEstimate),
	},
	//Invoice Routes
	Route{
		"ListInvoices",
		"GET",
		"/v0/invoices",
		Authenticate(controller.ListInvoice),
	},
	Route{
		"NewInvoice",
		"POST",
		"/v0/invoice/new",
		Authenticate(controller.AddInvoice),
	},
	Route{
		"GetInvoice",
		"GET",
		"/v0/invoice/{id}",
		Authenticate(controller.FetchInvoice),
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/v0/invoice/update/{id}",
		Authenticate(controller.ModifyInvoice),
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/v0/invoice/delete/{id}",
		Authenticate(controller.DeleteInvoice),
	},
	Route{
		"EmailInvoice",
		"GET",
		"/v0/invoice/email/{id}",
		Authenticate(controller.SendInvoice),
	},
	//Sales Receipts
	Route{
		"ListSales",
		"GET",
		"/v0/sales",
		Authenticate(controller.ListSale),
	},
	Route{
		"NewSale",
		"POST",
		"/v0/sale/new",
		Authenticate(controller.AddSale),
	},
	Route{
		"GetSale",
		"GET",
		"/v0/sale/{id}",
		Authenticate(controller.FetchSale),
	},
	Route{
		"UpdateSale",
		"PUT",
		"/v0/sale/update/{id}",
		Authenticate(controller.ModifySale),
	},
	Route{
		"DeleteSale",
		"DELETE",
		"/v0/sale/delete/{id}",
		Authenticate(controller.DeleteSale),
	},
	//Payment Routes
	Route{
		"ListPayments",
		"GET",
		"/v0/payments",
		Authenticate(controller.ListPayment),
	},
	Route{
		"NewPayment",
		"POST",
		"/v0/payment/new",
		Authenticate(controller.AcceptPayment),
	},
	Route{
		"GetPayment",
		"GET",
		"/v0/payment/{id}",
		Authenticate(controller.FetchPayment),
	},
	Route{
		"UpdatePayment",
		"POST",
		"/v0/payment/update/{id}",
		Authenticate(controller.ModifyPayment),
	},
	//Expense Routes
	Route{
		"ListExpenses",
		"GET",
		"/v0/expenses",
		Authenticate(controller.ListExpense),
	},
	Route{
		"NewExpense",
		"POST",
		"/v0/expense/new",
		Authenticate(controller.AddExpense),
	},
	Route{
		"GetExpense",
		"GET",
		"/v0/expense/{id}",
		Authenticate(controller.FetchExpense),
	},
	Route{
		"UpdateExpense",
		"PUT",
		"/v0/expense/update/{id}",
		Authenticate(controller.ModifyExpense),
	},
	Route{
		"DeleteExpense",
		"DELETE",
		"/v0/expense/delete/{id}",
		Authenticate(controller.DeleteExpense),
	},

	//Item Routes
	Route{
		"ListItem",
		"GET",
		"/v0/items",
		Authenticate(controller.ListItem),
	},
	Route{
		"NewItem",
		"POST",
		"/v0/item/new",
		Authenticate(controller.AddItem),
	},
	Route{
		"UpdateItem",
		"PUT",
		"/v0/item/update/{id}",
		Authenticate(controller.ModifyItem),
	},
	Route{
		"DeleteItem",
		"DELETE",
		"/v0/item/delete/{id}",
		Authenticate(controller.DeleteItem),
	},
	//Tax Routes
	Route{
		"ListTax",
		"GET",
		"/v0/taxes",
		Authenticate(controller.ListTax),
	},
	Route{
		"NewTax",
		"POST",
		"/v0/tax/new",
		Authenticate(controller.AddTax),
	},
	Route{
		"UpdateTax",
		"PUT",
		"/v0tax/update/{id}",
		Authenticate(controller.ModifyTax),
	},
	Route{
		"DeleteTax",
		"DELETE",
		"/v0tax/delete/{id}",
		Authenticate(controller.DeleteTax),
	},
}

//Configure and load the routers for the API
func NewRouter() *mux.Router {
	//Create a new Router and be lenient about it
	router := mux.NewRouter().StrictSlash(true)
	//Load it up
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
