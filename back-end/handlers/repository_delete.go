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
        var repository models.Repository
        var err error
        var nDeletedDocs int64
        var userRepositories []models.UserRepository

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

        repository, err = s.Datastore.GetRepository(repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the repository with id %s: %s", repositoryId, err.Error())})
            return
        }

        if repository.ID == "" {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to get the repository with the id %s: the repository wasn't found", repositoryId)})
            return
        }

        repository.Tags = userRepository.Tags

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
                    "repository id %s: the user repository wasn't found", userId, repositoryId)})
            return
        }

        if nDeletedDocs > 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the user repository with the user id %s and " + 
                    "repository id %s: the expected number of user repositories deleted: %d, got: %d", 
                    userId, repositoryId, 1, nDeletedDocs)})
            return
        }

        // Additional step:
        // Check if there is any other user associated with the same repository.
        // If not, we can delete the related repository from the repository table.        
        userRepositories, err = s.Datastore.GetAllUserRepositoriesByRepositoryId(repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the user repositories with the repository id %s: %s" + 
                    repositoryId, err.Error())})
            return
        }

        if len(userRepositories) == 0 {
            nDeletedDocs, err = s.Datastore.DeleteRepository(repositoryId)

            if err != nil {
                utils.RespondWithJson(w, http.StatusInternalServerError, 
                    map[string]string{"error": fmt.Sprintf("Failed to delete the repository with the id %s: %s", 
                        repositoryId, err.Error())})
                return
            }
    
            if nDeletedDocs == 0 {
                utils.RespondWithJson(w, http.StatusNotFound, 
                    map[string]string{"error": fmt.Sprintf("Failed to delete the repository with the id %s: " +
                        "the repository wasn't found", repositoryId)})
                return
            }
    
            if nDeletedDocs > 1 {
                utils.RespondWithJson(w, http.StatusInternalServerError, 
                    map[string]string{"error": fmt.Sprintf("Failed to delete the repository with the id %s: " + 
                        "the expected number of repositories deleted: %d, got: %d", repositoryId, 1, nDeletedDocs)})
                return
            }
        }

        utils.RespondWithJson(w, http.StatusOK, repository)
    })
}
