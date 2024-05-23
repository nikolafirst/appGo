package routes

import (
	"log/slog"
	"net/http"

	"appGo/pkg/api/apiv1"

	"github.com/go-chi/chi/v5"
)

// Router has base path /api/v1
func Router(handler apiv1.ServerInterface) http.Handler {
	router := chi.NewRouter()
	router.Mount(
		"/api", apiv1.HandlerWithOptions(
			handler, apiv1.ChiServerOptions{
				BaseURL: "/v1",
				ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
					slog.Error("handle error", slog.String("err", err.Error()))
				},
			},
		),
	)
	return router
}
