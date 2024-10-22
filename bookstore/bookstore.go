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

type Catalog map[int]Book

// Buy takes type Book as input and updates the remaining Book.Copies count
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

// Method GetAllBooks returns a slice of all books available in the catalogue
func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

// Method GetBook returns book information using its unique ID
// and returns an error if the ID was not found
func (c Catalog) GetBook(ID int) (Book, error) {
	_, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("Error: Book ID %d not found.", ID)
	}
	return c[ID], nil
}

// Method NetPriceCents returns the price of a book including any discounts
func (b Book) NetPriceCents() (price int) {
	if !(b.DiscountPercent < 1 || b.DiscountPercent > 100) {
		discountAmount := (b.PriceCents * b.DiscountPercent) / 100
		return b.PriceCents - discountAmount
	}
	return b.PriceCents
}
