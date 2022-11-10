package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

// Thread 1 -> main
func main() {
	// go task("A") // go routines // green threads
	// go task("B")
	// go task("C")
	// time.Sleep(15 * time.Second)

	canal := make(chan string)

	// Thread 2
	go func() {
		for {
			canal <- "Veio da Thread 2"
		}
	}()

	// Thread 1
	// msg := <-canal
	fmt.Println(<-canal)
}

// type Carro struct {
// 	Marca string // apontado para um lugar na memória
// 	Ano   int
// }

// func (c *Carro) MudaMarca(marca string) {
// 	c.Marca = marca // c é uma cópia do carro, está apontando para outro lugar na memória
// 	fmt.Println(c.Marca)
// }

// func main() {
// 	carro := Carro{Marca: "Fiat", Ano: 2010}
// 	carro.MudaMarca("Ford")
// 	fmt.Println(carro.Marca)
// }
