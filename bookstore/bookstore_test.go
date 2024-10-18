package bookstore_test

import (
	"testing"

	"bookstore"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Maria Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Maria Kondo",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorNoCopies(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Fatal("expected error buying book with zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go", ID: 1},
		{Title: "The Power of Go: Tools", ID: 2},
	}
	want := bookstore.Book{Title: "The Power of Go: Tools", ID: 2}

	got, _ := bookstore.GetBook(catalog, 2)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{{Title: "The Hobbit", ID: 3}}
	_, err := bookstore.GetBook(catalog, 4)
	if err == nil {
		t.Fatal("Book ID doesn't exist: expected an error, got nil")
	}
}

func TestGetBookEmptyCatalog(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 1)
	if err == nil {
		t.Fatal("Book Catalog empty: expected an error, got nil")
	}
}
