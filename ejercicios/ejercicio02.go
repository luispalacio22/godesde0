package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var total int
var texto string

func Ejercicio2() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el numero")
	for {
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}

		}
	}

	for i := 1; i <= 10; i++ {
		total = numero1 * i
		texto += fmt.Sprintln(numero1, " x ", i, " = ", total)
	}
	return texto
}
