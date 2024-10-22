package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title, Author               string
	Copies, ID                  int
	PriceCents, DiscountPercent int
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
func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

// GetBook returns book information using its unique ID
// and returns an error if the ID was not found
func GetBook(catalog map[int]Book, ID int) (Book, error) {
	_, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("Error: Book ID %d not found.", ID)
	}
	return catalog[ID], nil
}

// Method NetPriceCents returns the price of a book including any discounts
func (b Book) NetPriceCents() (price int) {
	if !(b.DiscountPercent < 1 || b.DiscountPercent > 100) {
		discountAmount := (b.PriceCents * b.DiscountPercent) / 100
		return b.PriceCents - discountAmount
	}
	return b.PriceCents
}

// old implementation has a complexity of O(n) (vs. new one with O(1))
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
