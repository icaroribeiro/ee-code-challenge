package routes

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/handlers"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

func AddDeleteUserRepositoryRoute(s *server.Server) Route {
    var route = Route {
            Name: "DeleteUserRepository",
            Method: "DELETE",
            Pattern: "/users/{userId}/repositories/{repositoryId}",
            HandlerFunc: handlers.DeleteUserRepository(s),
        }

    return route
}
