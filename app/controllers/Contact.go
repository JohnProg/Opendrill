package controllers

import (
	"github.com/gorilla/mux"

	"net/http"
	"log"
	"encoding/json"

	models "../models"
)


func ListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	contacts, _ := models.AllContact()
	if contacts == nil {
		return []models.Contact{}, nil
	}
	return contacts, nil
}

func GetContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, b := models.GetContact(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find contact " + Id, http.StatusNotFound}
	}
	return b, nil
}

func AddContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.Contact
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, contact := models.CreateContact(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create contact ", http.StatusNotFound}
	}
	return contact, nil
}

func UpdateContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.Contact
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON ", http.StatusNotFound}
	}
	err, contact := models.UpdateContact(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update contact " + Id + " to update", http.StatusNotFound}
	}
	return contact, nil
}

func RemoveContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	log.Println(Id)
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveContact(Id)

	if err != nil {
		return nil, &models.HandlerError{err, "Could not find contact " + Id + " to delete", http.StatusNotFound}
	}
	return deleted, nil
}
