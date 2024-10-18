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
