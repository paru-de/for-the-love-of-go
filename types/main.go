package main

import "fmt"

func main() {
	var title, author string
	var copies, edition int
	var promotion bool
	var specialOffer float64

	title = "All Quiet on the Western Front"
	author = "Remarque, Erich Maria"
	copies = 99
	edition = 1
	promotion = true
	specialOffer = 15

	fmt.Printf("%s: %s (%d. edition)\n", author, title, edition)
	fmt.Printf("Copies in stock: %d\n", copies)

	if promotion {
		fmt.Printf("Special Offer! %.2f%% off - buy now!\n", specialOffer)
	}
}
