package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
	"io/ioutil"
)

type Contact struct{
	Id int	`json:"id"`
	First string `json:"first"`
	Last string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
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
	var objContacts []Contact
	data, err := ioutil.ReadFile("contacts.json")
	if err != nil {
		fmt.Println("error while reading the file:", err)
	} else {
		err = json.Unmarshal(data, &objContacts)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
	}

	displayData := ContactsPageData{
		PageTitle: "Contacts List",
		Contacts: objContacts,
	}
	
	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, displayData)
}

func ContactsNewGet(w http.ResponseWriter, r *http.Request){
	contactInfo := Contact{Id: 10, First: "Test", Last: "2", Phone: "1234234", Email: "email@email.com"}
	tmpl, err := template.ParseFiles("new.html")
	if err != nil {
		fmt.Println("Error while parsing new.html template: ", err)	
	}
	err = tmpl.Execute(w, contactInfo)

	if err != nil {
		fmt.Println("Error while executing new.html template: ", err)	
	}
}

func ContactsNew(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "adding new contact in here")
}

