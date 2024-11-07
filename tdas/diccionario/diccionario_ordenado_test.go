package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
)

func TestDiccionarioOrdenadoPropios(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	dic.Guardar(5, 5)
	dic.Guardar(3, 3)
	require.True(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(3))
	require.EqualValues(t, 5, dic.Obtener(5))
	require.EqualValues(t, 5, dic.Obtener(5))
	require.EqualValues(t, 3, dic.Obtener(3))
	require.EqualValues(t, 5, dic.Borrar(5))
	require.False(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(3))
}
func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })

}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](func(a int, b int) int { return a - b })
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}
func TestUnElementDiccionarioOrdenado(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int { return strings.Compare(a, b) })
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.Equal(t, 10, dic.Obtener("A"))
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestReemplazoDatoHopscotchDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {

		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestReemplazoDatoDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosDiccionarioOrdenado(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestGuardarYBorrarRepetidasVecesDiccionarioOrdenado(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestValorNuloDiccionarioOrdenado(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](func(a string, b string) int { return strings.Compare(a, b) })
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestCadenaLargaParticularDiccionarioOrdenado(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}
func TestConClavesNumericasDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](func(a int, b int) int { return a - b })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}
func TestClaveVaciaDiccionarioOrdenado(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}
func TestIteradorInternoClavesDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](func(a string, b string) int { return strings.Compare(a, b) })
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarClave(cs[0], claves))
	require.NotEqualValues(t, -1, buscarClave(cs[1], claves))
	require.NotEqualValues(t, -1, buscarClave(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])

}
func buscarClave(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoValoresDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int { return strings.Compare(a, b) })

	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}
func TestIteradorInternoValoresConBorradosDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int { return strings.Compare(a, b) })

	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestVolumenIteradorCorteDiccionarioOrdenado(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })

	/* Inserta 'n' parejas en el funcionHash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false

			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda muchos elementos y comprueba que se iteren de forma ordenada")
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	anterior := 0
	dic.Iterar(func(clave int, valor int) bool {
		require.True(t, clave >= anterior)
		anterior = clave
		return true
	})
}

func TestIterarRango(t *testing.T) {
	t.Log("Itera con un rango especificado. Testea que los elementos iterados esten dentro de ese rango")
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	desde := 200
	hasta := 300
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	dic.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		require.True(t, clave >= desde, "El diccionario itero sobre un elemento menor al rango")
		require.True(t, clave <= hasta, "El diccionario itero sobre un elemento menor al rango")
		return true
	})
}
func TestIterarRangoNulo(t *testing.T) {
	t.Log("Revisa que al iterarse con valores nulos se itere una vez por elemento")
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	cantElementos := 500
	for i := 0; i < cantElementos; i++ {
		dic.Guardar(i, i)
	}
	contador := 0
	dic.IterarRango(nil, nil, func(clave int, valor int) bool {
		contador++
		return true
	})
	require.True(t, contador == cantElementos, "IteradorRango no itero una vez por elemento")
}

func TestIteradorNoLlegaAlFinalDiccionarioOrdenado(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int { return strings.Compare(a, b) })
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarClave(primero, claves))
	require.NotEqualValues(t, -1, buscarClave(segundo, claves))
	require.NotEqualValues(t, -1, buscarClave(tercero, claves))
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorDiccionarioOrdenado(t *testing.T) {
	t.Log("Crea un iterador, guarda muchos elementos y revisa que se recorran los elementos de forma ordenada y se itere cada uno una vez")
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })

	cantElementos := 500
	for i := 0; i < cantElementos; i++ {
		dic.Guardar(i, i)
	}
	it := dic.Iterador()
	anterior := 0
	contadorElementos := 0
	for it.HaySiguiente() {
		clave, _ := it.VerActual()
		require.True(t, clave >= anterior)
		anterior = clave
		contadorElementos++
		it.Siguiente()
	}
	require.EqualValues(t, cantElementos, contadorElementos)
}
func TestIteradorRangoDiccionarioOrdenado(t *testing.T) {
	t.Log("Crea un iterador, guarda muchos elementos y revisa que los elementos iterados esten dentro del rango ingresado")
	desde := 200
	hasta := 300
	cantElementos := 500
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })

	for i := 0; i < cantElementos; i++ {
		dic.Guardar(i, i)
	}

	it := dic.IteradorRango(&desde, &hasta)
	elementosIterados := 0
	for it.HaySiguiente() {
		clave, _ := it.VerActual()
		require.True(t, clave >= desde && clave <= hasta)
		elementosIterados++
		it.Siguiente()
	}
	require.EqualValues(t, elementosIterados, 101)

}

func TestIteradorInternoPruebaEspecifica(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 1; i <= 7; i++ {
		dic.Guardar(i, i)
	}
	desde := 2
	hasta := 5
	it := dic.IteradorRango(&desde, &hasta)
	contador := 0
	for it.HaySiguiente() {
		it.Siguiente()
		contador++
	}
	contador2 := 0
	contador3 := 2
	dic.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		require.True(t, clave == contador3)
		contador3++
		contador2++

		return true
	})
	require.EqualValues(t, 4, contador2)
	require.EqualValues(t, 4, contador2)
}
