package cola_prioridad_test

import (
	"tdas/cola_prioridad"
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

func TestCrearHeap(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	if heap == nil {
		t.Errorf("Se esperaba un heap no nulo")
	}
	if !heap.EstaVacia() {
		t.Errorf("Se esperaba que el heap estuviera vacío")
	}
}

func TestEncolar(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Encolar(5)
	if heap.EstaVacia() {
		t.Errorf("Se esperaba que el heap no estuviera vacío")
	}
	if heap.Cantidad() != 1 {
		t.Errorf("El tamaño del heap esperado era 1, fue %d", heap.Cantidad())
	}
	if heap.VerMax() != 5 {
		t.Errorf("El elemento máximo debía ser 5, fue %d", heap.VerMax())
	}
}

func TestDesencolar(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Encolar(5)
	heap.Encolar(10)
	heap.Encolar(1)
	if max := heap.Desencolar(); max != 10 {
		t.Errorf("Se esperaba que el máximo elemento sea 10, fue %d", max)
	}
	if max := heap.Desencolar(); max != 5 {
		t.Errorf("Se esperaba que el máximo elemento fuera 5, fue %d", max)
	}
	if max := heap.Desencolar(); max != 1 {
		t.Errorf("Se esperaba que el máximo elemento fuera 1, fue %d", max)
	}
	if !heap.EstaVacia() {
		t.Errorf("Se esperaba que el heap este vacío")
	}
}

func TestHeapSort(t *testing.T) {
	elementos := []int{5, 1, 10, 3, 2}
	ordenado := cola_prioridad.HeapSort(elementos, cmpInt)
	esperado := []int{1, 2, 3, 5, 10}
	for i, v := range ordenado {
		if v != esperado[i] {
			t.Errorf("Se esperaba que ordenado[%d] sea %d, dio %d", i, esperado[i], v)
		}
	}
}

func TestCrearHeapArr(t *testing.T) {
	elementos := []int{5, 1, 10, 3, 2}
	heap := cola_prioridad.CrearHeapArr(elementos, cmpInt)
	if max := heap.VerMax(); max != 10 {
		t.Errorf("Se esperaba que el máximo elemento sea 10, fue %d", max)
	}
	if cantidad := heap.Cantidad(); cantidad != 5 {
		t.Errorf("Se esperaba que la cantidad sea 5, fue %d", cantidad)
	}
}
