package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

// Configure the routes of the API endpoints.
func ConfigureRoutes(r *mux.Router, s *server.Server) *mux.Router {
    var routeList []Route

    // It refers to the operation linked to the service status.
    routeList = append(routeList, AddGetStatusRoute())

    // It refers to the operations linked to user repositories.
    routeList = append(routeList, AddGetAllUserGithubStarredRepositoriesRoute(s))
    routeList = append(routeList, AddCreateUserRepositoryRoute(s))
    routeList = append(routeList, AddGetAllUserRepositoriesRoute(s))
    routeList = append(routeList, AddGetAllUserRepositoriesRoute(s))
    routeList = append(routeList, AddUpdateUserRepositoryRoute(s))
    routeList = append(routeList, AddDeleteUserRepositoryRoute(s))

    for _, route := range routeList {
        r.Name(route.Name).
            Methods(route.Method).
            Path(route.Pattern).
            HandlerFunc(route.HandlerFunc)
    }

    return r
}
