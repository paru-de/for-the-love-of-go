package main

import "fmt"

func main() {
	// Book represents information about a book.
	type Book struct {
		Title, Author   string
		Copies, Edition int
		Promotion       bool
		SpecialOffer    float64
	}
	// Customer represents information about a customer.
	type Customer struct {
		Name, FirstName, Email, Street, State, Country, City string
		StreetNumber, ZIP                                    int
	}

	b := Book{}
	c := Customer{}

	b.Title = "All Quiet on the Western Front"
	b.Author = "Remarque, Erich Maria"
	b.Copies = 99
	b.Edition = 1
	b.Promotion = true
	b.SpecialOffer = 15

	c.Name = "Doe"
	c.FirstName = "John"
	c.Email = "booklover@mail.com"
	c.Street = "Hauptstr."
	c.StreetNumber = 5
	c.State = "Baden-Württemberg"
	c.Country = "DE"
	c.City = "Freiburg"
	c.ZIP = 79108

	fmt.Printf("Logged in as %s. Welcome %s %s!\n\n", c.Email, c.FirstName, c.Name)
	fmt.Printf("Shipping address:\n%s %s\n%s %d\n%d %s\n%s-%s\n\n", c.FirstName, c.Name, c.Street, c.StreetNumber, c.ZIP, c.City, c.Country, c.State)
	fmt.Printf("Search results:\n%s: %s (%d. edition)\n", b.Author, b.Title, b.Edition)
	fmt.Printf("Copies in stock: %d\n", b.Copies)

	if b.Promotion {
		fmt.Printf("Special Offer! %.2f%% off - buy now!\n", b.SpecialOffer)
	}
}
