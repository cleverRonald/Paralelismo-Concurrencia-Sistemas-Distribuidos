// NOMBRE: CLEVER RONALD LAURA AYAMAMANI
package main
import (
    "fmt"
    "sync"
	"time"
	"math/rand"
)
var Coches_Total int = 10
var Coches_N_S int =0
var Coches_S_N int =0
var Pasando string = "vacio"
var NumCoche int = 1
type Puente struct {
    sync.Mutex
}
type Coche struct {
	Direc string
    Direccion_N_S, Direccion_S_N *Puente
}
func Pasar(p Coche){
    if p.Direc=="Norte-Sur"{
		if (Coches_N_S-1)==0{
			p.Direccion_N_S.Lock()
			fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
			NumCoche= NumCoche + 1
			Coches_N_S = Coches_N_S - 1
	   //	 p.Direccion_N_S.Unlock() //
			p.Direccion_S_N.Unlock()
		} else {
		    p.Direccion_N_S.Lock()
		    fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
		    NumCoche= NumCoche + 1
		    Coches_N_S = Coches_N_S - 1
		    p.Direccion_N_S.Unlock()
		}
   } else {
	   if (Coches_S_N-1)==0{
		   p.Direccion_S_N.Lock()
		   fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
			NumCoche= NumCoche + 1
		   Coches_S_N = Coches_S_N - 1
	   //	p.Direccion_S_N.Unlock()
		   p.Direccion_N_S.Unlock()
	   } else {
		  p.Direccion_S_N.Lock()
		  fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
			NumCoche= NumCoche + 1
		  Coches_S_N = Coches_S_N - 1
		  p.Direccion_S_N.Unlock()
	   }
   }
}
func (p Coche) Pasar_Puente( c chan *Coche,wg *sync.WaitGroup) {
	c <- &p
	if Pasando=="vacio"{
		Pasando=p.Direc
		if Pasando=="Norte-Sur"{
			p.Direccion_N_S.Lock()
			p.Direccion_S_N.Lock()
			fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
			NumCoche= NumCoche + 1
			Coches_N_S = Coches_N_S - 1
			p.Direccion_N_S.Unlock()
		} else {
			p.Direccion_S_N.Lock()
			p.Direccion_N_S.Lock()
			fmt.Print("\n  ",NumCoche,"° Coche en pasar el PUENTE en direccion de ")
			fmt.Println("(",p.Direc,")")
			NumCoche= NumCoche + 1
			Coches_S_N = Coches_S_N - 1
			p.Direccion_S_N.Unlock()
		}	
	}else {
		if p.Direc==Pasando {
			Pasar(p)
		} else {
			Pasar(p)
		}
	}
	wg.Done()
}

func host(c chan *Coche, wg *sync.WaitGroup) {
    for {
        if len(c) == 1 {
            <-c
            time.Sleep(400 * time.Millisecond)
        }
    }
}
func main() {
	var i int
    var wg sync.WaitGroup 
	c := make(chan *Coche, 1) 
	num_coches := Coches_Total
	Direc_N_S :=new(Puente)
	Direc_S_N :=new(Puente)

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	Coches := make([]*Coche, num_coches)
	for i = 0; i < num_coches; i++ {
		q:=r1.Intn(2)
		if q==0{
			Coches[i] = &Coche{"Norte-Sur",Direc_N_S ,Direc_S_N }
		    Coches_N_S= Coches_N_S + 1
			
		} else{
			Coches[i] = &Coche{"Sur-Norte",Direc_N_S ,Direc_S_N }
		    Coches_S_N= Coches_S_N + 1	
		}
	}

	fmt.Println("\nTotal de Coches en direccion de Norte-Sur: ",Coches_N_S)
	fmt.Println("Total de Coches en direccion de Sur-Norte: ",Coches_S_N)
	
	wg.Add(num_coches)
	go host(c, &wg) // bloqueo de canales
	
    for i = 0; i < num_coches; i++ {
		go Coches[i].Pasar_Puente(c, &wg)
    }
   wg.Wait()
}