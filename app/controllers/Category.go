package controllers

import (
	"github.com/gorilla/mux"

	"net/http"
	"log"
	"encoding/json"

	models "../models"
)


func ListCategories(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	categories, _ := models.AllCategory()
	if categories == nil {
		return []models.Category{}, nil
	}
	return categories, nil
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, b := models.GetCategory(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + Id, http.StatusNotFound}
	}
	return b, nil
}

func AddCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.CreateCategory(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create category ", http.StatusNotFound}
	}
	return category, nil
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, category := models.UpdateCategory(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update category " + Id + " to update", http.StatusNotFound}
	}
	return category, nil
}

func RemoveCategory(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	log.Println(Id)
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveCategory(Id)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find category " + Id + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
