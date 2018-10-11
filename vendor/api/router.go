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
		controller.NewBusiness,
	},
	Route{
		"UpdateBusiness",
		"PUT",
		"/v0/business/update",
		controller.ModifyBusiness,
	},
	Route{
		"GetSubscription",
		"GET",
		"/v0/business/subscription",
		controller.ConfirmSubscription,
	},
	//Customer Routes
	Route{
		"ListCustomers",
		"GET",
		"/v0/customers",
		controller.ListCustomer,
	},
	Route{
		"NewCustomer",
		"POST",
		"/v0/customer/new",
		controller.AddCustomer,
	},
	Route{
		"GetCustomer",
		"GET",
		"/v0/customer/{id}",
		controller.FetchCustomer,
	},
	Route{
		"UpdateCustomer",
		"PUT",
		"/v0/customer/update/{id}",
		controller.ModifyCustomer,
	},
	Route{
		"DeleteCustomer",
		"DELETE",
		"/v0/customer/delete/{id}",
		controller.DeleteCustomer,
	},
	//Estimate Routes
	Route{
		"ListEstimates",
		"GET",
		"/v0/estimate",
		controller.ListEstimate,
	},
	Route{
		"NewEstimate",
		"POST",
		"/v0/estimate/new",
		controller.AddEstimate,
	},
	Route{
		"GetEstimate",
		"GET",
		"/v0/estimate/{id}",
		controller.FetchEstimate,
	},
	Route{
		"UpdateEstimate",
		"PUT",
		"/v0/estimate/update/{id}",
		controller.ModifyEstimate,
	},
	Route{
		"DeleteEstimate",
		"DELETE",
		"/v0/estimate/delete/{id}",
		controller.DeleteEstimate,
	},
	Route{
		"EmailEstimate",
		"GET",
		"/v0/estimate/email/{id}",
		controller.SendEstimate,
	},
	//Invoice Routes
	Route{
		"ListInvoices",
		"GET",
		"/v0/invoices",
		controller.ListInvoice,
	},
	Route{
		"NewInvoice",
		"POST",
		"/v0/invoice/new",
		controller.AddInvoice,
	},
	Route{
		"GetInvoice",
		"GET",
		"/v0/invoice/{id}",
		controller.FetchInvoice,
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/v0/invoice/update/{id}",
		controller.ModifyInvoice,
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/v0/invoice/delete/{id}",
		controller.DeleteInvoice,
	},
	Route{
		"EmailInvoice",
		"GET",
		"/v0/invoice/email/{id}",
		controller.SendInvoice,
	},
	//Sales Receipts
	Route{
		"ListSales",
		"GET",
		"/v0/sales",
		controller.ListSale,
	},
	Route{
		"NewSale",
		"POST",
		"/v0/sale/new",
		controller.AddSale,
	},
	Route{
		"GetSale",
		"GET",
		"/v0/sale/{id}",
		controller.FetchSale,
	},
	Route{
		"UpdateSale",
		"PUT",
		"/v0/sale/update/{id}",
		controller.ModifySale,
	},
	Route{
		"DeleteSale",
		"DELETE",
		"/v0/sale/delete/{id}",
		controller.DeleteSale,
	},
	//Payment Routes
	Route{
		"ListPayments",
		"GET",
		"/v0/payments",
		controller.ListPayment,
	},
	Route{
		"NewPayment",
		"POST",
		"/v0/payment/new",
		controller.AcceptPayment,
	},
	Route{
		"GetPayment",
		"GET",
		"/v0/payment/{id}",
		controller.FetchPayment,
	},
	Route{
		"UpdatePayment",
		"POST",
		"/v0/payment/update/{id}",
		controller.ModifyPayment,
	},
	//Expense Routes
	Route{
		"ListExpenses",
		"GET",
		"/v0/expenses",
		controller.ListExpense,
	},
	Route{
		"NewExpense",
		"POST",
		"/v0/expense/new",
		controller.AddExpense,
	},
	Route{
		"GetExpense",
		"GET",
		"/v0/expense/{id}",
		controller.FetchExpense,
	},
	Route{
		"UpdateExpense",
		"PUT",
		"/v0/expense/update/{id}",
		controller.ModifyExpense,
	},
	Route{
		"DeleteExpense",
		"DELETE",
		"/v0/expense/delete/{id}",
		controller.DeleteExpense,
	},

	//Item Routes
	Route{
		"ListItem",
		"GET",
		"/v0/items",
		controller.ListItem,
	},
	Route{
		"NewItem",
		"POST",
		"/v0/item/new",
		controller.AddItem,
	},
	Route{
		"UpdateItem",
		"PUT",
		"/v0/item/update/{id}",
		controller.ModifyItem,
	},
	Route{
		"DeleteItem",
		"DELETE",
		"/v0/item/delete/{id}",
		controller.DeleteItem,
	},
	//Tax Routes
	Route{
		"ListTax",
		"GET",
		"/v0/taxes",
		controller.ListTax,
	},
	Route{
		"NewTax",
		"POST",
		"/v0/tax/new",
		controller.AddTax,
	},
	Route{
		"UpdateTax",
		"PUT",
		"/v0tax/update/{id}",
		controller.ModifyTax,
	},
	Route{
		"DeleteTax",
		"DELETE",
		"/v0tax/delete/{id}",
		controller.DeleteTax,
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
