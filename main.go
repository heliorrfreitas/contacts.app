package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Contact struct{
	Id int
	First string
	Last string
	Phone string
	Email string
}

type ContactsPageData struct{
	PageTitle string
	Contacts []Contact
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomePlaceholder).Methods("GET")
	router.HandleFunc("/contacts", Contacts).Methods("GET")
	router.HandleFunc("/contacts/new", ContactsNewGet).Methods("GET")
	router.HandleFunc("/contacts/new", ContactsNew).Methods("POST")

	http.ListenAndServe(":80", router)
}

func HomePlaceholder(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "redirect to contacts")
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}
func Contacts(w http.ResponseWriter, r *http.Request){
	data := ContactsPageData{
		PageTitle: "Contacts",
		Contacts: []Contact{
			{Id: 2, First: "Edward", Last: "Newgate", Phone: "123-456-7890", Email: "newgate@example.comz"},
		},
	}
	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, data)
}

func ContactsNewGet(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "adding new contact page, in its own function")
}

func ContactsNew(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "adding new contact in here")
}
