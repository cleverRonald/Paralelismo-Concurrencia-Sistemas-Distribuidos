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
		go task_multiplicar(3)
  		go task_multiplicar(5)
		time.Sleep(time.Second)
		fmt.Println("")
	}
}

func task_multiplicar(x int){
	go multiplicacion(x,1)
	go multiplicacion(x,2)
	go multiplicacion(x,3)
}
func multiplicacion(x int, y int){
	
	res := x*y
	fmt.Println(" ",x,"*",y,"=",res)
	
}
