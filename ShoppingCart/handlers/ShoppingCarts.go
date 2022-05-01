package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/huavanthong/ShoppingCart/ShoppingCart/data"
)

// KeyShoppingCart is a key used for the ShoppingCart object in the context
type KeyShoppingCart struct{}

// ShoppingCart handler for getting and updating items in Shopping Cart
type ShoppingCarts struct {
	l *log.Logger
	v *data.Validation
}

// NewShoppingCarts return a new shopping cart handler with the given logger
func NewShoppingCarts(l *log.Logger, v *data.Validation) *ShoppingCarts {
	return &ShoppingCarts{l, v}
}

// ErrInvalidShoppingCartPath is an error message when the ShoppingCart path is not valid
var ErrInvalidShoppingCartPath = fmt.Errorf("Invalid Path, path should be /shoppingcart/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getShoppingCartID returns the shopping cart ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getUserIdShoppingCart(r *http.Request) int {
	// parse the shopping cart id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
