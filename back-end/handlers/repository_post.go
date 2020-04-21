package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
)

func CreateUserRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var userId string
        var err error
        var repository models.Repository
        var body string
        var userRepository models.UserRepository
        var repositoryAux models.Repository
        var nRowsAffected int64

        params = mux.Vars(r)

        userId = params["userId"]

        if userId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The user id is required and must be set to a non-empty value in the request URL"})
            return
        }

        err = json.NewDecoder(r.Body).Decode(&repository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to decode the request body: %s", err.Error())})
            return
        }

        if repository.ID == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The id field is required and must be set to a non-empty value"})
            return
        }

        if repository.Name == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The name field is required and must be set to a non-empty value"})
            return
        }

        if repository.Description == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The description field is required and must be set to a non-empty value"})
            return
        }
        
        if repository.URL == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The url field is required and must be set to a non-empty value"})
            return
        }

        if repository.Language == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The language field is required and must be set to a non-empty value"})
            return
        }

        body = fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","url":"%s","language":"%s"}`, 
            repository.ID, repository.Name, repository.Description, repository.URL, repository.Language)

        userRepository = models.UserRepository {
            UserID: userId,
            RepositoryID: repository.ID,
        }

        userRepository, err = s.Datastore.CreateUserRepository(userRepository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to create a new user repository with user id %s and repository id %s: %s", userId, repository.ID, err.Error())})
            return
        }

        // Additional step:
        // Check if the related repository already exists in the repository table.
        // If not, we must create it. Otherwise, we should update its data that may have changed over time.
        repositoryAux, err = s.Datastore.GetRepository(repository.ID)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the repository with id %s: %s", repository.ID, err.Error())})
            return
        }

        if repositoryAux.ID == "" {
            // Create the repository.
            repository, err = s.Datastore.CreateRepository(repository)

            if err != nil {
                utils.RespondWithJson(w, http.StatusInternalServerError, 
                    map[string]string{"error": fmt.Sprintf("Failed to create a new repository with %s: %s", body, err.Error())})
                return
            }
        } else {
            if !cmp.Equal(repository, repositoryAux) {
                // Update the repository.
                nRowsAffected, err = s.Datastore.UpdateRepository(repository.ID, repository)

                if err != nil {
                    utils.RespondWithJson(w, http.StatusInternalServerError, 
                        map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: %s", 
                            repository.ID, body, err.Error())})
                    return
                }

                if nRowsAffected == 0 {
                    utils.RespondWithJson(w, http.StatusConflict, 
                        map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: " +
                            "the repository wasn't found", repository.ID, body)})
                    return
                }

                if nRowsAffected != 1 {
                    utils.RespondWithJson(w, http.StatusInternalServerError, 
                        map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: " + 
                            "the expected number of repositories updated: %d, got: %d", repository.ID, body, 1, nRowsAffected)})
                    return
                }
            }
        }

        utils.RespondWithJson(w, http.StatusCreated, repository)
    })
}
