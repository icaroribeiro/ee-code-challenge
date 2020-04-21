package handlers

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/icaroribeiro/ee-code-challenge/back-end/graphql"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
)

func GetRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var repositoryId string
        var repository models.Repository
        var err error
        var userRepositories []models.UserRepository
        var tagMap map[string]bool
        var userRepository models.UserRepository
        var tags []string
        var tag string
        var tagsAux[]string

        params = mux.Vars(r)

        repositoryId = params["repositoryId"]

        if repositoryId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The repository id is required and must be set to a non-empty value in the request URL"})
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
        
        userRepositories, err = s.Datastore.GetAllUserRepositoriesByRepositoryId(repositoryId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the repository with the repository id %s: %s", repositoryId, err.Error())})
            return
        }

        tagMap = make(map[string]bool)

        for _, userRepository = range userRepositories {
            tags = userRepository.Tags

            for _, tag = range tags {
                if !(tagMap[tag]) {   
                    tagMap[tag] = true
                    tagsAux = append(tagsAux, tag)
                }
            }
        }

        if len(tagsAux) > 0 {
            repository.Tags = tagsAux
        }

        utils.RespondWithJson(w, http.StatusOK, repository)
    })
}

func GetAllUserGithubStarredRepositories(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var userId string
        var repositories []models.Repository
        var err error

        params = mux.Vars(r)

        userId = params["userId"]

        if userId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The user id is required and must be set to a non-empty value in the request URL"})
            return
        }

        repositories, err = graphql.GetAllUserGithubStarredRepositories(s.Token, userId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the Github starred repositories of the user " + 
                    "with username %s: %s", userId, err.Error())})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, repositories)
    })
}

func GetAllUserRepositories(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var params map[string]string
        var userId string
        var userRepositories []models.UserRepository
        var err error
        var userRepository models.UserRepository
        var repository models.Repository
        var repositories []models.Repository

        params = mux.Vars(r)

        userId = params["userId"]

        if userId == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The user id is required and must be set to a non-empty value in the request URL"})
            return
        }

        userRepositories, err = s.Datastore.GetAllUserRepositoriesByUserId(userId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the user repositories with user id %s: %s", userId, err.Error())})
            return
        }

        for _, userRepository = range userRepositories {
            repository, err = s.Datastore.GetRepository(userRepository.RepositoryID)

            if err != nil {
                utils.RespondWithJson(w, http.StatusInternalServerError, 
                    map[string]string{"error": fmt.Sprintf("Failed to get the repository with id %s: %s", userRepository.RepositoryID, err.Error())})
                return
            }
    
            if repository.ID == "" {
                utils.RespondWithJson(w, http.StatusNotFound, 
                    map[string]string{"error": fmt.Sprintf("Failed to get the repository with the id %s: the repository wasn't found", userRepository.RepositoryID)})
                return
            }

            repository.Tags = userRepository.Tags

            repositories = append(repositories, repository)
        }

        utils.RespondWithJson(w, http.StatusOK, repositories)
    })
}
