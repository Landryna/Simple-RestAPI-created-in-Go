package note

import "github.com/google/uuid"

type Note struct {
	Details  NoteDetails `json:"details"`
	Creation string      `json: "creation"`
	ID       uuid.UUID   `json: "id"`
}

type NoteDetails struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var Notes map[string]Note

func init() {
	Notes = make(map[string]Note)
}
