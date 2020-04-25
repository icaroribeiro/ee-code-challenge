package handlers_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetRepository(t *testing.T) {
    var repository models.Repository
    var body string
    var err error
    var bodyBytes []byte
    var method string
    var path string
    var request *http.Request
    var response *httptest.ResponseRecorder
    var expectedCode int
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

    repository, err = s.Datastore.CreateRepository(repository)

    if err != nil {
        t.Fatalf("Failed to create a new repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    t.Logf("Repository: %s", string(bodyBytes))

    method = "GET"

    path = fmt.Sprintf("/repositories/%s", repository.ID)

    request, err = http.NewRequest(method, path, nil)

    if err != nil {
        t.Fatalf("Failed to create the request: %s", err.Error())
    }

    t.Logf("Request: method=%s and path=%s", method, path)

    response = httptest.NewRecorder()

    r.ServeHTTP(response, request)

    expectedCode = http.StatusOK

    if expectedCode != response.Code {
        t.Errorf("Test failed, response: code=%d and body=%+v", response.Code, response.Body)
        return
    }

    err = json.NewDecoder(response.Body).Decode(&repositoryAux)

    if err != nil {
        t.Fatalf("Failed to parse the JSON response body: %s", err.Error())
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

    t.Logf("Test successful, response: code=%d and body=%s", response.Code, string(bodyBytes))
}

func TestGetAllUserRepositories(t *testing.T) {
    var userRepository models.UserRepository
    var body string
    var err error
    var bodyBytes []byte
    var repository models.Repository
    var method string
    var path string
    var request *http.Request
    var response *httptest.ResponseRecorder
    var expectedCode int
    var repositories []models.Repository
    var isFound bool
    var repositoryAux models.Repository

    userRepository = models.UserRepository{
        UserID:       utils.GenerateRandomString(10),
        RepositoryID: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"user_id":"%s","repository_id:"%s"}`, userRepository.UserID, userRepository.RepositoryID)

    userRepository, err = s.Datastore.CreateUserRepository(userRepository)

    if err != nil {
        t.Fatalf("Failed to create a new user repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(userRepository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user repository %+v: %s", userRepository, err.Error())
    }

    t.Logf("User Repository: %s", string(bodyBytes))

    repository = models.Repository{
        ID: userRepository.RepositoryID,
        Name: utils.GenerateRandomString(10),
        Description: utils.GenerateRandomString(10),
        URL: utils.GenerateRandomString(10),
        Language: utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","url":"%s","language":"%s"}`, 
        repository.ID, repository.Name, repository.Description, repository.URL, repository.Language)

    repository, err = s.Datastore.CreateRepository(repository)

    if err != nil {
        t.Fatalf("Failed to create a new repository with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(repository)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the repository %+v: %s", repository, err.Error())
    }

    t.Logf("Repository: %s", string(bodyBytes))

    method = "GET"

    path = fmt.Sprintf("/users/%s/repositories", userRepository.UserID)

    request, err = http.NewRequest(method, path, nil)

    if err != nil {
        t.Fatalf("Failed to create the request: %s", err.Error())
    }

    t.Logf("Request: method=%s and path=%s", method, path)

    response = httptest.NewRecorder()

    r.ServeHTTP(response, request)

    expectedCode = http.StatusOK

    if expectedCode != response.Code {
        t.Errorf("Test failed, the expected response code: %d, got: %d", expectedCode, response.Code)
        return
    }

    err = json.NewDecoder(response.Body).Decode(&repositories)

    if err != nil {
        t.Fatalf("Failed to parse the JSON response body: %s", err.Error())
    }

    isFound = false

    for _, repositoryAux = range repositories {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(repository, repositoryAux) {
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
