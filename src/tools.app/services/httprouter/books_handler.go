package httprouter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewBooksHandler returns a books http router handler.
func NewBooksHandler() *BooksHandler {
	books := make([]Book, 2)
	books = append(books, Book{
		ISDN:   "001",
		Title:  "Silence of the Lambs",
		Author: "Thomas Harris",
		Pages:  367,
	})
	books = append(books, Book{
		ISDN:   "002",
		Title:  "To Kill a Mocking Bird",
		Author: "Harper Lee",
		Pages:  320,
	})

	handler := &BooksHandler{
		BooksStore: make(map[string]*Book),
	}
	for _, book := range books {
		handler.BooksStore[book.ISDN] = &book
	}
	return handler
}

// ********* Models

// Book book object
type Book struct {
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// ********* Books Handlers

// BooksHandler books http handlers.
type BooksHandler struct {
	BooksStore map[string]*Book
}

// Index handler for home page.
// Get /index
func (handler BooksHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

// BookCreate handler for the books Create action.
// POST /books
func (handler *BooksHandler) BookCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	book := &Book{}
	if err := handler.populateModelFromHandler(r, book); err != nil {
		WriteErrResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	handler.BooksStore[book.ISDN] = book
	WriteOKResponse(w, book)
}

// BookIndex handler for the books index action.
// GET /books
func (handler BooksHandler) BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range handler.BooksStore {
		books = append(books, book)
	}
	WriteOKResponse(w, books)
}

// BookShow handler for the books Show action.
// GET /books/:isdn
func (handler BooksHandler) BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	book, ok := handler.BooksStore[isdn]
	if !ok {
		WriteErrResponse(w, http.StatusNotFound, "Book Record Not Found!")
		return
	}
	WriteOKResponse(w, book)
}

// populateModelFromHandler populates a model from the params in the Handler.
func (handler BooksHandler) populateModelFromHandler(r *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	defer r.Body.Close()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

// ********* HTTP Response

// JSONResponse json http response
type JSONResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// JSONErrResponse json http error response
type JSONErrResponse struct {
	Error *APIError `json:"error"`
}

// APIError json http error response
type APIError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// WriteOKResponse writes the response as a standard JSON response with StatusOK.
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: m}); err != nil {
		WriteErrResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteErrResponse writes the error response as a Standard API JSON response with a response code.
func WriteErrResponse(w http.ResponseWriter, errCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&JSONErrResponse{
		Error: &APIError{Status: errCode, Title: errMsg},
	})
}
