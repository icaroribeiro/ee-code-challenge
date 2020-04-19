package handlers

import (
    //"encoding/json"
    //"fmt"
    //"github.com/gorilla/mux"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
    //"strconv"
)

func UpdateUserRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        /*
        var params map[string]string
        var bookId int
        var err error
        var book models.Book
        var body string
        var authorsMap map[int]bool
        var i int
        var authorId int
        var author models.Author
        var nRowsAffected int64

        params = mux.Vars(r)

        if params["bookId"] == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The id is required and must be set to a non-empty value in the request URL"})
            return
        }

        bookId, err = strconv.Atoi(params["bookId"])

        if err != nil {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": fmt.Sprintf("Failed to convert the string %s to a numeric value", params["bookId"])})
            return
        }

        err = json.NewDecoder(r.Body).Decode(&book)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to decode the request body: %s", err.Error())})
            return
        }

        if book.Name == "" {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The name field is required and must be set to a non-empty value"})
            return
        }

        if book.Edition <= 0 {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The edition field is required and must be set to a value greater than 0"})
            return
        }

        if book.PublicationYear < 1 || book.PublicationYear > 9999 {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The publication year field is required and must be set to a value " +
                    "in the range from 1 to 9999"})
            return
        }

        if len(book.Authors) == 0 {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": "The authors field is required and must be set to an array containing " +
                    "at least one author id"})
            return
        }

        body = fmt.Sprintf(`{"name":"%s","edition":%d,"publication_year":%d`,
            book.Name, book.Edition, book.PublicationYear)

        // Verify if all the ids of the authors associated with the book are valid.
        // Additionally, checks if there are no duplicate ids of the authors.
        authorsMap = make(map[int]bool)

        for i, authorId = range book.Authors {
            author, err = s.Datastore.GetAuthor(authorId)

            if err != nil {
                utils.RespondWithJson(w, http.StatusInternalServerError, 
                    map[string]string{"error": fmt.Sprintf("Failed to add the author with id %d: %s", authorId, err.Error())})
                return
            }

            if author.ID == 0 {
                utils.RespondWithJson(w, http.StatusNotFound, 
                    map[string]string{"error": fmt.Sprintf("Failed to add the author with id %d: the author wasn't found", authorId)})
                return
            }

            if !(authorsMap[author.ID]) {
                authorsMap[author.ID] = true
            } else {
                utils.RespondWithJson(w, http.StatusBadRequest, 
                    map[string]string{"error": fmt.Sprintf("Failed to add the author with id %d: the id is duplicated", author.ID)})
                return
            }

            if i == 0 {
                body += fmt.Sprintf(`,"authors":[%d`, author.ID)
            } else {
                body += fmt.Sprintf(`,%d`, author.ID)
            }
        }

        body += `]}`

        nRowsAffected, err = s.Datastore.UpdateBook(bookId, book)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the book with the id %d with %s: %s", 
                    bookId, body, err.Error())})
            return
        }

        book.ID = bookId

        if nRowsAffected == 0 {
            utils.RespondWithJson(w, http.StatusConflict, 
                map[string]string{"error": fmt.Sprintf("Failed to update the book with the id %d with %s: " +
                    "the book wasn't found", bookId, body)})
            return
        }

        if nRowsAffected != 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to update the book with the id %d with %s: " + 
                    "the expected number of books updated: %d, got: %d", bookId, body, 1, nRowsAffected)})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, book)
        */
    })
}
