package handler

import "fmt"

var Menu = `
0. Exit
1. Add note
2. Show notes
3. Show note by id
4. Update note
5. Delete note

Choose a command: `

func RunMenu() {
	fmt.Println("__ Console Notes manager __")
	for {
		fmt.Print(Menu)
		var command string
		fmt.Scan(&command)
		switch command {
		case "0":
			fmt.Println("Exiting...")
			return
		case "1":
			CreateNote()
		case "2":
			GetNotes()
		case "3":
			GetNoteByID()
		case "4":
			UpdateNote()
		case "5":
			DeleteNote()
		default:
			fmt.Printf("Unknown command: %s, please try again\n", command)
		}
	}
}