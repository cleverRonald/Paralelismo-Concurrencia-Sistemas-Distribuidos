
/**Name: CLEVER RONALD LAURA AYAMAMANI **/

#include <stdio.h>
#include <omp.h>
#include <time.h>
#include <iostream>

using namespace std;

void DimencionDeMatrices(int n)
{

    int i,j;
    int matriz_a[n][n];
    int matriz_b[n][n];
    int matriz_resp[n][n];

    clock_t t_ini, t_fin;
    double milisegundos_multiplicacion_secuancial;
    double milisegundos_multiplicacion_paralela;

    ///LLENADO DE MATRICES CON DATOS ENTEROS RANDOWN-----------------------------------------------------

    for (i=0;i<n;i++)
    {
        for (j=0;j<n;j++)
        {   matriz_a[i][j]=rand();
            matriz_b[i][j]=rand();
        }
    }

    ///MULTIPLICACION DE MATRICES EN MODO SECUENCIAL--------------------------------------------------------

    t_ini = clock();
        for(i=0 ; i < n; i++)
            {
                for(j=0 ; j < n; j++)
                {
                    matriz_resp[i][j]=0;
                    for(int k = 0; k < n; k++)
                    {
                        matriz_resp[i][j]=matriz_resp[i][j]+ (matriz_a[i][k] * matriz_b[k][j]);
                    }
                }
            }
     t_fin = clock();
    milisegundos_multiplicacion_secuancial = (double)(t_fin - t_ini) / CLOCKS_PER_SEC;


    ///MULTIPLICACION DE MATRICES EN MODO PARALELA-------------------------------------------------------------

     t_ini = clock();
        #pragma omp parallel
        {
            #pragma omp for
            for(i=0 ; i < n; i++)
            {
                for(j=0 ; j < n; j++)
                {
                    matriz_resp[i][j]=0;
                    for( int k = 0; k < n; k++)
                    {
                        matriz_resp[i][j]=matriz_resp[i][j]+ (matriz_a[i][k] * matriz_b[k][j]);
                    }
                }
            }
        }

    t_fin = clock();
    milisegundos_multiplicacion_paralela = (double)(t_fin - t_ini) / CLOCKS_PER_SEC;

    ///-------------------------------------------------------------------------------------------------------


    cout<<"Tiempo que requirie para la multiplicacion de matrices en " <<endl ;
    cout<<endl ;
    cout<<"modo secuencial es: "<<milisegundos_multiplicacion_secuancial*1000.0<<" milisegundos"<<endl;
    cout<<"modo paralelo es: "<<milisegundos_multiplicacion_paralela*1000.0<<" milisegundos"<<endl;
    cout<<endl ;
    cout<<"En una dimension de matrices de :"<<n<<"x"<<n<<endl ;
    cout<<endl ;
    cout<<"La diferencia de tiempo entre el modo secuencial y paralelos es: " <<(milisegundos_multiplicacion_secuancial-milisegundos_multiplicacion_paralela)*1000.0<<" milisegundos"<<endl ;
}
int main()
{
    DimencionDeMatrices(414);
    return 0;
}
