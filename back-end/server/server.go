package server

import (
    "github.com/icaroribeiro/ee-code-challenge/back-end/postgresdb"
)

// This structure is an abstraction of the server that allows to "attach" some resources in order to make them
// available during the API requests. Here, it's used to store other structure that holds attributes to manage the data.
type Server struct {
    Datastore postgresdb.Datastore
}

func CreateServer(dbConfig postgresdb.DBConfig) (Server, error) {
    var s Server
    var err error

    // Initialize the database.
    s.Datastore, err = postgresdb.InitializeDB(dbConfig)

    if err != nil {
        return s, err
    }

    return s, nil
}
