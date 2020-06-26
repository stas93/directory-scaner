package main

import (
	"fmt"
	d "from-books/directory-scaner"
)

func main() {

	fmt.Println(d.Scan("/home/stas"))
	fmt.Println(d.Scan2("/home/stas"))
	//fmt.Println()
}
