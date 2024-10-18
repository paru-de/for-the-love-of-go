package bookstore

import (
	"errors"
	"slices"
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
func GetBook(catalog []Book, id int) (BookInfo Book, err error) {
	bookIndex := slices.IndexFunc(catalog, func(b Book) bool {
		return b.ID == id
	})

	if bookIndex != -1 {
		return catalog[bookIndex], nil
	}

	return BookInfo, errors.New("book doesn't exist")
}
