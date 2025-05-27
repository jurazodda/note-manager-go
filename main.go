package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var notes []Note
var nextID int = 1

func main() {
	fmt.Println("__ Console Notes manager __")
	for {
		fmt.Println("0. Exit")
		fmt.Println("1. Add note")
		fmt.Println("2. Show notes")
		fmt.Println("3. Show note by id")
		fmt.Println("4. Update note")
		fmt.Println("5. Delete note")
		fmt.Println()
		fmt.Print("Choose a command: ")

		var command int
		fmt.Scan(&command)

		switch command {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			createNote()
		case 2:
			getNotes()
		case 3:
			getNoteByID()
		case 4:
			updateNote()
		case 5:
			deleteNote()
		default:
			fmt.Printf("Unknown command: %d, please try again!\n", command)
		}
	}
}

func createNote() {
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

	notes = append(notes, Note{
		ID: nextID,
		Title: title,
		Content: content,
	})
	nextID++

	fmt.Println("Note created.")
	fmt.Println()
}

func getNotes() {
	fmt.Println("Notes:", notes)
	for i := range notes {
		fmt.Printf("%d. %s %s\n", notes[i].ID, notes[i].Title, notes[i].Content)
	}

	if len(notes) == 0 {
		fmt.Println("No notes to show!")
	}
	fmt.Println()
}

func getNoteByID() {
	fmt.Print("Enter the ID of the note you want to view: ")
	var noteID int
	fmt.Scan(&noteID)

	for i := range notes {
		if notes[i].ID == noteID {
			fmt.Printf("%d. %s %s\n", notes[i].ID, notes[i].Title, notes[i].Content)
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}

func updateNote() {
	fmt.Print("Enter the ID of the note you want to update: ")
	var noteID int 
	fmt.Scan(noteID)

	reader := bufio.NewReader(os.Stdin)
	for i := range notes {
		if notes[i].ID == noteID {
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

			notes[i].Title = title
			notes[i].Content = content

			fmt.Println("Note updated.")
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}

func deleteNote() {
	fmt.Print("Enter the ID of the note you want to delete: ")
	var noteID int
	fmt.Scan(&noteID)

	for i := range notes {
		if notes[i].ID == noteID {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Println("Note deleted.")
			fmt.Println()
			return
		}
	}

	fmt.Println("Note not found.")
	fmt.Println()
}
