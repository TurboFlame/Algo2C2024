package lista_test

import (
	"tdas/lista"
	"testing"
)

func TestEstaVacia(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	if !lista.EstaVacia() {
		t.Errorf("Una lista nueva deberia ser reconocida como vacia")
	}
}
func TestInsertarPrimero(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)

	if lista.VerPrimero() != 2 {
		t.Errorf("Esperado 2, pero se obtuvo %v", lista.VerPrimero())
	}
}

func TestInsertarUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	if lista.VerUltimo() != 2 {
		t.Errorf("Esperado 2, pero se obtuvo %v", lista.VerUltimo())
	}
}

func TestBorrarPrimero(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	borrado := lista.BorrarPrimero()
	if borrado != 2 {
		t.Errorf("Esperaba un 2 como elemento borrado, se obtiene %v", borrado)
	}
	if lista.VerPrimero() != 1 {
		t.Errorf("Esperaba un 1 como primer elemento, se obtiene %v", lista.VerPrimero())
	}
}

func TestIterInsertarPrimero(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	iter := lista.Iterador()
	iter.Insertar(1)
	if lista.VerPrimero() != 1 {
		t.Errorf("Se esperaba que el primer Valor sea el insertado, es %v", lista.VerPrimero())
	}
}

func TestIterInsertarUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)
	if lista.VerUltimo() != 2 {
		t.Errorf("Se esperaba que el ultimo elemento fuera 2, es %v", lista.VerUltimo())
	}
}

func TestIterInsertarMedio(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(5)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(4)
	if lista.VerPrimero() != 3 {
		t.Errorf("Se esperaba que el primer Valor sea 3, es %v", lista.VerPrimero())
	}
	if lista.VerUltimo() != 5 {
		t.Errorf("Se esperaba que el ultimo Valor sea 5, es %v", lista.VerUltimo())
	}
}

func TestIterBorrarPrimero(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	iter := lista.Iterador()
	borrado := iter.Borrar()
	if borrado != 2 {
		t.Errorf("Se esperaba que el elemento borrado sea 2, es %v", borrado)
	}
	if lista.VerPrimero() != 1 {
		t.Errorf("Se esperaba que el primer elemento sea 1, es %v", lista.VerPrimero())
	}
}
func TestIterBorrarUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	iter.Siguiente()
	borrado := iter.Borrar()
	if borrado != 2 {
		t.Errorf("Se esperaba que el elemento borrado sea 2, es %v", borrado)
	}
	if lista.VerUltimo() != 1 {
		t.Errorf("Se esperaba que el último elemento sea 1, es %v", lista.VerUltimo())
	}
}
func TestIterBorrarMedio(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	iter := lista.Iterador()
	iter.Siguiente()
	borrado := iter.Borrar()
	if borrado != 4 {
		t.Errorf("Se esperaba que el elemento borrado sea 4, es %v", borrado)
	}
	if iter.VerActual() == 4 {
		t.Errorf("Se esperaba que no haya un 4")
	}
}

func TestIterConCorte(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	suma := 0
	lista.Iterar(func(Valor int) bool {
		suma += Valor
		if Valor == 3 {
			return false
		}
		return true
	})

	if suma != 6 {
		t.Errorf("Se esperaba que la suma sea 6, pero fue %v", suma)
	}
}

func TestBorrarUnicoElemento(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)

	iter := lista.Iterador()
	borrado := iter.Borrar()

	if borrado != 1 {
		t.Errorf("Se esperaba que se borre el 1, pero se borro %v", borrado)
	}

	if lista.EstaVacia() == false {
		t.Errorf("La lista deberia estar vacia")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Se esperaba un panic al intentar iterar una lista vacia")
		}
	}()

	iter.Siguiente() //prueba de panic
}

func TestBorrarTodosLosElementos(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	for i := 0; i < 3; i++ {
		iter.Borrar()
	}

	if lista.EstaVacia() == false {
		t.Errorf("Se esperaba que la lista estuviera vacía")
	}
}

func TestAvanzarDespuesDeBorrarUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Borrar()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Se esperaba un pánico al intentar avanzar después de borrar el último elemento")
		}
	}()

	iter.Siguiente()
}

func TestVolumenInsertarUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	cantidad := 10000

	for i := 0; i < cantidad; i++ {
		lista.InsertarUltimo(i)
	}

	if lista.Largo() != cantidad {
		t.Errorf("Se esperaban %d elementos, pero la lista tiene %d", cantidad, lista.Largo())
	}

	ultimo := lista.VerUltimo()
	if ultimo != cantidad-1 {
		t.Errorf("Se esperaba %d como último elemento, pero se encontró %d", cantidad-1, ultimo)
	}
}

func TestVolumenVerUltimo(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	cantidad := 10000

	for i := 0; i < cantidad; i++ {
		lista.InsertarUltimo(i)
	}

	for i := 0; i < 10000; i++ {
		ultimo := lista.VerUltimo()
		if ultimo != cantidad-1 {
			t.Errorf("Se esperaba %d como último elemento, pero se encontró %d", cantidad-1, ultimo)
		}
	}
}
