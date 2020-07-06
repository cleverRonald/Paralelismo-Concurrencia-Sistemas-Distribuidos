	//AUTOR: CLEVER RONALD LAURA AYAMAMANI
package main
import (
	"fmt"
	"time"
)
var numero int
var operacion [8]string
var count,t int 
func main() {
	fmt.Println("Ingrese un Numero")
	fmt.Scanf("%d\n", &numero)
	t = numero
	fmt.Println("")
	i := 1
	for i < 4 {
		count =1	
		fmt.Println("RACE CONDITTION ",i)
		go suma()
		go multiplicacion()
		go resta()
		go division()
		go exponencial()
		time.Sleep(time.Second)
		fmt.Println("El resultado es: ",numero)	
		i =i+1
		fmt.Println(" Orden de ejecucuion de las operaciones")
		for j:=1; j < 6; j++{
			fmt.Println(operacion[j])
		}
		numero = t
		fmt.Println("")
		
	}
}
func suma(){	
	operacion[count] = "  Suma" 
	count=count+1
	numero = numero +4
	return 
}
func resta(){
	
	operacion[count] = "  Resta" 
	count=count+1
	numero = numero -2
	return
}
func multiplicacion(){
	
	operacion[count] = "  Multiplicacion" 
	count=count+1
	numero = numero *2
	return
}
func division(){
	
	operacion[count] = "  Division" 
	count=count+1
	numero = numero /2
	return
}
func exponencial(){
	
	operacion[count] = "  Exponencial" 
	count=count+1
	numero = numero *numero
	return
}