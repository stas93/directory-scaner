package bit_vector

import "fmt"

type collection []uint64

func NewCollection() *collection {
	return &collection{}
}
func (c *collection) Has(x int) {

}
func (c *collection) Add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Println(word, bit)
}
