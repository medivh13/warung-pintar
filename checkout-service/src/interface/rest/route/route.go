package route

import (
	checkoutHandler "warung-pintar/checkout-service/src/interface/rest/handlers/checkout"
	"warung-pintar/checkout-service/src/interface/rest/middleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func CheckoutAppRouter(ch checkoutHandler.CheckoutHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", CheckoutRouter(ch))

	return r
}
