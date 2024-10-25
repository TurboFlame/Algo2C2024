package cola

func CrearColaEnlazada[T any]() Cola[T] {
	cola := colaEnlazada[T]{primero: nil, ultimo: nil, tamanio: 0}
	return &cola
}

type colaEnlazada[T any] struct {
	primero nodoInterf[T]
	ultimo  nodoInterf[T]
	tamanio uint
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.verValor()
}
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodo(elem)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.setSiguiente(nuevoNodo)
	}
	cola.ultimo = nuevoNodo
	cola.tamanio++
	return
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := cola.primero.verValor()
	if cola.primero.verSiguiente() == nil {
		cola.primero, cola.ultimo = nil, nil
	} else {
		cola.primero = cola.primero.verSiguiente()
	}
	cola.tamanio--
	return valor
}

type nodo[T any] struct {
	valor     T
	siguiente nodoInterf[T]
}

func crearNodo[T any](elem T) nodoInterf[T] {

	return &nodo[T]{valor: elem, siguiente: nil}
}

func (nodo *nodo[T]) verValor() T {
	return nodo.valor
}
func (nodo *nodo[T]) setValor(elem T) {
	nodo.valor = elem
}
func (nodo *nodo[T]) verSiguiente() nodoInterf[T] {
	return nodo.siguiente
}
func (nodo *nodo[T]) setSiguiente(sigNodo nodoInterf[T]) {
	nodo.siguiente = sigNodo
}

type nodoInterf[T any] interface {
	verValor() T
	setValor(T)
	verSiguiente() nodoInterf[T]
	setSiguiente(nodoInterf[T])
}
