package diccionario

//ANOTACIONES
//Considerar switch en funcion de busqueda
//Crear una funcion Swap que cambie el nodo al que apunta un  puntero puntero
//Cambiar buscar para que reciba un puntero a puntero

type funcComparacion[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	comp     funcComparacion[K]
}

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	nuevoAbb := abb[K, V]{cantidad: 0, comp: funcionCmp, raiz: nil}
	return &nuevoAbb
}
func crearNodoABB[K comparable, V any]() *nodoAbb[K, V] {

	return &nodoAbb[K, V]{}
}

// func(ab *nodoAbb[K,V]) cambiarHijo(valorAbb **nodoAbb[K,V])  {}
/*func (ab *nodoAbb[K, V]) buscar(clave K, comp funcComparacion[K]) **nodoAbb[K, V] {
	if ab == nil || comp(clave, ab.clave) == 0 {
		println(&ab)
		return &ab
	} else if comp(clave, ab.clave) < 0 {
		return ab.izq.buscar(clave, comp)
	} else {
		return ab.der.buscar(clave, comp)
	}
}*/

func buscar[K comparable, V any](ab **nodoAbb[K, V], comp funcComparacion[K], clave K) **nodoAbb[K, V] {
	if *ab == nil || comp(clave, (*ab).clave) == 0 {
		return ab
	} else if comp(clave, (*ab).clave) < 0 {
		return buscar(&((*ab).izq), comp, clave)
	} else {
		return buscar(&((*ab).der), comp, clave)
	}
}

func (dic *abb[K, V]) Guardar(clave K, dato V) {

	nodo := buscar[K, V](&dic.raiz, dic.comp, clave)
	if *nodo == nil {
		dic.cantidad++
		*nodo = crearNodoABB[K, V]()
	}
	(*nodo).dato, (*nodo).clave = dato, clave
}

func (dic *abb[K, V]) Borrar(clave K) V {
	nodo := buscar[K, V](&dic.raiz, dic.comp, clave)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*nodo).dato
	if (*nodo).izq == nil && (*nodo).der == nil {
		*nodo = nil
	} else if (*nodo).izq != nil && (*nodo).der == nil {
		*nodo = (*nodo).izq
	} else if (*nodo).der != nil && (*nodo).izq == nil {
		*nodo = (*nodo).der
	} else if (*nodo).izq != nil && (*nodo).der != nil {
		nodoPrevio := (*nodo).izq.buscarMaximo()
		dic.Borrar(nodoPrevio.clave)
		(*nodo).clave, (*nodo).dato = nodoPrevio.clave, nodoPrevio.dato
	}
	dic.cantidad--
	return dato

}
func (abb *nodoAbb[K, V]) buscarMaximo() *nodoAbb[K, V] {
	if abb == nil {
		return nil
	}
	if abb.der != nil {
		return abb.der.buscarMaximo()
	}
	return abb
}

func (dic *abb[K, V]) Pertenece(clave K) bool {
	nodo := buscar[K, V](&dic.raiz, dic.comp, clave)
	return *nodo != nil
}
func (dic *abb[K, V]) Obtener(clave K) V {
	nodo := buscar[K, V](&dic.raiz, dic.comp, clave)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}
func (dic *abb[K, V]) Cantidad() int {
	return dic.cantidad
}
func (dic *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {

}
func (dic *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return nil
}
func (dic *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
}
func (dic *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] { return nil }
