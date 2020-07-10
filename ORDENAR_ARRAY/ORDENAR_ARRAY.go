//// NAME: CLEVER RONALD LAURA AYAMAMANI
///------------------------------------------------------------
package main
import (
	"fmt"
	"sync"
	"sort"
)
func OrdenarSubArreglo(SubA [] int , c chan int,wg *sync.WaitGroup){
	sort.Ints(SubA)
	tmp := len(SubA)
	for i:=0; i<tmp;i++{
		c <- SubA[i]
	}
	close(c)
	wg.Done()
}
func main(){
	var wg sync.WaitGroup
	var numero int
	fmt.Print("Cantidad de datos que tendra el array :  ")
	fmt.Scan(&numero)
	fmt.Println(" ")
	var ArrayPrincipal [200]int
	for i :=0; i <numero;i++{
		var tem int
		fmt.Print("Dato ",i+1, " : ")
		fmt.Scan( &tem)
		ArrayPrincipal[i] = tem
	}
	fmt.Println(" ")
	fmt.Println(" Array Principal:		   ",ArrayPrincipal[:numero])
	var num int
	if ((numero+1)%4) ==0{
		num = (numero/4)+1
	} else {
		num = numero/4
	}
	SubArray_1 := ArrayPrincipal[:num]
	SubArray_2 := ArrayPrincipal[num:num*2]
	SubArray_3 := ArrayPrincipal[num*2:num*3]
	SubArray_4 := ArrayPrincipal[num*3:numero]
	fmt.Print(" Sub arrays del Array principal:	 ")
	fmt.Println(SubArray_1," ",SubArray_2," ",SubArray_3," ",SubArray_4)
	
	SA_1 := 	make(chan int, len(SubArray_1))
	SA_2 := 	make(chan int, len(SubArray_2))
	SA_3 := 	make(chan int, len(SubArray_3))
	SA_4 := 	make(chan int, len(SubArray_4))

	wg.Add(4)

	go OrdenarSubArreglo(SubArray_1,SA_1,&wg) 
	go OrdenarSubArreglo(SubArray_2,SA_2,&wg) 
	go OrdenarSubArreglo(SubArray_3,SA_3,&wg) 
	go OrdenarSubArreglo(SubArray_4,SA_4,&wg) 
	
	count :=0
	for i := range SA_1 {
		SubArray_1[count] = i	
		count++
	}
	count =0
	for i := range SA_2 {
		SubArray_2[count] = i	
		count++
	}
	count =0
	for i := range SA_3 {
		SubArray_3[count] = i	
		count++
	}
	count =0
	for i := range SA_4 {
		SubArray_4[count] = i	
		count++
	}
	fmt.Print(" Sub arrays Ordenados:			 ")
	fmt.Println(SubArray_1," ",SubArray_2," ",SubArray_3," ",SubArray_4)
	fmt.Println(" Sub arrays en el array principal:  ",ArrayPrincipal[:numero])
	sort.Ints(ArrayPrincipal[:numero])
	fmt.Println(" Array principal ordenado:	    ",ArrayPrincipal[:numero])
	wg.Wait()
}