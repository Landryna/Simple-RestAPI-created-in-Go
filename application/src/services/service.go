package service

import (
	"demo/src/note"
	"demo/src/utils"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

const DATA_FORMAT = "Jan 2 2006 15:04:05"

func sendResponse(res http.ResponseWriter, json []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func ListNotes(res http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(note.Notes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	sendResponse(res, js)
}

func GetNote(res http.ResponseWriter, req *http.Request) {
	noteId := req.URL.Query().Get(":id")
	singleNote := note.Notes[noteId]
	if singleNote.Creation == "" {
		http.Error(res, "Note with provided ID does not exist", http.StatusNotFound)
	}

	js, err := json.Marshal(singleNote)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	sendResponse(res, js)
}

func CreateNote(res http.ResponseWriter, req *http.Request) {
	var noteDetails note.NoteDetails
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&noteDetails)
	if err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		// type missmatch
		case errors.As(err, &unmarshalTypeError):
			http.Error(res, err.Error(), http.StatusBadRequest)
		// empty body
		case errors.Is(err, io.EOF):
			http.Error(res, err.Error(), http.StatusBadRequest)
		// typicaly key not found
		default:
			http.Error(res, "invalid key", http.StatusBadRequest)
		}
	}

	id := uuid.New()
	note.Notes[id.String()] = note.Note{
		Details:  noteDetails,
		Creation: time.Now().Format(DATA_FORMAT),
		ID:       id,
	}

	response := utils.HttpResponse{
		Message:    "Note created.",
		StatusCode: http.StatusAccepted,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	sendResponse(res, js)
}

func RemoveNote(res http.ResponseWriter, req *http.Request) {
	noteId := req.URL.Query().Get(":id")
	singleNote := note.Notes[noteId]
	if singleNote.Creation == "" {
		http.Error(res, "Note with provided ID does not exist", http.StatusNotFound)
	}
	delete(note.Notes, noteId)

	response := utils.HttpResponse{
		Message:    "Note removed.",
		StatusCode: http.StatusAccepted,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	sendResponse(res, js)
}
