package main

import (
	"fmt"
	d "from-books/directory-scaner"
)

func main() {
	fmt.Println(d.Scan("/home/stas/Документы/books/go-en"))
	fmt.Println(d.Scan2("/home/stas/Документы/books/go-en"))
	//fmt.Println()
}
