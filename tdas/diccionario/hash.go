package diccionario

import "fmt"

// ANOTACIONES
// Cambiar utilizados por borrados
type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
	CONDICION_INCREMENTO_REDIM = 0.7
	CONDICION_DECREMENTO_REDIM = 0.2
	RATIO_REDIM                = 2
	LARGO_INICIAL              = 10
)

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccionario := hash[V, K]{elementos: crearTabla[K, V](LARGO_INICIAL), cantidad: 0, borrados: 0}
	return &diccionario
}
func crearTabla[K comparable, V any](largo uint) []celda[V, K] {
	tabla := make([]celda[V, K], largo)
	return inicializarCeldas(tabla)
}
func inicializarCeldas[K comparable, V any](celdas []celda[V, K]) []celda[V, K] {
	for i := 0; i < len(celdas); i++ {
		celdas[i] = crearCelda[V, K]()
	}
	return celdas
}

type hash[V any, K comparable] struct {
	elementos []celda[V, K]
	cantidad  uint
	borrados  uint
}

// Busca el primer espacio ocupado con la misma clave o el primer espacio vacio.
// Si no lo encuentra, busca en el siguiente elemento del array.
// Devuelve una  celda ocupada si la  clave ya esta en el array, vacia si la clave no esta.
func (dic *hash[V, K]) buscar(claveHashed int, clave K) *celda[V, K] {

	estado := dic.elementos[claveHashed].estado
	if estado == VACIO || (clave == dic.elementos[claveHashed].clave && estado != BORRADO) {
		return &dic.elementos[claveHashed]
	} else {
		return dic.buscar((claveHashed+1)%(len(dic.elementos)), clave)
	}
}

func (dic *hash[V, K]) Guardar(clave K, dato V) {
	claveHashed := funcionHash[K](clave, len(dic.elementos))
	celda := dic.buscar(claveHashed, clave)

	//Si la clave no pertenecia al diccionario, aumenta la cantidad de elementos
	if celda.estado == VACIO {
		dic.cantidad++
	}
	celda.clave = clave
	celda.valor = dato
	celda.estado = OCUPADO

	//Llama a redimension si los elementos guardados y borrados estan por encima del umbral
	if float64(dic.cantidad+dic.borrados) > float64(len(dic.elementos))*CONDICION_INCREMENTO_REDIM {
		dic.redim(len(dic.elementos) * RATIO_REDIM)
	}

}

func (dic *hash[V, K]) Pertenece(clave K) bool {
	claveHashed := funcionHash[K](clave, len(dic.elementos))
	celda := dic.buscar(claveHashed, clave)
	return celda.estado == OCUPADO
}

func (dic *hash[V, K]) Obtener(clave K) V {
	claveHashed := funcionHash[K](clave, len(dic.elementos))
	celda := dic.buscar(claveHashed, clave)
	if celda.estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
	return celda.valor
}
func (dic *hash[V, K]) Borrar(clave K) V {
	claveHashed := funcionHash[K](clave, len(dic.elementos))
	celda := dic.buscar(claveHashed, clave)
	if celda.estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
	celda.estado = BORRADO
	dic.cantidad--
	dic.borrados++

	//Llama a redimension  si los elementos estan bajo el umbral
	if float64(dic.cantidad) < (float64(len(dic.elementos))*CONDICION_DECREMENTO_REDIM) && len(dic.elementos) < LARGO_INICIAL {
		dic.redim(len(dic.elementos) / RATIO_REDIM)
	}
	return celda.valor
}

func (dic *hash[V, K]) Cantidad() int {
	return int(dic.cantidad)
}

func (dic *hash[V, K]) redim(nuevoLargo int) {
	nuevoDic := hash[V, K]{elementos: crearTabla[K, V](uint(nuevoLargo)), cantidad: 0, borrados: 0}
	dic.Iterar(func(clave K, dato V) bool {
		nuevoDic.Guardar(clave, dato)
		return true
	})
	dic.elementos = nuevoDic.elementos
}

func (dic *hash[V, K]) Iterar(visitar func(clave K, dato V) bool) {
	contador := dic.buscarOcupado(0)
	condicion := true
	for contador < uint(len(dic.elementos)) && condicion {
		condicion = visitar(dic.elementos[contador].clave, dic.elementos[contador].valor)
		contador++
		contador = dic.buscarOcupado(contador)
	}
}
func (dic *hash[V, K]) Iterador() IterDiccionario[K, V] {
	iterador := iteradorHash[K, V]{dic, 0}
	iterador.indice = iterador.hash.buscarOcupado(iterador.indice)
	return &iterador
}

type iteradorHash[K comparable, V any] struct {
	hash   *hash[V, K]
	indice uint
}

// Revisa si el iterador no esta fuera de limite y si la celda actual tiene un valor valido
func (iterador *iteradorHash[K, V]) HaySiguiente() bool {
	return iterador.indice < uint(len(iterador.hash.elementos)) && iterador.hash.elementos[iterador.indice].estado == OCUPADO
}
func (iterador *iteradorHash[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.hash.elementos[iterador.indice].clave, iterador.hash.elementos[iterador.indice].valor
}

// Si el indice actual es un valor valido, busca el siguiente valor ocupado.
func (iterador *iteradorHash[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.indice++
	iterador.indice = iterador.hash.buscarOcupado(iterador.indice)
}

// Recibe el indice actual y devuelve el proximo indice en el cual
// existe  una  celda ocupada
func (dic *hash[V, K]) buscarOcupado(indiceActual uint) uint {
	for indiceActual < uint(len(dic.elementos)) && dic.elementos[indiceActual].estado != OCUPADO {
		indiceActual++
	}
	return indiceActual
}

func crearCelda[V any, K comparable]() celda[V, K] {
	celda := celda[V, K]{estado: VACIO}
	return celda
}

type celda[V any, K comparable] struct {
	valor  V
	clave  K
	estado estado
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// Se utilizo la funcion de hashing FNV1 A.
// Fuente: https://github.com/romain-jacotin/FNV-1a/blob/master/fnv1a.go
func funcionHash[K comparable](clave K, longitud int) int {
	inputdata := convertirABytes[K](clave)
	var val uint32 = 2166136261
	for _, v := range inputdata {
		val ^= uint32(v)
		val *= 16777619
	}
	return int(val) % longitud
}
