Contact App in Go

The project uses golang in the backend 

This is supposed to be a web 1.0 version of a contact app

HTML/CSS/Golang are used to build a rest api

initiate the project

go mod init main

then I created the main file called main.go

then I imported the pckages used in the example such as fmt, net/http, and github.com/gorilla/mux

this last one required me to use another command to import it to the project, I think it is because it gets it from github. The command was the following:

go get github.com/gorilla/mux

then to run just ran go run main.go 

contacts.json is from https://github.com/bigskysoftware/contact-app


