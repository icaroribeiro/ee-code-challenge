package handlers_test

import (
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/postgresdb"
    "github.com/icaroribeiro/ee-code-challenge/back-end/router"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "log"
    "os"
    "testing"
)

var envVariablesMap map[string]string

var s server.Server

var r *mux.Router

func init() {
    var filenames []string
    var err error

    envVariablesMap = make(map[string]string)

    envVariablesMap["TEST_DB_USERNAME"] = ""
    envVariablesMap["TEST_DB_PASSWORD"] = ""
    envVariablesMap["TEST_DB_HOST"] = ""
    envVariablesMap["TEST_DB_PORT"] = ""
    envVariablesMap["TEST_DB_NAME"] = ""

    filenames = []string{"../.test.env"}

    err = utils.GetEnvVariables(filenames, envVariablesMap)

    if err != nil {
        log.Fatal(err.Error())
    }
}

// It serves as a wrapper around the testMain function that allows to defer other functions.
// At the end, it finally passes the returned exit code to os.Exit().
func TestMain(m *testing.M) {
    var exitVal int

    // Before the tests.
    utils.InitializeRandomization()

    exitVal = testMain(m)

    // After the tests.
    defer s.Datastore.Close()

    os.Exit(exitVal)
}

// It configures the settings before running the tests. It returns an integer denoting an exit code to be used 
// in the TestMain function. In the case, if the exit code is 0 it denotes success while all other codes denote failure.
func testMain(m *testing.M) int {
    var dbConfig postgresdb.DBConfig
    var err error

    dbConfig = postgresdb.DBConfig{
        Username: envVariablesMap["TEST_DB_USERNAME"],
        Password: envVariablesMap["TEST_DB_PASSWORD"],
        Host:     envVariablesMap["TEST_DB_HOST"],
        Port:     envVariablesMap["TEST_DB_PORT"],
        Name:     envVariablesMap["TEST_DB_NAME"],
    }

    s, err = server.CreateServer(dbConfig)

    if err != nil {
        log.Printf("Failed to configure the server: %s", err.Error())
        return 1
    }

    r = router.CreateRouter(&s)

    return m.Run()
}
