package cola_prioridad

type fcmpHeap[K any] func(K, K) int

type heap[T comparable] struct {
    datos    []T
    cantidad int
    cmp      fcmpHeap[T]
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
    return &heap[T]{datos: make([]T, 0), cantidad: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &heap[T]{datos: make([]T, len(arreglo)), cantidad: len(arreglo), cmp: funcion_cmp}
	copy(h.datos, arreglo)
	for i := h.cantidad/2 - 1; i >= 0; i-- {
		h.hundir(i)
	}
	return h
}

func (h *heap[T]) EstaVacia() bool {
    return h.cantidad == 0
}

func (h *heap[T]) Encolar(v T) {
    h.datos = append(h.datos, v)
    h.cantidad++
    h.darSoporte(h.cantidad - 1, v)
}

func (h *heap[T]) VerMax() T {
    if h.EstaVacia() {
        panic("La cola esta vacia")
    }
    return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
    if h.EstaVacia() {
        panic("La cola esta vacia")
    }
    max := h.datos[0]
    h.intercambiar(0, h.cantidad-1)
    h.cantidad--
    h.hundir(0)
    return max
}

func (h *heap[T]) Cantidad() int {
    return h.cantidad
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) []T {
    h := CrearHeapArr(elementos, funcion_cmp).(*heap[T])
    for i := h.cantidad - 1; i > 0; i-- {
        h.intercambiar(0, i)
        h.cantidad--
        h.hundir(0)
    }
    return h.datos
}

//Auxiliares 
func (h *heap[T]) darSoporte(i int, v T) {
    for i > 0 && h.cmp(h.datos[(i-1)/2], v) < 0 {
        h.datos[i] = h.datos[(i-1)/2]
        i = (i - 1) / 2
    }
    h.datos[i] = v
}

func (h *heap[T]) hundir(i int) {
    v := h.datos[i]
    for k := 2*i + 1; k < h.cantidad; k = 2*k + 1 {
        if k+1 < h.cantidad && h.cmp(h.datos[k], h.datos[k+1]) < 0 {
            k++
        }
        if h.cmp(v, h.datos[k]) >= 0 {
            break
        }
        h.datos[i] = h.datos[k]
        i = k
    }
    h.datos[i] = v
}

func (h *heap[T]) intercambiar(i, j int) {
    h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}
