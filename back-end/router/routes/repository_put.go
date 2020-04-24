package routes

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/handlers"
    "github.com/icaroribeiro/ee-code-challenge/back-end/middlewares"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

func AddUpdateRepositoryRoute(s *server.Server) Route {
    var route = Route {
            Name: "UpdateRepository",
            Method: "PUT",
            Pattern: "/repositories/{repositoryId}",
            HandlerFunc: middlewares.AdaptFunc(handlers.UpdateRepository(s)).
                With(middlewares.ValidateRequestHeaderFields(map[string]string{
                        "Content-Type": "application/json",
                    }),
            ),
        }

    return route
}

func AddUpdateUserRepositoryRoute(s *server.Server) Route {
    var route = Route {
            Name: "UpdateUserRepository",
            Method: "PUT",
            Pattern: "/users/{userId}/repositories/{repositoryId}",
            HandlerFunc: middlewares.AdaptFunc(handlers.UpdateUserRepository(s)).
                With(middlewares.ValidateRequestHeaderFields(map[string]string{
                        "Content-Type": "application/json",
                    }),
            ),
        }

    return route
}
