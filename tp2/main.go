package main

/*
ANOTACIONES
Mejorar funcion anonima agregar archivo
Cambiar separar tokens para hacer que la funcion reciba el separador y hacerla universal unificando separartok y separartokip
Manejar error ip invalida
usar separarTokens en compIP
*/
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"time"
)

const (
	TIEMPO_MAXIMO_REQUESTS            = 2 * time.Second
	CANT_MAX_REQUESTS                 = 5
	CANT_ARGUMENTOS_AGREGAR_ARCHIVO   = 2
	CANT_ARGUMENTOS_VER_VISITANTES    = 3
	CANT_ARGUMENTOS_VER_MAS_VISITADOS = 2
	CAMPOS_IP                         = 4
)

// Creo el struct para manejar mas facilmente las lineas de un archivo log
type lineaLog struct {
	IP    string
	fecha time.Time
	URL   string
}

type paquete struct {
	visitados  diccionario.Diccionario[string, uint]
	visitantes diccionario.DiccionarioOrdenado[string, uint]
}
type sitioYVisitas struct {
	URL      string
	cantidad uint
}

func crearPaquete() paquete {
	return paquete{visitantes: diccionario.CrearABB[string, uint](compIpMax), visitados: diccionario.CrearHash[string, uint]()}
}

func main() {
	miPaquete := crearPaquete()
	scanner := bufio.NewScanner(os.Stdin)
	entradaValida := true
	for scanner.Scan() && entradaValida {
		argumentos, cantArgumentos := separarTokens(scanner.Text(), ' ')
		entradaValida = procesarEntrada(argumentos, cantArgumentos, &miPaquete)
	}
}

func procesarEntrada(argumentos cola.Cola[string], cantArgumentos int, miPaquete *paquete) bool {

	entrada := argumentos.Desencolar()
	switch {
	case entrada == "agregar_archivo" && cantArgumentos == CANT_ARGUMENTOS_AGREGAR_ARCHIVO:
		rutaLog := argumentos.Desencolar()
		err := agregarArchivo(rutaLog, miPaquete)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", entrada)
			return false
		}
	case entrada == "ver_visitantes" && cantArgumentos == CANT_ARGUMENTOS_VER_VISITANTES:
		desde := argumentos.Desencolar()
		hasta := argumentos.Desencolar()
		verVisitantes(desde, hasta, miPaquete)
	case entrada == "ver_mas_visitados" && cantArgumentos == CANT_ARGUMENTOS_VER_MAS_VISITADOS:
		n, err := strconv.Atoi(argumentos.Desencolar())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", entrada)
			return false
		}
		verMasVisitados(n, miPaquete)
	default:
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", entrada)
		return false
	}
	return true
}
func agregarArchivo(rutaArchivo string, miPaquete *paquete) error {
	dictDOS := diccionario.CrearHash[string, []time.Time]()
	err := iterarArchivoYAplicar(rutaArchivo, func(lineaTexto string) error {

		linea := procesarLinea(lineaTexto)
		//Si la clave pertenece, carga el arreglo con linea.fecha como unico elemento. Sino, utiliza un append y agrega linea.fecha al final de la lista.
		set[string, []time.Time](dictDOS, linea.IP, []time.Time{linea.fecha}, func(lista []time.Time) []time.Time { return append(lista, linea.fecha) })
		set[string, uint](miPaquete.visitados, linea.URL, 1, func(num uint) uint { return num + 1 })
		miPaquete.visitantes.Guardar(linea.IP, 0)
		return nil
	})

	if err != nil {
		return err
	}
	busquedaDOS(dictDOS)
	fmt.Println("OK")
	return nil
}
func set[K comparable, V any](dict diccionario.Diccionario[K, V], clave K, valorDefault V, visita func(V) V) {
	if dict.Pertenece(clave) {
		dict.Guardar(clave, visita(dict.Obtener(clave)))
	} else {
		dict.Guardar(clave, valorDefault)
	}
}

