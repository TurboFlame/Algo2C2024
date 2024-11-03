package lista

type nodo[T any] struct {
	valor     T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

func crearNodo[T any](valor T) *nodo[T] {
	return &nodo[T]{
		valor:     valor,
		siguiente: nil,
	}
}
func (lista *listaEnlazada[T]) panicLista() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{
		primero: nil,
		largo:   0,
		ultimo:  nil,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}
func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevo := crearNodo[T](valor)
	nuevo.siguiente = lista.primero
	lista.primero = nuevo
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevo := crearNodo(valor)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.siguiente = nuevo
	}
	lista.ultimo = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.panicLista()
	borrado := lista.primero.valor
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return borrado
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.panicLista()
	return lista.primero.valor
}
func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.panicLista()
	return lista.ultimo.valor
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.valor) {
			break
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		lista:  lista,
		actual: lista.primero,
	}
}

func (iter *iteradorListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.valor
}
func (iter *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}
func (iter *iteradorListaEnlazada[T]) Siguiente() {
	if iter.HaySiguiente() {
		iter.anterior = iter.actual // Actualiza el nodo anterior
		iter.actual = iter.actual.siguiente
		return
	}
	panic("El iterador termino de iterar")
}
func (iter *iteradorListaEnlazada[T]) Insertar(valor T) {
	nuevo := crearNodo(valor)
	if iter.anterior == nil {
		nuevo.siguiente = iter.lista.primero
		iter.lista.primero = nuevo
	} else {
		nuevo.siguiente = iter.actual
		iter.anterior.siguiente = nuevo
	}
	if iter.actual == nil {
		iter.lista.ultimo = nuevo
	}

	iter.actual = nuevo
	iter.lista.largo++
}
func (iter *iteradorListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	borrado := iter.actual.valor
	if iter.actual == iter.lista.primero {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}

	if iter.actual == iter.lista.ultimo {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente
	iter.lista.largo--

	return borrado
}
