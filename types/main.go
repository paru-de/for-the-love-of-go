package main

import "fmt"

func main() {
	var title, author string
	var copies, edition int
	title = "All Quiet on the Western Front"
	author = "Remarque, Erich Maria"
	copies = 99
	edition = 1

	fmt.Printf("%s: %s (%d. edition)\n", author, title, edition)
	fmt.Printf("Copies in stock: %d\n", copies)
}