func verVisitantes(desde string, hasta string, miPaquete *paquete) {
	fmt.Println("Visitantes:")
	miPaquete.visitantes.IterarRango(&desde, &hasta, func(ip string, dato uint) bool {
		fmt.Print("\t", ip, "\n")
		return true
	})
	fmt.Println("OK")
}

func verMasVisitados(cuantos int, miPaquete *paquete) {
	visitados := make([]sitioYVisitas, miPaquete.visitados.Cantidad())
	contador := 0
	miPaquete.visitados.Iterar(func(clave string, dato uint) bool {
		visitados[contador] = sitioYVisitas{URL: clave, cantidad: dato}
		contador++
		return true
	})
	heapVisitados := cola_prioridad.CrearHeapArr(visitados, compURL) //Ordeno con heapify para mantener la complejidad O(s)

	contador = 0
	fmt.Println("Sitios más visitados:")
	for contador < cuantos && !heapVisitados.EstaVacia() { //O(k(log(s)). K veces se utiliza desencolar en un heap de s elementos.
		duo := heapVisitados.Desencolar()
		fmt.Print("\t", duo.URL, " - ", duo.cantidad, "\n")
		contador++
	}
	fmt.Println("OK")
}

func iterarArchivoYAplicar(rutaArchivo string, funcionAplicada func(cadena string) error) error {
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return errors.New("Error al abrir archivo " + rutaArchivo)
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		err := funcionAplicada(s.Text())
		if err != nil {
			return errors.New("Error al interpretar archivo " + rutaArchivo)
		}
	}
	return nil
}

func procesarLinea(linea string) lineaLog {
	palabras := strings.Fields(linea)
	ip := palabras[0]
	fecha := palabras[1]
	URL := palabras[3]
	fechaParseada, _ := time.Parse("2006-01-02T15:04:05-07:00", fecha)
	lineaLog := lineaLog{ip, fechaParseada, URL}
	return lineaLog
}

func busquedaDOS(dict diccionario.Diccionario[string, []time.Time]) {
	slice := make([]string, 0)

	dict.Iterar(func(ip string, tiempos []time.Time) bool {
		for i := CANT_MAX_REQUESTS - 1; i < len(tiempos); i++ {
			if tiempos[i].Sub(tiempos[i-(CANT_MAX_REQUESTS-1)]) < TIEMPO_MAXIMO_REQUESTS { // Chequeo tiempos en grupos de 5
				slice = append(slice, ip)
				break
			}
		}
		return true
	})
	heap := cola_prioridad.CrearHeapArr(slice, compIpMin)
	for !heap.EstaVacia() {
		fmt.Printf("DoS: %s\n", heap.Desencolar())
	}
}

func compURL(a, b sitioYVisitas) int { return int(a.cantidad) - int(b.cantidad) }

func compIpMin(a, b string) int {
	return compIp(a, b, func(a, b int) int { return b - a })
}
func compIpMax(a, b string) int {
	return compIp(a, b, func(a, b int) int { return a - b })
}
func separarTokens(cadena string, separador rune) (cola.Cola[string], int) {
	colaStrings := cola.CrearColaEnlazada[string]()
	cadena += string(separador) //Se agrega un espacio al final de la cadena para asegurarse que no queden elementos residuales en temp.
	temp := ""
	contador := 0
	for _, caracter := range []rune(cadena) {
		if caracter == separador && temp != "" {
			colaStrings.Encolar(temp)
			temp = ""
			contador++
		} else if caracter != separador {
			temp += string(caracter)
		}
	}
	return colaStrings, contador
}

func compIp(a, b string, visitar func(a, b int) int) int {
	colaA, tamA := separarTokens(a, '.')
	colaB, tamB := separarTokens(b, '.')
	if tamA != CAMPOS_IP || tamB != CAMPOS_IP {
		panic("IP no valida (cantidad incorrecta de campos)")
	}
	for i := 0; i < CAMPOS_IP; i++ {
		numA, err1 := strconv.Atoi(colaA.Desencolar())
		numB, err2 := strconv.Atoi(colaB.Desencolar())
		if err1 != nil || err2 != nil {
			panic("IP no valida (no numerica)")
		}
		if numA != numB {
			return visitar(numA, numB)
		}
	}
	return 0
}
