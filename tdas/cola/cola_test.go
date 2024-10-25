package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "EstaVacia no devuelve True sobre una cola vacia")
	require.Panics(t, func() { cola.VerPrimero() }, "La funcion no hace panic al ver el primero sobre una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "La funcion no hizo Panic con el mensaje especificado")
	require.Panics(t, func() { cola.Desencolar() }, "La funcion no hace panic al desencolar sobre una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "La funcion no hizo el panic con el mensaje especificado")

	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	require.False(t, cola.EstaVacia(), "EstaVacia devuelve True sobre una cola no vacia")
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()

	require.True(t, cola.EstaVacia(), "La cola no se comporta como recien creada al llenarla y vaciarla")
	require.Panics(t, func() { cola.VerPrimero() }, "La cola no se comporta como recien creada al llenarla y vaciarla")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "La cola no se comporta como recien creada al llenarla y vaciarla")

}

func TestFIFO(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	require.Equal(t, 1, cola.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola no esta vacia luego de desencolar todos sus elementos")
	cola2 := TDACola.CrearColaEnlazada[string]()
	cola2.Encolar("1")
	cola2.Encolar("2")
	cola2.Encolar("3")
	require.Equal(t, "1", cola2.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, "1", cola2.Desencolar())
	require.Equal(t, "2", cola2.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, "2", cola2.Desencolar())
	require.Equal(t, "3", cola2.VerPrimero(), "El primer valor de la cola no es el esperado")
	require.Equal(t, "3", cola2.Desencolar())
	require.True(t, cola2.EstaVacia(), "La cola no esta vacia luego de desencolar todos sus elementos")
	cola2.Encolar("uno")
	cola2.Encolar("dos")
	cola2.Encolar("tres")
	require.Equal(t, "uno", cola2.Desencolar(), "No retiene propiedades FIFO luego de vaciar y reencolar elementos")
	require.Equal(t, "dos", cola2.Desencolar(), "No retiene propiedades FIFO luego de vaciar y reencolar elementos")
	require.Equal(t, "tres", cola2.Desencolar(), "No retiene propiedades FIFO luego de vaciar y reencolar elementos")

}

func TestVolumen(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero(), "El primer valor no es el esperado en la prueba de volumen")
	}
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.Desencolar(), "No se desencolo el valor esperado en la prueba de volumen")
	}
	require.True(t, cola.EstaVacia(), "La cola no esta vacia al terminar la prueba de volumen")
}
