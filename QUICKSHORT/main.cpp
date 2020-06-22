
/**Name: CLEVER RONALD LAURA AYAMAMANI **/

#include <stdio.h>
#include <time.h>
#include <omp.h>
#include <iostream>
using namespace std;
int thread_count = omp_get_num_procs();

///QUICKSHORT DE MODO SECUENCIAL ------------------------------------------------------------------------

void QUICKSHORT_SECUENCIAL(int *A, int izq, int der) {
 int piv = A[izq];
 int i = izq;
 int j = der;
 int aux;

     while (i < j) {
          while (A[i] <= piv && i < j) i++;
          while (A[j] > piv) j--;
          if (i < j) {
               aux  = A[i];
               A[i] = A[j];
               A[j] = aux;
          }
     }
     A[izq] = A[j];
     A[j] = piv;
     if (izq<j - 1)
        QUICKSHORT_SECUENCIAL(A, izq, j - 1);
     if (j + 1 <der)
        QUICKSHORT_SECUENCIAL(A, j + 1, der);
}

///QUICKSHORT DE MODO PARALELA ------------------------------------------------------------------------

void QUICKSHORT_PARALELO(int *A, int izq, int der) {
 int piv = A[izq];
 int i = izq;
 int j = der;
 int aux;

    while (i < j) {
        while (A[i] <= piv && i < j) i++;
        while (A[j] > piv) j--;
        if (i < j) {
            aux  = A[i];
            A[i] = A[j];
            A[j] = aux;
        }
    }
     A[izq] = A[j];
     A[j] = piv;

    if (izq<j - 1){
                #pragma omp task
                    QUICKSHORT_PARALELO(A, izq, j - 1);
         }
    if (j + 1 <der){
                #pragma omp task
                    QUICKSHORT_PARALELO(A, j + 1, der);
        }
}
void NumeroDeDatos (int n)
{
    int i;
    int A[n];
    int B[n];
    for (int i = 0; i < n; i++) {
        A[i] = rand();
        B[i] = rand();
    }

    clock_t t_ini, t_fin;
    double TIEMPO_SECUENCIAL,TIEMPO_PARALELO;

    ///SECUENCIAL -------------------------------------------------------------------------

    t_ini = clock();
    QUICKSHORT_SECUENCIAL(A,0,n-1);
    t_fin = clock();

    TIEMPO_SECUENCIAL = (double)(t_fin - t_ini) / CLOCKS_PER_SEC;

    ///PARALELA -------------------------------------------------------------------------

    t_ini = clock();
    #pragma omp parallel num_threads(thread_count)
	{
		#pragma omp single nowait
		{
            QUICKSHORT_PARALELO(B, 0, n-1);
        }
	}
    t_fin = clock();

    TIEMPO_PARALELO= (double)(t_fin - t_ini) / CLOCKS_PER_SEC;

    ///-------------------------------------------------------------------------------------------------------

    cout<<"Tiempo que requirie para el ordenamiento QUICKSHORT en " <<endl ;
    cout<<endl ;
    cout<<"modo secuencial es: "<<TIEMPO_SECUENCIAL*1000.0<<" milisegundos"<<endl;
    cout<<"modo paralelo es: "<<TIEMPO_PARALELO*1000.0<<" milisegundos"<<endl;
    cout<<endl ;
    cout<<"En una arreglo de :"<<n<<" numeros"<<endl ;
    cout<<endl ;
    cout<<"La diferencia de tiempo entre el modo secuencial y paralelos es: " <<(TIEMPO_PARALELO-TIEMPO_SECUENCIAL)*1000.0<<" milisegundos"<<endl ;
}

int main() {

    NumeroDeDatos(20000);


 return 0;
}
