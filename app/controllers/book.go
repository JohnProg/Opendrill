package controllers

import (
	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"

	models "../models"
)

func ListBooks(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	books, _ := models.AllBooks()
	if books == nil {
		return []models.Book{}, nil
	}
	return books, nil
}

func GetBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, b := models.GetBook(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find book " + Id, http.StatusNotFound}
	}
	return b, nil
}

func AddBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.Book
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, book := models.CreateBook(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create book ", http.StatusNotFound}
	}
	return book, nil
}

func UpdateBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.Book
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, book := models.UpdateBook(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update book " + Id + " to update", http.StatusNotFound}
	}
	return book, nil
}

func RemoveBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveBook(Id)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find book " + Id + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
