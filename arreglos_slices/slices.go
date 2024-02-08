package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 52, 2, 1, 6, 4, 8, 6, 4}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]   //Slice creado desde un vector desde la posicion 3
	porcion2 := arreglo[:6]  //Slice creado de un vector desde la pocision 0 hasta la 6
	porcion3 := arreglo[6:7] //Slice creado desde un vector desde la posicion 6 a la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
