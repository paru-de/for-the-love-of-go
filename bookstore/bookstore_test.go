package bookstore_test

import (
	"testing"

	"bookstore"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Maria Kondo",
		Copies: 2,
	}
}
