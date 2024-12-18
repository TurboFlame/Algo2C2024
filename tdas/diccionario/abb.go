package diccionario

import "tdas/pila"

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

func CrearABB[K comparable, V any](funcionCmp funcComparacion[K]) DiccionarioOrdenado[K, V] {
	nuevoAbb := abb[K, V]{cantidad: 0, comp: funcionCmp, raiz: nil}
	return &nuevoAbb
}
func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}
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
		*nodo = crearNodoABB[K, V](clave, dato)
	} else {
		(*nodo).clave, (*nodo).dato = clave, dato
	}
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
		nodoMenor := (*nodo).izq.buscarMaximo()
		dic.Borrar(nodoMenor.clave)
		dic.cantidad++
		(*nodo).clave, (*nodo).dato = nodoMenor.clave, nodoMenor.dato
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
func (abb *nodoAbb[K, V]) buscarMinimo() *nodoAbb[K, V] {
	if abb == nil {
	}
	if abb.izq != nil {
		return abb.izq.buscarMinimo()
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
	dic.raiz.iterarArbol(dic.comp, nil, nil, visitar)

}
func (abb *nodoAbb[K, V]) iterarArbol(comp funcComparacion[K], min *K, max *K, visitar func(clave K, dato V) bool) bool {
	if abb == nil {
		return true
	}
	if min != nil && comp(abb.clave, *min) < 0 {
		return abb.der.iterarArbol(comp, min, max, visitar)
	} //Si la clave es menor al rango, se desplaza a la derecha
	if max != nil && comp(abb.clave, *max) > 0 {
		return abb.izq.iterarArbol(comp, min, max, visitar)
	} //Si la clave es mayor al rango, se desplaza a la izquierda
	return abb.izq.iterarArbol(comp, min, max, visitar) && visitar(abb.clave, abb.dato) && abb.der.iterarArbol(comp, min, max, visitar) //Recorrido inorder que se detiene al encontrar un false
}

func (dic *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	dic.raiz.iterarArbol(dic.comp, desde, hasta, visitar)
}

func (dic *abb[K, V]) Iterador() IterDiccionario[K, V] { return dic.inicializarIterador(nil, nil) }

type iteradorAbb[K comparable, V any] struct {
	pila  pila.Pila[nodoAbb[K, V]]
	dic   *abb[K, V]
	desde *K
	hasta *K
}

func (it *iteradorAbb[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return it.pila.VerTope().clave, it.pila.VerTope().dato
}
func (it *iteradorAbb[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia() && (it.hasta == nil || it.dic.comp(it.pila.VerTope().clave, *it.hasta) <= 0)
}
func (it *iteradorAbb[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := it.pila.Desapilar()
	if desapilado.der != nil {
		desapilado.der.apilarIzquierdos(it.pila, it.dic.comp, it.desde, it.hasta)
	}
}
func (abb *nodoAbb[K, V]) apilarIzquierdos(pila pila.Pila[nodoAbb[K, V]], comp funcComparacion[K], min *K, max *K) {
	if abb == nil {
		return
	}

	if (min == nil || comp(abb.clave, *min) >= 0) && (max == nil || comp(abb.clave, *max) <= 0) { //Revisa si esta dentro del rango. Si es asi, apila y llama recursivamente a izquierdos
		pila.Apilar(*abb)
		abb.izq.apilarIzquierdos(pila, comp, min, max)
	} else if min == nil || comp(abb.clave, *min) >= 0 { //Si el elemento es mayor al rango, no lo apila pero llama a apilar sus izquierdos
		abb.izq.apilarIzquierdos(pila, comp, min, max)
	} else if abb.der != nil { //Si ninguna se cumple, llama a apilarIzquierdos en el hijo derecho
		abb.der.apilarIzquierdos(pila, comp, min, max)
	}
}

func (dic *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return dic.inicializarIterador(desde, hasta)
}

func (dic *abb[K, V]) inicializarIterador(desde *K, hasta *K) IterDiccionario[K, V] {
	it := iteradorAbb[K, V]{pila: pila.CrearPilaDinamica[nodoAbb[K, V]](), dic: dic, desde: desde, hasta: hasta}
	it.dic.raiz.apilarIzquierdos(it.pila, it.dic.comp, it.desde, it.hasta)
	return &it
}
