// STRUCTS {STRUCTURES}
// Defining a structure
// Keywords: type and struct

package main

import "fmt"

func main() {

	type Book struct {
		Title   string
		Author  string
		Subject string
		BookID  int
		Read    bool
	}

	book1 := Book{
		Title:   "Shawshank Redemption",
		Author:  "Stephen King",
		Subject: "Novel about prisoners",
		BookID:  123,
		Read:    true,
	}

	fmt.Println("Title:", book1.Title)
	fmt.Println("Author:", book1.Author)
	fmt.Println("Subject:", book1.Subject)
	fmt.Println("Book ID:", book1.BookID)
	fmt.Println("Read already ?:", book1.Read)
}
