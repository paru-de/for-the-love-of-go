package bookstore

import "errors"

type Book struct {
	Title, Author string
	Copies        int
}

// Buy takes type Book as input and updates the remaining Book.Copies count
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
