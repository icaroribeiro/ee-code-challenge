package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
)

func UpdateRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var repositoryId string
        var err error
        var repository models.Repository
        var body string
        var nRowsAffected int64

        params = mux.Vars(r)

        repositoryId = params["repositoryId"]

        if repositoryId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The repository id is required and must be set to a non-empty value in the request URL"})
            return
        }

        err = json.NewDecoder(r.Body).Decode(&repository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to decode the request body: %s", err.Error())})
            return
        }

        repository.ID = repositoryId

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

        body = fmt.Sprintf(`{"name":"%s","description":"%s","url":"%s","language":"%s"}`, 
            repository.Name, repository.Description, repository.URL, repository.Language)

        nRowsAffected, err = s.Datastore.UpdateRepository(repositoryId, repository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: %s", 
                    repositoryId, body, err.Error())})
            return
        }

        if nRowsAffected == 0 {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: " +
                    "the repository wasn't found", repositoryId, body)})
            return
        }

        if nRowsAffected != 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the repository with the id %s with %s: " + 
                    "the expected number of repositories updated: %d, got: %d", repositoryId, body, 1, nRowsAffected)})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, repository)
    })
}

func UpdateUserRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var userId string
        var repositoryId string
        var err error
        var userRepository models.UserRepository
        var body string
        var tagMap map[string]bool
        var i int
        var tag string
        var nRowsAffected int64
        var repository models.Repository

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

        err = json.NewDecoder(r.Body).Decode(&userRepository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to decode the request body: %s", err.Error())})
            return
        }

        body = fmt.Sprintf(`{"tags":[`)

        // Verify if all the tags associated with the repository are valid.
        // Additionally, checks if there are no duplicate tags.
        tagMap = make(map[string]bool)

        for i, tag = range userRepository.Tags {
            if tag == "" {
                utils.RespondWithJson(w, http.StatusNotFound, 
                    map[string]string{"error": fmt.Sprintf("Failed to add one of the tags: there is an empty value")})
                return
            }

            if !(tagMap[tag]) {
                tagMap[tag] = true
            } else {
                utils.RespondWithJson(w, http.StatusBadRequest, 
                    map[string]string{"error": fmt.Sprintf("Failed to add the tag %s: the tag is duplicated", tag)})
                return
            }

            if i == 0 {
                body += fmt.Sprintf(`"%s"`, tag)
            } else {
                body += fmt.Sprintf(`,"%s"`, tag)
            }
        }

        body += `]}`

        nRowsAffected, err = s.Datastore.UpdateUserRepository(userId, repositoryId, userRepository)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the user repository with the user id %s and repository id %s with %s: %s", 
                    userId, repositoryId, body, err.Error())})
            return
        }

        if nRowsAffected == 0 {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to update the user repository with the user id %s and repository id %s with %s: " + 
                    "the user repository wasn't found", userId, repositoryId, body)})
            return
        }

        if nRowsAffected != 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the user repository with the user id %s and repository id %s with %s: " + 
                    "the expected number of user repositories updated: %d, got: %d", userId, repositoryId, body, 1, nRowsAffected)})
            return
        }

        repository, err = s.Datastore.GetRepository(repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the repository with id %s: %s", repositoryId, err.Error())})
            return
        }

        repository.Tags = userRepository.Tags

        utils.RespondWithJson(w, http.StatusOK, repository)
    })
}
