package repository

import (
	"fmt"
	"github.com/jurazodda/note-manager-go/internal/errs"
	"github.com/jurazodda/note-manager-go/internal/models"
)

var notes []models.Note
var nextID int 

func CreateNote(note models.Note) (models.Note, error) {
	nextID++
	note.ID = nextID
	notes = append(notes, note)

	return note, nil
}

func GetNotes() ([]models.Note, error) {
	if len(notes) == 0 {
		fmt.Println("Notes:", notes)
		return nil, nil
	}

	return notes, nil
}

func GetNoteByID(noteID int) (models.Note, error) {
	for i := range notes {
		if notes[i].ID == noteID {
			return notes[i], nil
		}
	}
	return models.Note{}, errs.ErrNoteNotFound
}

func UpdateNote(note models.Note) (models.Note, error) {
	for i := range notes {
		if notes[i].ID == note.ID {
			notes[i].Title = note.Title
			notes[i].Content = note.Content
			return notes[i], nil
		}
	}
	return models.Note{}, errs.ErrNoteNotFound
}

func DeleteNote(noteID int) error {
	for i := range notes {
		if notes[i].ID == noteID {
			notes = append(notes[:i], notes[i+1:]...)
			return nil
		}
	}

	return errs.ErrNoteNotFound
}