package service

import (
	"github.com/jurazodda/note-manager-go/internal/errs"
	"github.com/jurazodda/note-manager-go/internal/models"
	"github.com/jurazodda/note-manager-go/internal/repository"
)

func CreateNote(note models.Note) (models.Note, error) {
	if note.Title == "" {
		return models.Note{}, errs.ErrNoteTitleRequired
	}

	if note.Content == "" {
		return models.Note{}, errs.ErrNoteContentRequired
	}

	note, err := repository.CreateNote(note)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func GetNotes() ([]models.Note, error) {
	notes, err := repository.GetNotes()
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func GetNoteByID(noteID int) (models.Note, error) {
	if noteID <= 0 {
		return models.Note{}, errs.ErrNoteIDRequired
	}

	note, err := repository.GetNoteByID(noteID)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func UpdateNote(note models.Note) (models.Note, error) {
	if note.ID <= 0 {
		return models.Note{}, errs.ErrNoteIDRequired
	}

	if note.Title == "" {
		return models.Note{}, errs.ErrNoteTitleRequired
	}

	if note.Content == "" {
		return models.Note{}, errs.ErrNoteContentRequired
	}

	note, err := repository.UpdateNote(note)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func DeleteNote(noteID int) error {
	if noteID <= 0 {
		return errs.ErrNoteIDRequired
	}

	err := repository.DeleteNote(noteID)
	if err != nil {
		return err
	}

	return nil
}
