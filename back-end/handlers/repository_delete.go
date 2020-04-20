package handlers

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
)

func DeleteUserRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var userId string
        var repositoryId string
        var userRepository models.UserRepository
        var err error
        var nDeletedDocs int64

        params = mux.Vars(r)

        userId = params["userId"]

        if userId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The user id is required and must be set to a non-empty value in the request URL"})
            return
        }

        repositoryId = params["repositoryId"]

        if repositoryId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The repository id is required and must be set to a non-empty value in the request URL"})
            return
        }

        userRepository, err = s.Datastore.GetUserRepository(userId, repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the user repository with the user id %s and " +
                    "repository id %s: %s", userId, repositoryId, err.Error())})
            return
        }

        nDeletedDocs, err = s.Datastore.DeleteUserRepository(userId, repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the user repository with the user id %s and " + 
                    "repository id %s: %s", userId, repositoryId, err.Error())})
            return
        }

        if nDeletedDocs == 0 {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the user repository with the user id %s and " + 
                    "repository id %s: the user repository wasn't found", userId, repositoryId, err.Error())})
            return
        }

        if nDeletedDocs > 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the user repository with the user id %s and " + 
                    "repository id %s: the expected number of user repositories deleted: %d, got: %d", 
                    userId, repositoryId, 1, nDeletedDocs)})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, userRepository)
    })
}
