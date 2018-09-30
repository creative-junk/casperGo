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
		"/business/setup",
		controller.NewBusiness,
	},
	Route{
		"UpdateBusiness",
		"PUT",
		"/business/update",
		controller.ModifyBusiness,
	},
	Route{
		"GetSubscription",
		"GET",
		"/business/subscription",
		controller.ConfirmSubscription,
	},
	//Customer Routes
	Route{
		"ListCustomers",
		"GET",
		"/customers",
		controller.ListCustomer,
	},
	Route{
		"NewCustomer",
		"POST",
		"/customer/new",
		controller.AddCustomer,
	},
	Route{
		"GetCustomer",
		"GET",
		"/customer/{id}",
		controller.FetchCustomer,
	},
	Route{
		"UpdateCustomer",
		"PUT",
		"/customer/update/{id}",
		controller.ModifyCustomer,
	},
	Route{
		"DeleteCustomer",
		"DELETE",
		"/customer/delete/{id}",
		controller.DeleteCustomer,
	},
	//Estimate Routes
	Route{
		"ListEstimates",
		"GET",
		"/estimate",
		controller.ListEstimate,
	},
	Route{
		"NewEstimate",
		"POST",
		"/estimate/new",
		controller.AddEstimate,
	},
	Route{
		"GetEstimate",
		"GET",
		"/estimate/{id}",
		controller.FetchEstimate,
	},
	Route{
		"UpdateEstimate",
		"PUT",
		"/estimate/update/{id}",
		controller.ModifyEstimate,
	},
	Route{
		"DeleteEstimate",
		"DELETE",
		"/estimate/delete/{id}",
		controller.DeleteEstimate,
	},
	Route{
		"EmailEstimate",
		"GET",
		"/estimate/email/{id}",
		controller.SendEstimate,
	},
	//Invoice Routes
	Route{
		"ListInvoices",
		"GET",
		"/invoices",
		controller.ListInvoice,
	},
	Route{
		"NewInvoice",
		"POST",
		"/invoice/new",
		controller.AddInvoice,
	},
	Route{
		"GetInvoice",
		"GET",
		"/invoice/{id}",
		controller.FetchInvoice,
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/invoice/update/{id}",
		controller.ModifyInvoice,
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/invoice/delete/{id}",
		controller.DeleteInvoice,
	},
	Route{
		"EmailInvoice",
		"GET",
		"/invoice/email/{id}",
		controller.SendInvoice,
	},
	//Sales Receipts
	Route{
		"ListSales",
		"GET",
		"/sales",
		controller.ListSale,
	},
	Route{
		"NewSale",
		"POST",
		"/sale/new",
		controller.AddSale,
	},
	Route{
		"GetSale",
		"GET",
		"/sale/{id}",
		controller.FetchSale,
	},
	Route{
		"UpdateSale",
		"PUT",
		"/sale/update/{id}",
		controller.ModifySale,
	},
	Route{
		"DeleteSale",
		"DELETE",
		"/sale/delete/{id}",
		controller.DeleteSale,
	},
	//Payment Routes
	Route{
		"ListPayments",
		"GET",
		"/payments",
		controller.ListPayment,
	},
	Route{
		"NewPayment",
		"POST",
		"/payment/new",
		controller.AcceptPayment,
	},
	Route{
		"GetPayment",
		"GET",
		"/payment/{id}",
		controller.FetchPayment,
	},
	Route{
		"UpdatePayment",
		"POST",
		"/payment/update/{id}",
		controller.ModifyPayment,
	},
	//Expense Routes
	Route{
		"ListExpenses",
		"GET",
		"/expenses",
		controller.ListExpense,
	},
	Route{
		"NewExpense",
		"POST",
		"/expense/new",
		controller.AddExpense,
	},
	Route{
		"GetExpense",
		"GET",
		"/expense/{id}",
		controller.FetchExpense,
	},
	Route{
		"UpdateExpense",
		"PUT",
		"/expense/update/{id}",
		controller.ModifyExpense,
	},
	Route{
		"DeleteExpense",
		"DELETE",
		"/expense/delete/{id}",
		controller.DeleteExpense,
	},

	//Item Routes
	Route{
		"ListItem",
		"GET",
		"/items",
		controller.ListItem,
	},
	Route{
		"NewItem",
		"POST",
		"/item/new",
		controller.AddItem,
	},
	Route{
		"UpdateItem",
		"PUT",
		"item/update/{id}",
		controller.ModifyItem,
	},
	Route{
		"DeleteItem",
		"DELETE",
		"item/delete/{id}",
		controller.DeleteItem,
	},
	//Tax Routes
	Route{
		"ListTax",
		"GET",
		"/taxes",
		controller.ListTax,
	},
	Route{
		"NewTax",
		"POST",
		"/tax/new",
		controller.AddTax,
	},
	Route{
		"UpdateTax",
		"PUT",
		"tax/update/{id}",
		controller.ModifyTax,
	},
	Route{
		"DeleteTax",
		"DELETE",
		"tax/delete/{id}",
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
