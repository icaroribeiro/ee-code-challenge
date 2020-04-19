package router

import (
	"github.com/gorilla/mux"
	"github.com/icaroribeiro/ee-code-challenge/back-end/router/routes"
	"github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

func CreateRouter(s *server.Server) *mux.Router {
	var r *mux.Router
	r = mux.NewRouter().StrictSlash(true)
	return routes.ConfigureRoutes(r, s)
}
