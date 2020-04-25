package postgresdb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "testing"
)

func TestUpdateUserRepository(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var nRowsAffected int64

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

    t.Logf("User Repository: %s", body)

    userRepository = models.UserRepository{
        UserID:       userRepository.UserID,
        RepositoryID: userRepository.RepositoryID,
        Tags:         []string{utils.GenerateRandomString(10)},
    }

    body = fmt.Sprintf(`{"tags":["%s"]}`, userRepository.Tags[0])

    t.Logf("New user repository data: %s", body)

    nRowsAffected, err = datastore.UpdateUserRepository(userRepository.UserID, userRepository.RepositoryID, userRepository)

    if err != nil {
        t.Fatalf("Failed to update the user repository with the user id %s and " + 
            "repository id %s with %s: %s", userRepository.UserID, userRepository.RepositoryID, body, err.Error())
    }

    if nRowsAffected == 0 {
        t.Errorf("Test failed, the user repository with the user id %s and repository id %s wasn't found", 
            userRepository.UserID, userRepository.RepositoryID)
        return
    }

    if nRowsAffected != 1 {
        t.Errorf("Test failed, the expected number of user repositories updated: %d, got: %d", 1, nRowsAffected)
        return
    }

    bodyBytes, err = json.Marshal(userRepository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user repository %+v: %s", userRepository, err.Error())
    }

    t.Logf("Test successful, the updated user repository: %s", string(bodyBytes))
}

func TestUpdateRepository(t *testing.T) {
    var repository models.Repository
    var body string
    var err error
    var bodyBytes []byte
    var nRowsAffected int64

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

    t.Logf("Repository: %s", body)

    repository = models.Repository{
        ID: repository.ID,
        Name: utils.GenerateRandomString(10),
        Description: utils.GenerateRandomString(10),
        URL: utils.GenerateRandomString(10),
        Language: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"name":"%s","description":"%s","url":"%s","language":"%s"}`, 
        repository.Name, repository.Description, repository.URL, repository.Language)

    t.Logf("New repository data: %s", body)

    nRowsAffected, err = datastore.UpdateRepository(repository.ID, repository)

    if err != nil {
        t.Fatalf("Failed to update the repository with the id %s with %s: %s", repository.ID, body, err.Error())
    }

    if nRowsAffected == 0 {
        t.Errorf("Test failed, the repository with the id %s wasn't found", repository.ID)
    }

    if nRowsAffected != 1 {
        t.Errorf("Test failed, the expected number of repositorys updated: %d, got: %d", 1, nRowsAffected)
        return
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    t.Logf("Test successful, the updated repository: %s", string(bodyBytes))
}
