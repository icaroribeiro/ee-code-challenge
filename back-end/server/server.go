package server

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/postgresdb"
)

// This structure is an abstraction of the server that allows to "attach" some resources in order to make them
// available during the API requests. Here, it's used to store the Github personal access token as well as
// other structure that holds attributes to manage the data.
type Server struct {
    Token string
    Datastore postgresdb.Datastore
}

func CreateServer(token string, dbConfig postgresdb.DBConfig) (Server, error) {
    var s Server
    var err error

    s.Token = token

    // Initialize the database.
    s.Datastore, err = postgresdb.InitializeDB(dbConfig)

    if err != nil {
        return s, err
    }

    return s, nil
}
