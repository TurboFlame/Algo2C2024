package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestLIFO(t *testing.T) {
	//Comprobacion de LIFO
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia(), "La pila no esta vacia luego de desapilar sus elementos")
	//Comprobacion que al vaciar y volver a llenar mantiene las propiedades LIFO
	pila.Apilar(16)
	pila.Apilar(32)
	pila.Apilar(8)
	require.Equal(t, 8, pila.Desapilar(), "La pila no mantiene sus propiedades FIFO luego de vaciar y desapilar sus elementos")
	require.Equal(t, 32, pila.Desapilar(), "La pila no mantiene sus propiedades FIFO luego de vaciar y desapilar sus elementos")
	require.Equal(t, 16, pila.Desapilar(), "La pila no mantiene sus propiedades FIFO luego de vaciar y desapilar sus elementos")
	require.True(t, pila.EstaVacia(), "La pila no esta vacia luego de desapilar sus elementos")
	//Comprobacion LIFO con strings
	pila2 := TDAPila.CrearPilaDinamica[string]()
	pila2.Apilar("1")
	pila2.Apilar("2")
	pila2.Apilar("3")
	require.Equal(t, "3", pila2.Desapilar())
	require.Equal(t, "2", pila2.Desapilar())
	require.Equal(t, "1", pila2.Desapilar())
	require.True(t, pila2.EstaVacia(), "La pila no esta vacia luego de desapilar sus elementos")
}

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.Desapilar() }, "Al desapilar una pila recien creada, no tira panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Al desapilar una pila recien creada, no tira el panic con el mensaje de error correcto")
	require.Panics(t, func() { pila.VerTope() }, "Al ver el tope de una pila recien creada, no tira panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Al ver el tope de una pila recien creada, no tira panic")
	require.True(t, pila.EstaVacia(), "La funcion no devuelve True sobre una pila vacia")

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.False(t, pila.EstaVacia(), "EstaVacia devuelve True sobre una pila recien creada")
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.Panics(t, func() { pila.Desapilar() }, "Al desapilar una pila vaciada, no tira panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Al desapilar una pila vaciada, no tira el panic con el mensaje de error correcto")
	require.Panics(t, func() { pila.VerTope() }, "Al ver el tope de una pila vaciada, no tira panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Al ver el tope de una pila vaciada, no tira panic")
	require.True(t, pila.EstaVacia(), "La funcion no devuelve True sobre una pila vaciada")

}

func TestPruebaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= 10000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "No se apilo el valor esperado en la prueba de volumen")
	}
	for i := 10000; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar(), "No se desapilo el valor esperado en la prueba de volumen")
	}
	require.True(t, pila.EstaVacia(), "La pila no esta vacia luego de la prueba de volumen")
	require.Panics(t, func() { pila.Desapilar() }, "Al desapilar la funcion vacia luego de la prueba de volumen, la funcion no tira panic")
	require.Panics(t, func() { pila.VerTope() }, "Al ver el tope de la pila luego de apilar y desapilar en la prueba de volumen, la funcion no tira panic")

}
