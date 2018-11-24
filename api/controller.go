//Handler methods for our endpoints
package api

import (
	"encoding/json"
	"firebase.google.com/go"
	gContext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
var user User

const bearer  = "Bearer"

type Controller struct {
	repository Repository
}

func initializeApp() *firebase.App  {
	opt := option.WithCredentialsFile("casper-invoicing-firebase-adminsdk-i4kdh-a96801f157.json")
	app,err := firebase.NewApp(context.Background(),nil,opt)
	if err != nil{
		log.Fatalf("error initializing app: %v\n",err)
	}
	return app
}

func  getUserId(r *http.Request) string  {
	//Get the User object from the Request
	user := gContext.Get(r,"decoded").(User)
	//Extract the ID and return it
	return user.ID
}
func Authenticate(next http.HandlerFunc) http.HandlerFunc  {

	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		//Get Token from the Request headers
			authorizationHeader := r.Header.Get("Authorization")

			if authorizationHeader !="" {

				bearerToken := strings.Split(authorizationHeader,"")
				if len(bearerToken)==2 {
					
					//Initialize SDK
					app := initializeApp()
					//Verify token
					client, err := app.Auth(context.Background())
					if err != nil {
						log.Fatalf("error getting Auth client: %v\n", err)

					}
					token, err := client.VerifyIDToken(r.Context(), bearerToken[1])
					if err != nil {
						json.NewEncoder(w).Encode(Exception{Message: "Invalid Authentication Token"})
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Invalid Token"))
					}
					userId := token.UID
					user.ID = userId
					gContext.Set(r, "decoded", user)
					next(w, r)

				}else{
					json.NewEncoder(w).Encode(Exception{Message:"Authorization Failed"})
				}
			}else {
				json.NewEncoder(w).Encode(Exception{Message:"An authorization Header is required"})
			}

	})
}
//NewBusiness POST
func (c *Controller) NewBusiness(w http.ResponseWriter, r *http.Request) {
	var business Business

	business.UserId = getUserId(r)

	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Setting Up Business", err)
		w.WriteHeader(http.StatusInternalServerError) //500 Error
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Setting Up Business", err)
	}
	//Unmarshal the json into our Business Struct
	if err := json.Unmarshal(body, &business); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError) //500 Error
			return
		}
	}

	//The unMarshal works, write the object to the DB
	success := c.repository.setupBusiness(business)

	//Handle the error
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated) // 201
	return
}

