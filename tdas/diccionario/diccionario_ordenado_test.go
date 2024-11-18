package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
)

var TAMS_VOLUMEN_DICCIONARIO_ORDENADO = []int{1250, 2500, 5000, 10000, 20000, 40000}

//Se redujo el numero de iteraciones del ABB para su benchmark por la naturaleza mas lenta de este comparada al Hash

func compInt(a int, b int) int { return a - b }

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario ordenado vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}
func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementDiccionarioOrdenado(t *testing.T) {
	t.Log("Comprueba que Diccionario ordenado con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.Equal(t, 10, dic.Obtener("A"))
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}
func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario ordenado, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}
func TestReemplazoDatoDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
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
func TestReemplazoDatoHopscotchDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](compInt)
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
func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario ordenado, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

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
func TestConClavesNumericasDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compInt)
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
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}
func TestValorNuloDiccionarioOrdenado(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}
func TestGuardarYBorrarRepetidasVecesDiccionarioOrdenado(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. ")

	dic := TDADiccionario.CrearABB[int, int](compInt)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}
func TestIteradorInternoClavesDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cantidad := 0
	dic.Iterar(func(clave string, dato *int) bool {
		require.EqualValues(t, claves[cantidad], clave) //Requiere que las claves se iteren en el orden Gato-Perro-Vaca
		cantidad++
		return true
	})
	require.EqualValues(t, 3, cantidad, "El iterador no itero una vez por elemento")
}
func TestIteradorInternoValoresDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
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

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
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
func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves se iteren de forma ordenada y sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, claves[0], primero)

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.EqualValues(t, claves[1], segundo)
	require.EqualValues(t, valores[1], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, clave3, tercero)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorNoLlegaAlFinalDiccionarioOrdenado(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
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

	require.EqualValues(t, claves[0], primero)
	require.EqualValues(t, claves[1], segundo)
	require.EqualValues(t, claves[2], tercero)
}
func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario ordenado. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_DICCIONARIO_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenDiccionarioOrdenado(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenDiccionarioOrdenado(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
	}
	rand.Shuffle(len(valores), func(i, j int) { valores[i], valores[j] = valores[j], valores[i] })

	for _, i := range valores {
		dic.Guardar(fmt.Sprintf("%08d", i), i)
	}
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(fmt.Sprintf("%08d", valores[i]))
		if !ok {
			break
		}
		ok = dic.Obtener(fmt.Sprintf("%08d", valores[i])) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(fmt.Sprintf("%08d", valores[i])) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(fmt.Sprintf("%08d", valores[i]))
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func TestVolumenIteradorCorteDiccionarioOrdenado(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compInt)
	rango := 10000
	for i := 0; i < rango; i++ {
		intAleatorio := rand.Intn(rango)
		dic.Guardar(intAleatorio, intAleatorio)
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

func BenchmarkIteradorDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario ordenado. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas, asegurandose que esten ordenados. Se ejecuta cada prueba b.N " +
		"veces para generar un benchmark")
	for _, n := range TAMS_VOLUMEN_DICCIONARIO_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorDiccionarioOrdenado(b, n)
			}
		})
	}
}
func ejecutarPruebasVolumenIteradorDiccionarioOrdenado(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
	}
	rand.Shuffle(n, func(i, j int) { valores[i], valores[j] = valores[j], valores[i] })

	for i := 0; i < n; i++ {
		dic.Guardar(fmt.Sprintf("%08d", valores[i]), &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		require.True(b, c1 > clave, "El diccionario no itero de forma ordenada o contiene elementos iguales")
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}
func TestIterarRango(t *testing.T) {
	t.Log("Itera con un rango especificado. Testea que los elementos iterados esten dentro de ese rango")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	desde := 200
	hasta := 300
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	dic.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		require.True(t, clave >= desde, "El diccionario itero sobre un elemento menor al rango")
		require.True(t, clave <= hasta, "El diccionario itero sobre un elemento mayor al rango")
		return true
	})
}

func TestIterarRangoValoresNil(t *testing.T) {
	t.Log("Comprueba que al iterar rango con valores nulos se itere sobre todo el diccionario")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	contador := 0
	dic.IterarRango(nil, nil, func(clave int, valor int) bool {
		contador++
		return true
	})
	require.True(t, contador == 500, "No se itero una vez cada elemento del diccionario")
}
func TestIterarDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda muchos elementos en orden aleatorio y comprueba que se iteren de forma ordenada")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	for i := 0; i < 500; i++ {
		intRandom := rand.Intn(500)
		dic.Guardar(intRandom, intRandom)
	}
	anterior := 0
	dic.Iterar(func(clave int, valor int) bool {
		require.True(t, clave >= anterior)
		anterior = clave
		return true
	})
}
func TestIteradorEsOrdenado(t *testing.T) {
	t.Log("Comprueba que el iterador itere de manera ordenada el diccionario luego de ingresar claves en orden aleatorio")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	for i := 0; i < 500; i++ {
		intRandom := rand.Intn(500)
		dic.Guardar(intRandom, intRandom)
	}
	it := dic.Iterador()
	anterior := -1
	for it.HaySiguiente() {
		clave, _ := it.VerActual()
		require.True(t, clave > anterior, "El iterador itero una clave igual o mayor a la anterior")
		it.Siguiente()
	}
}
func TestIteradorRangoValoresNil(t *testing.T) {
	t.Log("Comprueba que un iterador con valores nulos itere sobre todo el diccionario")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	it := dic.IteradorRango(nil, nil)
	contador := 0
	for it.HaySiguiente() {
		contador++
		it.Siguiente()
	}
	require.True(t, contador == 500, "El iterador no itero una vez cada elemento del diccionario")
}
func TestIteradorRango(t *testing.T) {
	t.Log("Comprueba que un iterador con rango itere solo dentro del rango especificado")
	dic := TDADiccionario.CrearABB[int, int](compInt)
	desde := 200
	hasta := 300
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	it := dic.IteradorRango(&desde, &hasta)
	for it.HaySiguiente() {
		clave, _ := it.VerActual()
		require.True(t, clave >= desde, "El iterador itero un elemento menor al rango establecido")
		require.True(t, clave <= hasta, "El iterador itero un elemento mayor al rango establecido")
		it.Siguiente()
	}
}

func buscarClave(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}
