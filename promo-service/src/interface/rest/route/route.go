package route

import (
	itemHandler "warung-pintar/promo-service/src/interface/rest/handlers/item"
	"warung-pintar/promo-service/src/interface/rest/middleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func ItemAppRouter(ch itemHandler.ItemHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", ItemRouter(ch))

	return r
}
