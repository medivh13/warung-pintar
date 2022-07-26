package route

import (
	"net/http"

	checkoutHandler "warung-pintar/checkout-service/src/interface/rest/handlers/checkout"

	"github.com/go-chi/chi/v5"
)

func CheckoutRouter(h checkoutHandler.CheckoutHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Checkout)
	return r
}
