package main

import (
	"fmt"
)

func main() {
	/*var cores[5]string;

	  cores[0] = "Azul"
	  cores[1] = "Verde"
	  cores[2] = "Lar"*/

	cores := [...]string{"Azul", "Branco", "Verde"}

	fmt.Println(len(cores))

	fmt.Println("Primeira cor : ", cores[0])

}
