package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/huavanthong/ShoppingCart/ShoppingCart/data"
	"github.com/huavanthong/ShoppingCart/ShoppingCart/handlers"

	"github.com/nicholasjackson/env"
)

// declare environment server
var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")

func main() {

	env.Parse()

	// create logger only for product-api
	l := log.New(os.Stdout, "event-feed-api", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	sch := handlers.NewShoppingCarts(l, v)

	// create a new server mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/shoppingcart", sch.ListAll)
	getR.HandleFunc("/shoppingcart/{userid:[0-9]+}", sch.ListSingle)

}
