package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

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
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool { return got[i].ID < got[j].ID })
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Title: "For the Love  of Go", ID: 1},
		2: {Title: "The Power of Go: Tools", ID: 2},
	}
	want := bookstore.Book{Title: "The Power of Go: Tools", ID: 2}

	got, err := bookstore.GetBook(catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {},
	}
	_, err := bookstore.GetBook(catalog, 2)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()

	if want != got {
		t.Errorf("NetPriceCents: want %d, got %d", want, got)
	}
}
