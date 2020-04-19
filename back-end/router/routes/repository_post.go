package routes

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/handlers"
    "github.com/icaroribeiro/ee-code-challenge/back-end/middlewares"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

func AddCreateUserRepositoryRoute(s *server.Server) Route {
    var route = Route {
            Name: "CreateUserRepository",
            Method: "POST",
            Pattern: "/users/{userId}/repository",
            HandlerFunc: middlewares.AdaptFunc(handlers.CreateUserRepository(s)).
                With(middlewares.ValidateRequestHeaderFields(map[string]string{
                        "Content-Type": "application/json",
                    }),
            ),
        }

    return route
}
