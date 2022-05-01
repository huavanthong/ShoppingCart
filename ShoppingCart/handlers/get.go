package handlers

import (
	"net/http"

	"github.com/huavanthong/ShoppingCart/ShoppingCart/data"
)

// ListAll handles GET requests and returns all current products
func (s *ShoppingCarts) ListAll(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] get all records")

	prods := data.GetShoppingCarts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		s.l.Println("[ERROR] serializing shopping carts", err)
	}
}

// ListSingle handles GET requests
func (s *ShoppingCarts) ListSingle(rw http.ResponseWriter, r *http.Request) {

	id := getUserIdShoppingCart(r)

	s.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetUserById(id)

	switch err {
	case nil:

	case data.UserIdForShoppingCartNotFound:
		s.l.Println("[ERROR] fetching user id", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		s.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		s.l.Println("[ERROR] serializing product", err)
	}
}
