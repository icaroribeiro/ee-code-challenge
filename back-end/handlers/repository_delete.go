package handlers

import (
    //"fmt"
    //"github.com/gorilla/mux"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
    //"strconv"
)

func DeleteUserRepository(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        /*var params map[string]string
        var bookId int
        var book models.Book
        var err error
        var nDeletedDocs int64

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

        book, err = s.Datastore.GetBook(bookId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the book with the id %d: %s", bookId, err.Error())})
            return
        }

        nDeletedDocs, err = s.Datastore.DeleteBook(bookId)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the book with the id %d: %s", bookId, err.Error())})
            return
        }

        if nDeletedDocs == 0 {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the book with the id %d: the book wasn't found", bookId)})
            return
        }

        if nDeletedDocs > 1 {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to delete the book with the id %d: the expected number of " +
                    "books deleted: %d, got: %d", bookId, 1, nDeletedDocs)})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, book)
        */
    })
}
