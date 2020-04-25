package postgresdb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "testing"
)

func TestDeleteUserRepository(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var nRowsDeleted int64

    userRepository = models.UserRepository{
        UserID:       utils.GenerateRandomString(10),
        RepositoryID: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"user_id":"%s","repository_id:"%s"}`, userRepository.UserID, userRepository.RepositoryID)

    userRepository, err = datastore.CreateUserRepository(userRepository)

    if err != nil {
        t.Fatalf("Failed to create a new user repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(userRepository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user repository %+v: %s", userRepository, err.Error())
    }

    t.Logf("User Repository: %s", string(bodyBytes))

    nRowsDeleted, err = datastore.DeleteUserRepository(userRepository.UserID, userRepository.RepositoryID)

    if err != nil {
        t.Fatalf("Failed to delete the user repository with the user id %s and " + 
            "repository id %s: %s", userRepository.UserID, userRepository.RepositoryID, err.Error())
    }

    if nRowsDeleted == 0 {
        t.Errorf("Test failed, the user repository with the user id %s and repository id %s wasn't found", 
            userRepository.UserID, userRepository.RepositoryID)
        return
    }

    if nRowsDeleted != 1 {
        t.Errorf("Test failed, the expected number of user repositories deleted: %d, got: %d", 1, nRowsDeleted)
        return
    }

    t.Logf("Test successful, the deleted user repository: %s", string(bodyBytes))
}

func TestDeleteRepository(t *testing.T) {
    var repository models.Repository
    var body string
    var err error
    var bodyBytes []byte
    var nRowsDeleted int64

    repository = models.Repository{
        ID: utils.GenerateRandomString(10),
        Name: utils.GenerateRandomString(10),
        Description: utils.GenerateRandomString(10),
        URL: utils.GenerateRandomString(10),
        Language: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","url":"%s","language":"%s"}`, 
        repository.ID, repository.Name, repository.Description, repository.URL, repository.Language)

    repository, err = datastore.CreateRepository(repository)

    if err != nil {
        t.Fatalf("Failed to create a new repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    t.Logf("Repository: %s", string(bodyBytes))

    nRowsDeleted, err = datastore.DeleteRepository(repository.ID)

    if err != nil {
        t.Fatalf("Failed to delete the repository with the id %s: %s", repository.ID, err.Error())
    }

    if nRowsDeleted == 0 {
        t.Errorf("Test failed, the repository with the id %s wasn't found", repository.ID)
        return
    }

    if nRowsDeleted != 1 {
        t.Errorf("Test failed, the expected number of repositories deleted: %d, got: %d", 1, nRowsDeleted)
        return
    }

    t.Logf("Test successful, the deleted repository: %s", string(bodyBytes))
}
