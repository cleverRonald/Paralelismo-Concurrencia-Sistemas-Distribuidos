	//AUTOR: CLEVER RONALD LAURA AYAMAMANI
package main
import (
	"fmt"
	"time"
)
var numero int
func main() {
	for i := 1; i < 4;i++ {
		fmt.Println("INTERLEAVING ",i)
		go task(1)
  		go task(2)
		time.Sleep(time.Second)
		fmt.Println("")
	}
}

func task(x int){
	go instruction(x,1)
	go instruction(x,2)
	go instruction(x,3)
}
func instruction(x int, y int){
	fmt.Println(" Tarea ",x,", instruccion ",y)
	
}
