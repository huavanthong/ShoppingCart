package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
