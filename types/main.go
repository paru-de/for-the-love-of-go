package main

import "fmt"

func main() {
	type Book struct {
		Title, Author   string
		Copies, Edition int
		Promotion       bool
		SpecialOffer    float64
	}
	b := Book{}
	b.Title = "All Quiet on the Western Front"
	b.Author = "Remarque, Erich Maria"
	b.Copies = 99
	b.Edition = 1
	b.Promotion = true
	b.SpecialOffer = 15

	fmt.Printf("%s: %s (%d. edition)\n", b.Author, b.Title, b.Edition)
	fmt.Printf("Copies in stock: %d\n", b.Copies)

	if b.Promotion {
		fmt.Printf("Special Offer! %.2f%% off - buy now!\n", b.SpecialOffer)
	}
}
