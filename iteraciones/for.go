package iteraciones

import (
	"fmt"
)

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

}

func Iterar1() {

	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
