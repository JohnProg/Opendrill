package controllers

import (
	"github.com/gorilla/mux"

	"net/http"
	"encoding/json"

	models "../models"
)


func ListListContacts(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	list_contacts, _ := models.AllListContact()
	if list_contacts == nil {
		return []models.ListContact{}, nil
	}
	return list_contacts, nil
}

func GetListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id es not valid", http.StatusBadRequest}
	}
	err, b := models.GetListContact(Id)
	if err != nil{
		return nil, &models.HandlerError{err, "Could not find list contact" + Id, http.StatusNotFound}
	}
	return b, nil	
}


func AddListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	var payload models.ListContact

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{err, "Could not parse JSON", http.StatusNotFound}
	}
	err, list_contact := models.CreateListContact(payload)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not create list contact", http.StatusNotFound}
	}
	return list_contact, nil
} 

func UpdateListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError){
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	var payload models.ListContact
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, &models.HandlerError{ err, "Could not parse JSON", http.StatusNotFound}
	}
	err, list_contact := models.UpdateListContact(payload, Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not update list contact" + Id + "to update", http.StatusNotFound}
	}
	return list_contact, nil
}

func RemoveListContact(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	Id := mux.Vars(r)["id"]
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}
	err, deleted := models.RemoveListContact(Id)
	if err != nil {
		return nil, &models.HandlerError{err, "Could not find list contact" + Id + " to delete", http.StatusNotFound}
	}
	return deleted, nil

}