// UpdateBusiness PUT
func (c *Controller) ModifyBusiness(w http.ResponseWriter, r *http.Request) {
	// Create a Struct
	var business Business

	//Read the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	//Handle Error
	if err != nil {
		log.Fatalln("Error Updating Business", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close request body and handle any errors
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Business", err)
	}
	//Unmarshal the json
	if err := json.Unmarshal(body, &business); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//The Unmarshal works, update the DB
	success := c.repository.updateBusiness(business) // updates the product in the DB
	//Handle any errors
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set up the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Invoices /invoices GET
func (c *Controller) ListInvoice(w http.ResponseWriter, r *http.Request) {
	//Get All Invoices from the DB
	invoices := c.repository.getInvoices()

	//Marshal them into JSON, ignore the error for now
	//TODO Consider handling errors that occur here
	data, _ := json.Marshal(invoices)

	//Set up the Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewInvoice POST
func (c *Controller) AddInvoice(w http.ResponseWriter, r *http.Request) {
	//Setup our Struct
	var invoice Invoice
	//Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	//Handle the errors
	if err != nil {
		log.Fatalln("Error Setting up Invoice: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the request Body and handle any errors
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Invoice")
	}
	//Unmarshal the json into our Struct and handle errors
	if err := json.Unmarshal(body, &invoice); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//Unmarshalling worked so lets save to the DB
	success := c.repository.addInvoice(invoice)
	//Handle errors
	if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
	}
	//Setup the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return

}

// UpdateInvoice PUT
func (c *Controller) ModifyInvoice(w http.ResponseWriter, r *http.Request) {
	//Setup a Struct
	var invoice Invoice

	//Read the request body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	//Handle errors
	if err != nil {
		log.Fatalln("Error Updating Invoice", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the request body and handle errors
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Invoice", err)
	}
	//Unmarshal the JSON into our Struct
	if err := json.Unmarshal(body, &invoice); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//The Unmarshal worked so lets update the database
	success := c.repository.modifyInvoice(invoice) // updates the product in the DB
	//Handle errors
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set the Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//FetchInvoice GET
func (c *Controller) FetchInvoice(w http.ResponseWriter, r *http.Request) {
	//Process the URL variables from the Request
	vars := mux.Vars(r)
	//Grab the ID
	id := vars["id"]
	// Fetch the Invoice from the DB
	invoice := c.repository.fetchInvoice(id)
	//Marshal teh invoice to JSON
	data, _ := json.Marshal(invoice)

	//Set Response Headers and return the Invoice
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//Delete an Invoice
func (c *Controller) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	//Process variables from the Request
	vars := mux.Vars(r)
	//Grab the ID
	id := vars["id"]
	// Delete the Invoice and handle errors
	if err := c.repository.deleteInvoice(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	//Set the Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Customers /customers GET
func (c *Controller) ListCustomer(w http.ResponseWriter, r *http.Request) {
	//Get All Customers
	customers := c.repository.getCustomers()
	//Marshal then into JSON
	data, _ := json.Marshal(customers)
	//Send the Invoices back within the Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewCustomer POST
func (c *Controller) AddCustomer(w http.ResponseWriter, r *http.Request) {
	//Set up a Customer Struct
	var customer Customer
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	//Handle Errors
	if err != nil {
		log.Fatalln("Error Setting Up Customer: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Customer", err)
	}

	if err := json.Unmarshal(body, &customer); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//Update this users Id
	customer.UserId=user.ID

	success := c.repository.addCustomer(customer)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return
}

//FetchCustomer GET
func (c *Controller) FetchCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	customer := c.repository.fetchCustomer(id)
	data, _ := json.Marshal(customer)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// ModifyCustomer PUT
func (c *Controller) ModifyCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Customer", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Customer", err)
	}

	if err := json.Unmarshal(body, &customer); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyCustomer(customer) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Delete a Customer
func (c *Controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteCustomer(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Estimates /estimates GET
func (c *Controller) ListEstimate(w http.ResponseWriter, r *http.Request) {
	estimates := c.repository.getEstimates()
	data, _ := json.Marshal(estimates)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewEstimate POST
func (c *Controller) AddEstimate(w http.ResponseWriter, r *http.Request) {
	var estimate Estimate
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Creating Estimate: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Estimate", err)
	}
	if err := json.Unmarshal(body, &estimate); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.AddEstimate(estimate)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

//FetchEstimate GET
func (c *Controller) FetchEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	estimate := c.repository.fetchEstimate(id)
	data, _ := json.Marshal(estimate)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdateEstimate PUT
func (c *Controller) ModifyEstimate(w http.ResponseWriter, r *http.Request) {
	var estimate Estimate
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Estimate", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Estimate", err)
	}

	if err := json.Unmarshal(body, &estimate); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyEstimate(estimate) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Delete an Estimate
func (c *Controller) DeleteEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteEstimate(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Expenses /expenses GET
func (c *Controller) ListExpense(w http.ResponseWriter, r *http.Request) {
	expenses := c.repository.getExpenses()
	data, _ := json.Marshal(expenses)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewExpense POST
func (c *Controller) AddExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Creating Expense: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Expense", err)
	}
	if err := json.Unmarshal(body, &expense); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.AddExpense(expense)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

//FetchExpense GET
func (c *Controller) FetchExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	expense := c.repository.fetchExpense(id)
	data, _ := json.Marshal(expense)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdateExpense PUT
func (c *Controller) ModifyExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Expense", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Expense", err)
	}

	if err := json.Unmarshal(body, &expense); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyExpense(expense) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Delete an Expense
func (c *Controller) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteExpense(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Invoices /sales GET
func (c *Controller) ListSale(w http.ResponseWriter, r *http.Request) {
	sales := c.repository.getSales()
	data, _ := json.Marshal(sales)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewSale POST
func (c *Controller) AddSale(w http.ResponseWriter, r *http.Request) {
	var sale Sale
	//Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Saving Sale: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Sales Receipt")
	}

	if err := json.Unmarshal(body, &sale); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.addSale(sale)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

// UpdateSale PUT
func (c *Controller) ModifySale(w http.ResponseWriter, r *http.Request) {
	var sale Sale
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Sale", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Sale", err)
	}

	if err := json.Unmarshal(body, &sale); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifySale(sale) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//FetchSale GET
func (c *Controller) FetchSale(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	sale := c.repository.fetchSale(id)
	data, _ := json.Marshal(sale)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//Delete a Sales Receipt
func (c *Controller) DeleteSale(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteSale(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Items /items GET
func (c *Controller) ListItem(w http.ResponseWriter, r *http.Request) {
	items := c.repository.getItems()
	data, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewItem POST
func (c *Controller) AddItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Creating Item: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Item", err)
	}
	if err := json.Unmarshal(body, &item); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.AddItem(item)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

// UpdateItem PUT
func (c *Controller) ModifyItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Item", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Item", err)
	}

	if err := json.Unmarshal(body, &item); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyItem(item) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//FetchItem GET
func (c *Controller) FetchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	item := c.repository.fetchItem(id)
	data, _ := json.Marshal(item)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//Delete an Item
func (c *Controller) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteItem(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Payments /payments GET
func (c *Controller) ListPayment(w http.ResponseWriter, r *http.Request) {
	invoices := c.repository.getInvoices()
	data, _ := json.Marshal(invoices)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewPayment POST
func (c *Controller) AcceptPayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Creating Payment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Payment", err)
	}
	if err := json.Unmarshal(body, &payment); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.AddPayment(payment)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

//FetchPayment GET
func (c *Controller) FetchPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	payment := c.repository.fetchPayment(id)
	data, _ := json.Marshal(payment)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdatePayment PUT
func (c *Controller) ModifyPayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Payment", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Payment", err)
	}

	if err := json.Unmarshal(body, &payment); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyPayment(payment) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Delete a Payment
func (c *Controller) DeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deletePayment(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}

//Items /taxes GET
func (c *Controller) ListTax(w http.ResponseWriter, r *http.Request) {
	taxes := c.repository.getTaxes()
	data, _ := json.Marshal(taxes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//NewTax POST
func (c *Controller) AddTax(w http.ResponseWriter, r *http.Request) {
	var tax Tax
	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error Creating Item: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Creating Item", err)
	}
	if err := json.Unmarshal(body, &tax); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
		success := c.repository.AddTax(tax)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return

}

//FetchTax GET
func (c *Controller) FetchTax(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	tax := c.repository.fetchTax(id)
	data, _ := json.Marshal(tax)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdateTax PUT
func (c *Controller) ModifyTax(w http.ResponseWriter, r *http.Request) {
	var tax Tax
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error Updating Tax", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Updating Tax", err)
	}

	if err := json.Unmarshal(body, &tax); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.repository.modifyTax(tax) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Delete a Tax
func (c *Controller) DeleteTax(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.repository.deleteTax(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

}
func (c *Controller) ConfirmSubscription(w http.ResponseWriter, r *http.Request) {

}
func (c *Controller) SendEstimate(w http.ResponseWriter, r *http.Request) {

}
func (c *Controller) SendInvoice(w http.ResponseWriter, r *http.Request) {

}
