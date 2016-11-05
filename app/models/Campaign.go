package models

type Campaign struct{
	Id string
	Name string
	FromName string
	ReplyEmail string
	Subject string
	Status bool
	//Foreign Keys
	Template int
	ListContact int
}