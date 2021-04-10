package main

import (
	"demo/src/controllers"
	"demo/src/services"
	"fmt"
	"github.com/gorilla/pat"
	"log"
	"net/http"
)

func main() {
	fmt.Print("	Initializing server")
	notesServer := controllers.Server{Handler: pat.New()}

	notesServer.HandleGetRequestPath("/notes/{id}", service.GetNote)
	notesServer.HandleDeleteRequestPath("/notes/{id}", service.RemoveNote)
	notesServer.HandleGetRequestPath("/notes", service.ListNotes)
	notesServer.HandlePostRequestPath("/notes", service.CreateNote)
	log.Fatal(http.ListenAndServe(":9999", notesServer.Handler))
}
