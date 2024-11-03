package cola_prioridad

type fcmpHeap[K any] func(K, K) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, 1), cantidad: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &heap[T]{datos: append([]T{nil}, arreglo...), cantidad: len(arreglo), cmp: funcion_cmp}
	for i := h.cantidad / 2; i > 0; i-- {
		h.hundir(i)
	}
	return h
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(v T) {
	h.cantidad++
	h.darSoporte(h.cantidad, v)
}	

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.dato(1)
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := h.dato(1)
	h.intercambiar(1, h.cantidad)
	h.cantidad--
	h.hundir(1)
	return max
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) []T {
	h := CrearHeapArr(elementos, funcion_cmp)
	for i := h.cantidad; i > 1; i-- {
		h.intercambiar(1, i)
		h.cantidad--
		h.hundir(1)
	}
	return h.datos[1:]
}

func (h *heap[T]) dato(i int) T {
	return h.datos[i]
}

func (h *heap[T]) darSoporte(i int, v T) {
	h.datos = append(h.datos, v)
	for i > 1 && h.cmp(h.dato(i/2), v) < 0 {
		h.datos[i] = h.dato(i/2)
		i /= 2
	}
	h.datos[i] = v
}

func (h *heap[T]) hundir(i int) {
	v := h.dato(i)
	for k := 2 * i; k <= h.cantidad; k *= 2 {
		if k < h.cantidad && h.cmp(h.dato(k), h.dato(k+1)) < 0 {
			k++
		}
		if h.cmp(v, h.dato(k)) >= 0 {
			break
		}
		h.datos[i] = h.dato(k)
		i = k
	}
	h.datos[i] = v
}

func (h *heap[T]) intercambiar(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}


