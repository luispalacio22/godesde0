package ejer_interfaces

import (
	"fmt"

	"github.com/luispalacio22/godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("soy un/a %s, y estoy respirando \n", hu.Sexo())
}
