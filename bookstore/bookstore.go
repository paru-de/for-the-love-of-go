package bookstore

import (
	"errors"
)

type Book struct {
	Title, Author string
	Copies, ID    int
}

// Buy takes type Book as input and updates the remaining Book.Copies count
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

// GetAllBooks returns a slice of all books available in the catalogue
func GetAllBooks(catalog []Book) []Book {
	return catalog
}

// GetBook returns book information using its unique ID
// and returns an error if the ID was not found
func GetBook(catalog map[int]Book, ID int) (Book, error) {
	_, ok := catalog[ID]
	if !ok {
		return Book{}, errors.New("Error: Book ID not found.")
	}
	return catalog[ID], nil
}

// old implementation
// func GetBook(catalog []Book, id int) (bookData Book, err error) {
// 	bookIndex := slices.IndexFunc(catalog, func(b Book) bool {
// 		return b.ID == id
// 	})
//
// 	if bookIndex != -1 {
// 		return catalog[bookIndex], nil
// 	}
//
// 	return bookData, errors.New("book doesn't exist")
// }
