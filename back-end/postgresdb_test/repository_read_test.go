package postgresdb_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "testing"
)

func TestGetUserRepository(t *testing.T) {
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

    userRepositoryAux, err = datastore.CreateUserRepository(userRepository)

    if err != nil {
        t.Fatalf("Failed to create a new user repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(userRepository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user repository %+v: %s", userRepository, err.Error())
    }

    t.Logf("User Repository: %s", string(bodyBytes))

    userRepositoryAux, err = datastore.GetUserRepository(userRepository.UserID, userRepository.RepositoryID)

    if err != nil {
        t.Fatalf("Failed to get the user repository with the user id %s and " + 
            "repository id %s: %s", userRepository.UserID, userRepository.RepositoryID, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !(cmp.Equal(userRepository, userRepositoryAux)) {
        bodyBytesAux, err = json.Marshal(userRepositoryAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned user repository %+v: %s", userRepositoryAux, err.Error())
        }

        t.Errorf("Test failed, the expected user repository returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the returned user repository: %s", string(bodyBytes))
}

func TestGetAllUserRepositoriesByUserId(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var userRepositories []models.UserRepository
    var isFound bool
    var userRepositoryAux models.UserRepository

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

    userRepositories, err = datastore.GetAllUserRepositoriesByUserId(userRepository.UserID)

    if err != nil {
        t.Fatalf("Failed to get the list of all user repositories: %s", err.Error())
    }

    isFound = false

    for _, userRepositoryAux = range userRepositories {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(userRepository, userRepositoryAux) {
            isFound = true
            break
        }
    }

    if !isFound {
        t.Errorf("Test failed, the user repository not found in the list of all user repositories: %s", string(bodyBytes))
        return
    }

    t.Logf("Test successful, the user repository found in the list of all user repositories: %s", string(bodyBytes))
}

func TestGetAllUserRepositoriesByRepositoryId(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var userRepositories []models.UserRepository
    var isFound bool
    var userRepositoryAux models.UserRepository

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

    userRepositories, err = datastore.GetAllUserRepositoriesByRepositoryId(userRepository.RepositoryID)

    if err != nil {
        t.Fatalf("Failed to get the list of all user repositories: %s", err.Error())
    }

    isFound = false

    for _, userRepositoryAux = range userRepositories {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(userRepository, userRepositoryAux) {
            isFound = true
            break
        }
    }

    if !isFound {
        t.Errorf("Test failed, the user repository not found in the list of all user repositories: %s", string(bodyBytes))
        return
    }

    t.Logf("Test successful, the user repository found in the list of all user repositories: %s", string(bodyBytes))
}

func TestGetRepository(t *testing.T) {
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

    repository, err = datastore.CreateRepository(repository)

    if err != nil {
        t.Fatalf("Failed to create a new repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    t.Logf("Repository: %s", string(bodyBytes))

    repositoryAux, err = datastore.GetRepository(repository.ID)

    if err != nil {
        t.Fatalf("Failed to get the repository with the id %s: %s", repository.ID, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !(cmp.Equal(repository, repositoryAux)) {
        bodyBytesAux, err = json.Marshal(repositoryAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned repository %+v: %s", repositoryAux, err.Error())
        }

        t.Errorf("Test failed, the expected repository returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the returned repository: %s", string(bodyBytes))
}
