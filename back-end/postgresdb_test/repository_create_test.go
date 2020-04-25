package postgresdb_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "testing"
)

func TestCreateUserRepository(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var userRepositoryAux models.UserRepository
    var bodyBytesAux []byte

    userRepository = models.UserRepository{
        UserID:       utils.GenerateRandomString(10),
        RepositoryID: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"user_id":"%s","repository_id:"%s"}`, userRepository.UserID, userRepository.RepositoryID)

    t.Logf("User Repository: %s", body)

    userRepositoryAux, err = datastore.CreateUserRepository(userRepository)

    if err != nil {
        t.Fatalf("Failed to create a new user repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(userRepository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user repository %+v: %s", userRepository, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !cmp.Equal(userRepository, userRepositoryAux) {
        bodyBytesAux, err = json.Marshal(userRepositoryAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned user repository %+v: %s", userRepositoryAux, err.Error())
        }

        t.Errorf("Test failed, the expected user repository returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the created user repository: %s", string(bodyBytes))
}

func TestCreateRepository(t *testing.T) {
    var repository models.Repository
    var body string
    var err error
    var bodyBytes []byte
    var repositoryAux models.Repository
    var bodyBytesAux []byte

    repository = models.Repository{
        ID: utils.GenerateRandomString(10),
        Name: utils.GenerateRandomString(10),
        Description: utils.GenerateRandomString(10),
        URL: utils.GenerateRandomString(10),
        Language: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","url":"%s","language":"%s"}`, 
        repository.ID, repository.Name, repository.Description, repository.URL, repository.Language)

    t.Logf("Repository: %s", body)

    repositoryAux, err = datastore.CreateRepository(repository)

    if err != nil {
        t.Fatalf("Failed to create a new repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !cmp.Equal(repository, repositoryAux) {
        bodyBytesAux, err = json.Marshal(repositoryAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned repository %+v: %s", repositoryAux, err.Error())
        }

        t.Errorf("Test failed, the expected repository returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the created repository: %s", string(bodyBytes))
}
