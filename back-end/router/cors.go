package router

import (
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "net/http"
)

// This function enables CORS (Cross Origin Resource Sharing) so that
// an API can be accessible by JavaScript in-browser client-side code.
func LoadCORS(r *mux.Router) http.Handler {
    var corsOpts *cors.Cors
    var origins []string
    var headers []string
    var methods []string

    origins = []string{"*"}
    headers = []string{"*"}
    methods = []string{http.MethodGet,
        http.MethodPost,
        http.MethodPut,
        http.MethodPatch,
        http.MethodDelete,
        http.MethodOptions,
    }

    corsOpts = cors.New(cors.Options{
        AllowedOrigins: origins,
        AllowedHeaders: headers,
        AllowedMethods: methods,
    })

    return corsOpts.Handler(r)
}
