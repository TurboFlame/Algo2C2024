package lista

type Lista[T any] interface {
	//EstaVacia devuelve true si la lista no tiene nodos, false en caso contrario
	EstaVacia() bool
	//InsertarPrimero inserta un elemento al comienzo de la lista
	InsertarPrimero(T)
	//InsertarUltimo inserta un elemento en el último lugar de la lista
	InsertarUltimo(T)
	//BorrarPrimero elimina el primer elemento de la lista y lo devuelve
	//Precondición: que la lista no este vacía
	BorrarPrimero() T
	//VerPrimero devuelve el primer elemento de la lista
	//Precondición: que la lista no esté vacía
	VerPrimero() T
	//VerUltimo devuelve el ultimo elemento de la lista
	//Precondición: que la lista no esté vacía
	VerUltimo() T
	//Largo devuelve el número de elementos de la lista
	Largo() int
	//Iterar recorre la lista, aplicando la función visitar a cada uno de sus elementos
	//Hasta que la lista termine o la función visitar devuelva false
	Iterar(visitar func(T) bool)
	//Iterador devuelve un iterador externo para la lista, permite recorrerla y modificarla
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//VerActual devuelve el valor actual del nodo al que apunta el iterador
	//Precondición: que no haya llegado al ultimo elemento de la lista
	VerActual() T
	//HaySiguiente devuelve true si el iterador tiene un siguiente elemento, false en caso contrario
	HaySiguiente() bool
	//Siguiente avanza al iterador al siguiente elemento de la lista
	//Precondición: que no se encuentre en el último elemento
	Siguiente()
	//Insertar agrega un nuevo elemento en la posición actual
	Insertar(T)
	//Borrar borra el elemento actual y devuelve el elemento actual al que apunta el iterador
	//Precondición: que no se encuentre en el último elemento
	Borrar() T
}
