package handler

import (
	"bufio"
	"fmt"
	"log"
	"github.com/jurazodda/note-manager-go/internal/models"
	"github.com/jurazodda/note-manager-go/internal/service"
	"os"
	"strings"
)

func CreateNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	title = strings.TrimSpace(title)

	fmt.Print("Enter note content: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	content = strings.TrimSpace(content)

	toCreate := models.Note{
		Title: title,
		Content: content,
	}

	note, err := service.CreateNote(toCreate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Note created:", note)
}

func GetNotes() {
	notes, err := service.GetNotes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range notes {
		fmt.Printf("ID: %d. Title: %s Content: %s\n", notes[i].ID, notes[i].Title, notes[i].Content)
	}
}

func GetNoteByID() {
	fmt.Print("Enter the ID of the note you want to view: ")
	var noteID int
	fmt.Scan(&noteID)

	note, err := service.GetNoteByID(noteID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d. Title: %s Content: %s\n", note.ID, note.Title, note.Content)
}

func UpdateNote() {
	fmt.Print("Enter the ID of the note you want to update: ")
	var noteID int
	fmt.Scan(&noteID)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a new title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	title = strings.TrimSpace(title)

	fmt.Print("Enter a new content: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	content = strings.TrimSpace(content)

	toUpdate := models.Note{
		ID: noteID,
		Title: title,
		Content: content,
	}

	note, err := service.UpdateNote(toUpdate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Note updated:", note)
}

func DeleteNote() {
	fmt.Print("Enter the ID of the note you want to delete: ")
	var noteID int
	fmt.Scan(&noteID)

	err := service.DeleteNote(noteID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Note deleted.")
}
