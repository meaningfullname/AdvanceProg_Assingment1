package main

import (
	"fmt"
)

type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[int]Book
}

var lib Library

func Add(id int, title string, author string) {
	book := Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	lib.Books[id] = book
}

func Borrow(id int) {
	if book, exists := lib.Books[id]; exists {
		if book.IsBorrowed {
			fmt.Println("This book is already borrowed.")
		} else {
			book.IsBorrowed = true
			lib.Books[id] = book
			fmt.Println("You borrowed:", book.Title)
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func ReturnBook(id int) {
	if book, exists := lib.Books[id]; exists {
		if !book.IsBorrowed {
			fmt.Println("This book was not borrowed.")
		} else {
			book.IsBorrowed = false
			lib.Books[id] = book
			fmt.Println("You returned:", book.Title)
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func ListBooks() {
	if len(lib.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	for _, book := range lib.Books {
		status := "Available"
		if book.IsBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
	}
}

func main() {
	var id int = 0
	var choice int
	var title, author string

	lib.Books = make(map[int]Book)

	for {
		fmt.Println("\n1: Add a book")
		fmt.Println("2: Borrow a book")
		fmt.Println("3: Return a book")
		fmt.Println("4: List books")
		fmt.Println("5: Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add a new book
			fmt.Println("Input Title:")
			fmt.Scan(&title)
			fmt.Println("Input Author:")
			fmt.Scan(&author)
			Add(id, title, author)
			id++
		case 2:

			var borrowID int
			fmt.Println("Which book you want to borrow? (Enter ID):")
			fmt.Scan(&borrowID)
			Borrow(borrowID)
		case 3:

			var returnID int
			fmt.Println("Which book you want to return? (Enter ID):")
			fmt.Scan(&returnID)
			ReturnBook(returnID)
		case 4:

			ListBooks()
		case 5:

			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
