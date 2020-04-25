package handlers_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestCreateUserRepository(t *testing.T) {
    var userId string
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

    userId = utils.GenerateRandomString(10)

    repository = models.Repository{
        ID: utils.GenerateRandomString(10),
        Name: utils.GenerateRandomString(10),
        Description: utils.GenerateRandomString(10),
        URL: utils.GenerateRandomString(10),
        Language: utils.GenerateRandomString(10),
    }

    method = "POST"

    path = fmt.Sprintf("/%s/%s/%s", "users", userId, "repository")

    body = fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","url":"%s","language":"%s"}`, 
        repository.ID, repository.Name, repository.Description, repository.URL, repository.Language)

    request, err = http.NewRequest(method, path, strings.NewReader(body))

    if err != nil {
        t.Fatalf("Failed to create the request: %s", err.Error())
    }

    request.Header.Set("Content-Type", "application/json")

    t.Logf("Request: method=%s, path=%s and body=%s", method, path, body)

    response = httptest.NewRecorder()

    r.ServeHTTP(response, request)

    expectedCode = http.StatusCreated

    if expectedCode != response.Code {
        t.Errorf("Test failed, response: code=%d and body=%+v", response.Code, response.Body)
        return
    }

    err = json.NewDecoder(response.Body).Decode(&repositoryAux)

    if err != nil {
        t.Fatalf("Failed to parse the JSON response body: %s", err.Error())
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

    t.Logf("Test successful, response: code=%d and body=%s", response.Code, string(bodyBytes))
}
