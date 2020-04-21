package routes

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/handlers"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
)

func AddGetRepositoryRoute(s *server.Server) Route {
    var route = Route {
            Name: "GetRepository",
            Method: "GET",
            Pattern: "/repositories/{repositoryId}",
            HandlerFunc: handlers.GetRepository(s),
        }

    return route
}

func AddGetAllUserGithubStarredRepositoriesRoute(s *server.Server) Route {
    var route = Route {
            Name: "GetAllUserGithubStarredRepositories",
            Method: "GET",
            Pattern: "/users/{userId}/githubStarredRepositories",
            HandlerFunc: handlers.GetAllUserGithubStarredRepositories(s),
        }

    return route
}

func AddGetAllUserRepositoriesRoute(s *server.Server) Route {
    var route = Route {
            Name: "GetAllUserRepositories",
            Method: "GET",
            Pattern: "/users/{userId}/repositories",
            HandlerFunc: handlers.GetAllUserRepositories(s),
        }

    return route
}
