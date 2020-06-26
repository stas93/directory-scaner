package paralel_clock

import (
	"fmt"
	"time"
)

// Беспонечные каналы
func l1() {
	start := make(chan int)
	cack := make(chan int)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second * 1)
			start <- i
		}
	}()
	go func() {
		for {
			el := <-start
			cack <- el * 2
		}
	}()
	for {
		el := <-cack
		fmt.Println(el)
	}
}

// Каналы с закрытием + строготипизированные, одни на запись другие на закрытие.
func l2() {
	start := make(chan int)
	cack := make(chan int)
	go func(start chan<- int) {
		for i := 0; i < 50; i++ {
			start <- i
		}
		close(start)
	}(start)
	go func(start <-chan int, cack chan<- int) {
		for el := range start {
			cack <- (el * 2)
		}
		close(cack)
	}(start, cack)
	for el := range cack {
		fmt.Println(el)
	}
}
func Listen() {
	l2()
}
