package main

import (
	"fmt"
	d "from-books/directory-scaner"
)

func main() {
	parse := "/home/stas/Документы"

	fmt.Println(d.Scan(parse))
	fmt.Println(d.Scan2(parse))
	fmt.Println(d.Scan3(parse))
}
