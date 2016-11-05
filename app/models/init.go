package models

import (
	"gopkg.in/mgo.v2"
)

// error response contains everything we need to use http.Error
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

var (
	db 				*mgo.Database
	books 			*mgo.Collection
	contacts 		*mgo.Collection
	list_contacts 	*mgo.Collection
	categories 		*mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB

	books = db.C("books")
	contacts = db.C("contacts")
	list_contacts = db.C("list_contacts")
	categories = db.C("categories")

	contacts.EnsureIndexKey("email")
}
