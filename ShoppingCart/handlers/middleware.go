package handlers

import (
	"context"
	"net/http"

	"github.com/huavanthong/ShoppingCart/ShoppingCart/data"
)

func (s *ShoppingCarts) MiddlewareValidateShoppingCart(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		shoppingCart := &data.ShoppingCart{}

		err := data.FromJSON(shoppingCart, r.Body)
		if err != nil {
			s.l.Println("[ERROR] deserializing shopping cart", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := s.v.Validate(shoppingCart)
		if len(errs) != 0 {
			s.l.Println("[ERROR] validating shopping cart", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyShoppingCart{}, shoppingCart)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
