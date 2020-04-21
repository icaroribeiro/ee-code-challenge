package graphql

import (
    "net/http"
    "context"
    "golang.org/x/oauth2"
    "github.com/shurcooL/githubv4"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
)

type Query struct {
    User struct {
        StarredRepositories struct {
            TotalCount int
            Edges []StarredRepositoryEdge
        } `graphql:"starredRepositories(last:100)"`
    } `graphql:"user(login:$login)"`
}

type StarredRepositoryEdge struct {
    Node struct {
        Id string
        Name string
        Description string
        URL string
        PrimaryLanguage struct {
            Name string
        }
    }
}

func GetAllUserGithubStarredRepositories(token string, login string) ([]models.Repository, error) {
    var tokenSource oauth2.TokenSource
    var httpClient *http.Client
    var client *githubv4.Client
    var variables map[string]interface{}
    var err error
    var query Query
    var repositories []models.Repository
    var totalRepositoriesNumber int
    var repository models.Repository
    var i int

    tokenSource = oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )

    httpClient = oauth2.NewClient(context.Background(), tokenSource)

    client = githubv4.NewClient(httpClient)

    variables = map[string]interface{}{
		"login": githubv4.String(login),
    }

    err = client.Query(context.Background(), &query, variables)

    if err != nil {
        return repositories, err
    }

    // Get the last 100 starred repositories.
    totalRepositoriesNumber = query.User.StarredRepositories.TotalCount

    if totalRepositoriesNumber > 100 {
        totalRepositoriesNumber = 100
    }

    for i = 0; i < totalRepositoriesNumber; i = i + 1 {
        repository.ID = query.User.StarredRepositories.Edges[i].Node.Id
        repository.Name = query.User.StarredRepositories.Edges[i].Node.Name
        repository.Description = query.User.StarredRepositories.Edges[i].Node.Description
        repository.URL = query.User.StarredRepositories.Edges[i].Node.URL
        repository.Language = query.User.StarredRepositories.Edges[i].Node.PrimaryLanguage.Name
        repositories = append(repositories, repository)
    }

    return repositories, nil
}
