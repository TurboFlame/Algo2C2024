package pila

/* Definición del struct pila proporcionado por la cátedra. */
const TAM_INICIAL int = 10
const VARIACION_REDIM int = 2
const RATIO_REDUCCION int = 4
const CANT_INICIAL int = 0

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := pilaDinamica[T]{datos: make([]T, TAM_INICIAL), cantidad: CANT_INICIAL}
	return &pila
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(dato T) {

	p.datos[p.cantidad] = dato
	p.cantidad++
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * VARIACION_REDIM)
	}
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	p.cantidad--
	if p.cantidad <= cap(p.datos)/RATIO_REDUCCION {
		p.redimensionar(cap(p.datos) / VARIACION_REDIM)
	}
	return p.datos[p.cantidad]
}

func (p *pilaDinamica[T]) redimensionar(nuevoLargo int) {
	nuevaPila := make([]T, nuevoLargo)
	copy(nuevaPila, p.datos)
	p.datos = nuevaPila
	return
}
