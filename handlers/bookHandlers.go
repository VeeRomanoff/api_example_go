package handlers

import (
	"SEMI_TRASH_API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	json.NewEncoder(w).Encode(models.DataBase)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{
			Message: "failed to convert string to int",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	book, ok := findBookById(id, &models.DataBase)
	if !ok {
		msg := models.Message{
			Message: "book not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var book models.Book
	// анмаршаллинг. первый вызов - ИЗ ЧЕГО (толо запроса), второй вызов - КУДА
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		msg := models.Message{
			Message: "failed to read request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// сколько-то там длинна нашего массива - + 1 - это айди следующего элемента в "бд"
	var newBookID int = len(models.DataBase) + 1
	book.ID = newBookID
	models.DataBase = append(models.DataBase, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	// работа с id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{
			Message: "failed to convert string to int",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// нахождение книги в "бд"
	bookToChange, ok := findBookById(id, &models.DataBase)
	var newBook models.Book
	if !ok {
		msg := models.Message{
			Message: "book not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{
			Message: "provided json file is invalid",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	newBook.ID = bookToChange.ID
	for i, b := range models.DataBase {
		if b.ID == newBook.ID {
			models.DataBase[i] = newBook
			break
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{
			Message: "failed to convert string to int",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	bookToDelete, ok := findBookById(id, &models.DataBase)
	if !ok {
		msg := models.Message{
			Message: "book not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	removeBookById(bookToDelete.ID, &models.DataBase)

	msg := models.Message{
		Message: "book deleted",
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(msg)
}

func findBookById(id int, db *[]models.Book) (*models.Book, bool) {
	for _, b := range *db {
		if b.ID == id {
			return &b, true
		}
	}
	return &models.Book{}, false
}

func removeBookById(id int, db *[]models.Book) {
	for i, b := range *db {
		if b.ID == id {
			*db = append((*db)[:i], (*db)[i+1:]...)
			return
		}
	}
}
