package models

import (
	"gopkg.in/mgo.v2/bson"
)

// book model
type Book struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `json:"title"`
	Author string        `json:"author"`
}

func AllBooks() (books2 []Book, err error) {
	err = books.
		Find(nil).
		All(&books2)
	return
}

func CreateBook(book Book) (err error, book2 Book) {
	book2 = book
	book2.Id = bson.NewObjectId()

	if err := books.Insert(book2); err != nil {
		return err, book
	}
	return nil, book2
}

func GetBook(Id string) (err error, book Book) {
	bid := bson.ObjectIdHex(Id)
	err = books.
		FindId(bid).
		One(&book)
	return
}

func RemoveBook(Id string) (err error, deleted bool) {
	deleted = false
	bid := bson.ObjectIdHex(Id)
	err = books.
		Remove(bson.M{"_id": bid})
	if err != nil {
		return err, deleted
	}
	deleted = true
	return nil, deleted
}

func UpdateBook(book Book, Id string) (err error, book2 Book) {
	book2 = book
	bid := bson.ObjectIdHex(Id)
	err = books.Update(bson.M{"_id": bid},
		bson.M{
			"title": book2.Title,
			"author": book2.Author,
			"_id":    bid,
		})
	if err != nil {
		return err, book
	}
	book2.Id = bid
	return nil, book2
}
