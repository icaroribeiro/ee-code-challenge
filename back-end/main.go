package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/postgresdb"
    "github.com/icaroribeiro/ee-code-challenge/back-end/router"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "log"
    "net/http"
    "os"
    "os/signal"
)

var envVariablesMap map[string]string

func init() {
    var filenames []string
    var err error

    envVariablesMap = make(map[string]string)

    // The environment variable of the personal access token required to authenticate to Github.
    envVariablesMap["GITHUB_PERSONAL_ACCESS_TOKEN"] = ""

    // The environment variables related to the database settings.
    envVariablesMap["DB_USERNAME"] = ""
    envVariablesMap["DB_PASSWORD"] = ""
    envVariablesMap["DB_HOST"] = ""
    envVariablesMap["DB_PORT"] = ""
    envVariablesMap["DB_NAME"] = ""

    // The environment variables related to the HTTP server.
    envVariablesMap["HTTP_SERVER_HOST"] = ""
    envVariablesMap["HTTP_SERVER_PORT"] = ""

    // The environment files from where the variables will be loaded.
    filenames = []string{".env"}

    err = utils.GetEnvVariables(filenames, envVariablesMap)

    if err != nil {
        log.Fatal(err.Error())
    }
}

func main() {
    var token string
    var dbConfig postgresdb.DBConfig
    var s server.Server
    var err error
    var r *mux.Router
    var httpAddress string

    token = envVariablesMap["GITHUB_PERSONAL_ACCESS_TOKEN"]

    dbConfig = postgresdb.DBConfig{
        Username: envVariablesMap["DB_USERNAME"],
        Password: envVariablesMap["DB_PASSWORD"],
        Host:     envVariablesMap["DB_HOST"],
        Port:     envVariablesMap["DB_PORT"],
        Name:     envVariablesMap["DB_NAME"],
    }

    // Create the server.
    s, err = server.CreateServer(token, dbConfig)

    if err != nil {
        log.Fatal("Failed to configure the server: ", err.Error())
    }

    // Create the router by arranging the routes.
    r = router.CreateRouter(&s)

    httpAddress = fmt.Sprintf("%s:%s", envVariablesMap["HTTP_SERVER_HOST"], envVariablesMap["HTTP_SERVER_PORT"])

    log.Printf("Starting the HTTP server connection on %s", httpAddress)

    go func() {
        err = http.ListenAndServe(httpAddress, r)

        if err != nil {
            log.Fatalf("Failed to start the HTTP server connection to %s: %s", httpAddress, err.Error())
        }
    }()

    // Graceful disconnect.
    WaitForShutdown()

    err = s.Datastore.Close()

    if err != nil {
        log.Fatalf("Failed to close the database: %s", err.Error())
    }

    log.Println("Done")
}

func WaitForShutdown() {
    var interruptChan chan os.Signal

    // Create a channel to receive OS signals.
    interruptChan = make(chan os.Signal)

    // Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
    // ignoring other incoming signals.
    signal.Notify(interruptChan, os.Interrupt)

    // Block the main routine so that to keep it running until a signal is received.
    // If the main routine is shut down, the child one that is serving the server will shut down as well.
    <-interruptChan

    log.Println("Shutting down the server...")
}
