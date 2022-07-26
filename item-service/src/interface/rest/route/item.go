package route

import (
	"net/http"

	itemHandler "warung-pintar/item-service/src/interface/rest/handlers/item"

	"github.com/go-chi/chi/v5"
)

func ItemRouter(h itemHandler.ItemHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.GetItems)
	return r
}
