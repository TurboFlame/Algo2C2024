package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"tdas/cola"
	"tdas/pila"
)

const MIN_EXPONENCIAL = 0
const MIN_RAIZ = 0
const MIN_BASE_LOG = 2
const MIN_LOG = 0
const EXCEPCION_DIVISION = 0
const CONDICION_TERNARIO = 0

func main() {
	s := bufio.NewScanner(os.Stdin)
	iterarArchivoAplicarEImprimir(s, calculadoraPolacaInversa)
}

func iterarArchivoAplicarEImprimir(s *bufio.Scanner, funcionAplicada func(cadena string) (int64, error)) {
	for s.Scan() {
		resul, err := funcionAplicada(s.Text())
		if err != nil {
			fmt.Printf("ERROR\n")

		} else {
			fmt.Printf("%d\n", resul)
		}
	}
}

func operarTresArgumentos(miPila pila.Pila[int64], operador func(val1 int64, val2 int64, val3 int64) (int64, error)) error {
	val3, err := intentarDesapilar(miPila)
	val2, err := intentarDesapilar(miPila)
	val1, err := intentarDesapilar(miPila)
	if err != nil {
		return err
	}
	resul, err := operador(val1, val2, val3)
	miPila.Apilar(resul)
	return err
}
func operarDosArgumentos(miPila pila.Pila[int64], operador func(val1 int64, val2 int64) (int64, error)) error {
	val2, err := intentarDesapilar(miPila)
	val1, err := intentarDesapilar(miPila)
	if err != nil {
		return err
	}
	resul, err := operador(val1, val2)
	miPila.Apilar(resul)
	return err
}

func operarUnArgumento(miPila pila.Pila[int64], operador func(val int64) (int64, error)) error {
	val, err := intentarDesapilar(miPila)
	if err != nil {
		return err
	}
	resul, err := operador(val)
	miPila.Apilar(resul)
	return err
}

func convStringAInt64(cadena string) (int64, error) {
	res, err := strconv.Atoi(cadena)

	if err != nil {
		return 0, err
	}
	return int64(res), nil
}

func intentarDesapilar(miPila pila.Pila[int64]) (int64, error) {
	if miPila.EstaVacia() {
		return 0, errors.New("pila vacia")
	}
	return miPila.Desapilar(), nil
}

/*
Esta funcion recibe un string con la operacion en notacion polaca inversa
y devuelve el resultado de la misma.
*/
func calculadoraPolacaInversa(cadena string) (int64, error) {
	var err error = nil
	miPila := pila.CrearPilaDinamica[int64]()
	miCola := separarTokens(cadena)

	for !miCola.EstaVacia() { //Colocar en una funcion
		token := miCola.Desencolar()

		switch token {
		case "+":
			err = operarDosArgumentos(miPila, suma)
		case "-":
			err = operarDosArgumentos(miPila, resta)
		case "*":
			err = operarDosArgumentos(miPila, multiplicacion)
		case "/":
			err = operarDosArgumentos(miPila, division)
		case "?":
			err = operarTresArgumentos(miPila, ternario)
		case "sqrt":
			err = operarUnArgumento(miPila, raiz)
		case "log":
			err = operarDosArgumentos(miPila, logaritmo)
		case "^":
			err = operarDosArgumentos(miPila, exponencial)
		default:
			val, err := convStringAInt64(token)

			if err != nil {
				return 0, err
			}
			miPila.Apilar(val)

		}

		if err != nil {
			return 0, err
		}
		//Luego de operar cada caracter, retorna si existe un error.
	}

	resulFinal, err := intentarDesapilar(miPila)
	if !miPila.EstaVacia() {
		err = errors.New("pila no vacia al final de la operacion")
	}

	return resulFinal, err
}

func separarTokens(cadena string) cola.Cola[string] {
	colaStrings := cola.CrearColaEnlazada[string]()
	cadena += " " //Se agrega un espacio al final de la cadena para asegurarse que no queden elementos residuales en temp.
	temp := ""
	for _, caracter := range cadena {
		if caracter == ' ' && temp != "" {
			colaStrings.Encolar(temp)
			temp = ""
		} else if string(caracter) != " " {
			temp += string(caracter)
		}
	}
	return colaStrings
}

func suma(val1 int64, val2 int64) (int64, error) {
	return val1 + val2, nil
}
func resta(val1 int64, val2 int64) (int64, error) {
	return val1 - val2, nil
}
func division(val1 int64, val2 int64) (int64, error) {
	if val2 == EXCEPCION_DIVISION {
		return 0, errors.New("division por cero")
	}
	return val1 / val2, nil
}
func exponencial(val1 int64, val2 int64) (int64, error) {
	if val2 < MIN_EXPONENCIAL {
		return 0, errors.New("exponente menor a 0")
	}
	return int64(math.Pow(float64(val1), float64(val2))), nil
}
func raiz(val int64) (int64, error) {
	if val < MIN_RAIZ {
		return 0, errors.New("raiz cuadrada por num menor a 0 o igual al indice")
	}
	return int64(math.Sqrt(float64(val))), nil
}
func ternario(val1 int64, val2 int64, val3 int64) (int64, error) {
	if val1 == CONDICION_TERNARIO {
		return val3, nil
	} else {
		return val2, nil
	}
}
func multiplicacion(val1 int64, val2 int64) (int64, error) {
	return val1 * val2, nil
}

func logaritmo(val1 int64, val2 int64) (int64, error) {
	if val2 < MIN_BASE_LOG {
		return 0, errors.New("logaritmo con base menor a 2")
	} else if val1 <= MIN_LOG {
		return 0, errors.New("logaritmo con argumento menor a 0")
	}

	return int64(math.Log(float64(val1)) / math.Log(float64(val2))), nil
}
