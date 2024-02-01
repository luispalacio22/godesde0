package main

import (
	"fmt"

	"github.com/luispalacio22/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvirtiendoTexto(1568)
	fmt.Println(estado)
	fmt.Println(texto)
}
