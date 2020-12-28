package main

import (
	"fmt"
)

func sliceFatia() {
	sabores := []string{"Frango com Catupiry", "Mussarela", "Calabresa", "Marguerita"}
	fatia := sabores[0:len(sabores)]
	fmt.Println(fatia)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}

func main() {
	times := []string{"Flamengo", "Vasco", "Botafogo", "Fluminense"}
	fmt.Println(times)
	sliceFatia()
}
