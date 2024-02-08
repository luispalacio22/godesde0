package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][20]int

func MuestroArregleo() {
	tabla[7] = 33
	tabla[2] = 23

	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][15] = 15
	fmt.Println(matriz)
}
