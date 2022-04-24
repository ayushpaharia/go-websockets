package routes

import (
	"chat-app/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func Mux() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
