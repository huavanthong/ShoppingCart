package controller

import (
	"net/http"

	"github.com/huavanthong/ShoppingCart/ShoppingCart/data"
)

// ListSingle handles GET requests
func (s *ShoppingCart) ListSingle(rw http.ResponseWriter, r *http.Request) {

	id := handler.getUserIdShoppingCart(r)

	s.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetUserById(id)

	switch err {
	case nil:

	case data.UserIdForShoppingCartNotFound:
		p.l.Println("[ERROR] fetching user id", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
