package main

import (
	"fmt"
	"runtime"

	"github.com/luispalacio22/godesde0/ejercicios"
	"github.com/luispalacio22/godesde0/files"
	"github.com/luispalacio22/godesde0/iteraciones"
	"github.com/luispalacio22/godesde0/teclado"
	"github.com/luispalacio22/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvirtiendoTexto(1568)
	fmt.Println(estado)
	fmt.Println(texto)
	os := runtime.GOOS
	if os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es: ", os)
	} else {
		fmt.Println("Esto es windows, es: ", os)
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("esto es linux")
	case "windows":
		fmt.Println("esto es windows")
	default:
		fmt.Println("esto es: ", os)
	}

	numero, text := ejercicios.EjercicioInd("1512")
	fmt.Println(numero)
	fmt.Println(text)

	teclado.IngresoNumeros()

	iteraciones.Iterar()
	fmt.Println(ejercicios.Ejercicio2())
	files.GrabaTabla()
	files.SumaTabla()
	files.LeoArchivo()
	files.LeoArchivo2()
}
