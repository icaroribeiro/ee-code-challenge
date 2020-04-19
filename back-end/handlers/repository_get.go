package handlers

import (
    //"fmt"
    //"github.com/gorilla/mux"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/icaroribeiro/ee-code-challenge/back-end/server"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/services"
    //"github.com/icaroribeiro/ee-code-challenge/back-end/utils"
    "net/http"
    //"strconv"
)

func GetAllUserGithubStarredRepositories(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        /*
        var filter services.Filter
        var err error
        var filtersMap map[string]interface{}
        var authors []models.Author
        var authorFilter map[string]interface{}
        var author models.Author
        var authorsIds []int
        var sort services.Sort
        var ordersMap map[string]string
        var books []models.Book
        var totalCount int
        var pagination models.Pagination
        var page services.Page
        var numberOfSkippedRecords int
        var i int
        var listings []interface{}

        // Filtering scheme.
        filter, err = services.GetFilter(r.URL.Query())

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the filtering fields from the URL: %s", err.Error())})
            return
        }

        // Prepare the filtering fields.
        filtersMap = make(map[string]interface{})

        if !filter.IsEmpty() {
            if filter.Name != "" {
                filtersMap["name"] = filter.Name
            }

            if filter.Edition != 0 {
                filtersMap["edition"] = filter.Edition
            }

            if filter.PublicationYear != 0 {
                filtersMap["publication_year"] = filter.PublicationYear
            }

            // Filter the books by the name of their authors.
            if filter.Author != "" {
                authorFilter = make(map[string]interface{})

                authorFilter["name"] = filter.Author

                authors, err = s.Datastore.GetAllAuthors(authorFilter, nil)

                if len(authors) > 0 {
                    for _, author = range authors {
                        authorsIds = append(authorsIds, author.ID)
                    }

                    filtersMap["authors"] = authorsIds
                }
            }
        }

        // Ordering scheme.
        sort, err = services.GetOrder(r.URL.Query())

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the ordering fields from the URL: %s", err.Error())})
            return
        }

        // Prepare the ordering fields.
        ordersMap = make(map[string]string)

        if !sort.IsEmpty() {
            switch sort.Param {
            case "id", "name", "publication_year":
                ordersMap[sort.Param] = sort.Order
            default:
            }
        }

        // Get all records by using the filtering and ordering settings.
        books, err = s.Datastore.GetAllBooks(filtersMap, ordersMap)

        if err != nil {
            utils.RespondWithJson(w, http.StatusInternalServerError, 
                map[string]string{"error": fmt.Sprintf("Failed to get the list of all books: %s", err.Error())})
            return
        }

        totalCount = len(books)

        if totalCount == 0 {
            pagination = models.Pagination{
                PageNumber: 1,
                PageSize:   0,
                TotalCount: 0,
                Listings:   nil,
            }

            utils.RespondWithJson(w, http.StatusOK, pagination)
            return
        }

        // Pagination scheme.
        page, err = services.GetPage(r.URL.Query())

        if err != nil {
            utils.RespondWithJson(w, http.StatusBadRequest, 
                map[string]string{"error": fmt.Sprintf("Failed to get the pagination fields from the URL: %s", err.Error())})
            return
        }

        // Prepare the pagination fields.
        if page.IsEmpty() {
            page.Number = 1
            page.Size = totalCount
        } else {
            services.DefinePaginationSettings(totalCount, &page, &numberOfSkippedRecords)
        }

        // In the case of there is no skip, the numberOfSkippedRecords is 0 and
        // therefore the records will be gathered from the beginning.
        for i = numberOfSkippedRecords; i < (page.Size + numberOfSkippedRecords) ; i++ {
            listings = append(listings, books[i])
        }

        pagination = models.Pagination{
            PageNumber: page.Number,
            PageSize:   page.Size,
            TotalCount: totalCount,
            Listings:   listings,
        }

        utils.RespondWithJson(w, http.StatusOK, pagination)
        */
    })
}

func GetAllUserRepositories(s *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        /*var params map[string]string
        var bookId int
        var book models.Book
        var err error

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

        if book.ID == 0 {
            utils.RespondWithJson(w, http.StatusNotFound, 
                map[string]string{"error": fmt.Sprintf("Failed to get the book with the id %d: the book wasn't found", bookId)})
            return
        }

        utils.RespondWithJson(w, http.StatusOK, book)
        */
    })
}
