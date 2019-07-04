package httprouter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	myutils "tools.app/utils"
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
func (handler *BooksHandler) BookCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	book := &Book{}
	if err := handler.populateModelFromHandler(r, book); err != nil {
		myutils.WriteErrJSONResp(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	handler.BooksStore[book.ISDN] = book
	myutils.WriteOKJSONResp(w, book)
}

// BookIndex handler for the books index action.
// GET /books
func (handler BooksHandler) BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range handler.BooksStore {
		books = append(books, book)
	}
	myutils.WriteOKJSONResp(w, books)
}

// BookShow handler for the books Show action.
// GET /books/:isdn
func (handler BooksHandler) BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	book, ok := handler.BooksStore[isdn]
	if !ok {
		myutils.WriteErrJSONResp(w, http.StatusNotFound, "Book Record Not Found!")
		return
	}
	myutils.WriteOKJSONResp(w, book)
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
