package utils

import (
	"SEMI_TRASH_API/handlers"
	"github.com/gorilla/mux"
)

/*
	Раньше в первой версии TRASH_API мы инициализовали связку роутер - хендлер в функции main, это не есть хорошо
	лучше создавать отдельную функцию для этого. мы создали аж целый пакет
	матчим все адреса, с исполнителями (хендлерами, которые будут обрабатывать запросы)
*/

// фнкция ничего не возвращает, она просто изменяет внутреннее сосотояние роутера
func BuildManyBooksResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}

func BuildBookResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	r.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	r.HandleFunc(prefix+"/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc(prefix+"/{id}", handlers.DeleteBook).Methods("DELETE")
}
