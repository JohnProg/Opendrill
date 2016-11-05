package router

import (
	"../app/controllers"
	"../app/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("../public/"))

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	// router.Handle("/books", handler(controllers.ListBooks)).Methods("GET")
	// router.Handle("/books", handler(controllers.AddBook)).Methods("POST")
	// router.Handle("/books/{id}", handler(controllers.GetBook)).Methods("GET")
	// router.Handle("/books/{id}", handler(controllers.UpdateBook)).Methods("PUT")
	// router.Handle("/books/{id}", handler(controllers.RemoveBook)).Methods("DELETE")
	//Category
	router.Handle("/categories/", middlewares.Handler(controllers.ListCategories)).Methods("GET")
	router.Handle("/categories/", middlewares.Handler(controllers.AddCategory)).Methods("POST")
	router.Handle("/categories/{id}", middlewares.Handler(controllers.GetCategory)).Methods("GET")
	router.Handle("/categories/{id}", middlewares.Handler(controllers.UpdateCategory)).Methods("PUT")
	router.Handle("/categories/{id}", middlewares.Handler(controllers.RemoveCategory)).Methods("DELETE")
	//Contact
	router.Handle("/contacts/", middlewares.Handler(controllers.ListContacts)).Methods("GET")
	router.Handle("/contacts/", middlewares.Handler(controllers.AddContact)).Methods("POST")
	router.Handle("/contacts/{id}", middlewares.Handler(controllers.GetContact)).Methods("GET")
	router.Handle("/contacts/{id}", middlewares.Handler(controllers.UpdateContact)).Methods("PUT")
	router.Handle("/contacts/{id}", middlewares.Handler(controllers.RemoveContact)).Methods("DELETE")
	//List Contact
	router.Handle("/list-contacts/", middlewares.Handler(controllers.ListListContacts)).Methods("GET")
	router.Handle("/list-contacts/", middlewares.Handler(controllers.AddListContact)).Methods("POST")
	router.Handle("/list-contacts/{id}", middlewares.Handler(controllers.GetListContact)).Methods("GET")
	router.Handle("/list-contacts/{id}", middlewares.Handler(controllers.UpdateListContact)).Methods("PUT")
	router.Handle("/list-contacts/{id}", middlewares.Handler(controllers.RemoveListContact)).Methods("DELETE")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.Public)))
	http.Handle("/", router)
}
