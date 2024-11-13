package cola_prioridad

const (
	CAPACIDAD          = 10
	FACTOR_REDIMENSION = 2
	LIMITE_REDIMENSION = 4
)

type fcmpHeap[T any] func(T, T) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, 0, CAPACIDAD), cantidad: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &heap[T]{datos: make([]T, len(arreglo)), cantidad: len(arreglo), cmp: funcion_cmp}
	copy(h.datos, arreglo)
	heapify(h.datos, h.cmp)
	return h
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(v T) {
	if h.cantidad == cap(h.datos) {
		h.redimensionar(cap(h.datos) * FACTOR_REDIMENSION)
	}
	if h.cantidad < len(h.datos) {
		h.datos[h.cantidad] = v
	} else {
		h.datos = append(h.datos, v)
	}
	h.cantidad++
	upHeap(h.datos, h.cantidad-1, h.cmp)
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
	downHeap(h.datos, 0, h.cantidad, h.cmp)
	if h.cantidad > 0 && h.cantidad == len(h.datos)/LIMITE_REDIMENSION {
		h.redimensionar(len(h.datos) / FACTOR_REDIMENSION)
	}
	return max
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) []T {
	cantidad := len(elementos)
	heapify(elementos, funcion_cmp)
	for i := cantidad - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, 0, i, funcion_cmp)
	}
	return elementos
}

// Auxiliares
func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downHeap(arr, i, len(arr), cmp)
	}
}

func upHeap[T any](arr []T, pos_hijo int, cmp func(T, T) int) {
	if pos_hijo == 0 {
		return
	}
	pos_padre := (pos_hijo - 1) / 2
	if cmp(arr[pos_padre], arr[pos_hijo]) < 0 {
		arr[pos_padre], arr[pos_hijo] = arr[pos_hijo], arr[pos_padre]
		upHeap(arr, pos_padre, cmp)
	}
}

func downHeap[T any](arr []T, pos_padre, n int, cmp func(T, T) int) {
	v := arr[pos_padre]
	for k := 2*pos_padre + 1; k < n; k = 2*k + 1 {
		if k+1 < n && cmp(arr[k], arr[k+1]) < 0 {
			k++
		}
		if cmp(v, arr[k]) >= 0 {
			break
		}
		arr[pos_padre] = arr[k]
		pos_padre = k
	}
	arr[pos_padre] = v
}

func (h *heap[T]) intercambiar(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

func (h *heap[T]) redimensionar(capacidad int) {
	nuevosDatos := make([]T, h.cantidad, capacidad)
	copy(nuevosDatos, h.datos)
	h.datos = nuevosDatos
}
