package main

/*
ANOTACIONES
Cambiar IpAInt e IntAIp a una funcion de comparacion
Revisitar funcion de comparacion
Cambiar ABB por Hash en funcion DDOS
Implementar counting sort con radix sort para ordenamiento de IPs manteniendo el O(n) o heapsort
Mejorar funcion anonima agregar archivo
Cambiar separar tokens para hacer que la funcion reciba el separador y hacerla universal
*/
import (
	"bufio"
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
	TIEMPO_DOS = 2 * time.Second
	NUMERO_DOS = 5
)

// Creo el struct para manejar mas facilmente las lineas de un archivo log
type lineaLog struct {
	IP    string
	fecha time.Time
	URL   string
}

type paquete struct {
	visitados  diccionario.Diccionario[string, uint]
	visitantes diccionario.DiccionarioOrdenado[string, int]
}
type duo struct {
	URL      string
	cantidad uint
}

func crearPaquete() paquete {
	return paquete{visitantes: diccionario.CrearABB[string, int](compIpMax), visitados: diccionario.CrearHash[string, uint]()}
}

func main() {
	miPaquete := crearPaquete()
	scanner := bufio.NewScanner(os.Stdin)
	entradaValida := true
	for scanner.Scan() && entradaValida {
		entradaValida = procesarEntrada(separarTokens(scanner.Text()), &miPaquete)
	}
}

func procesarEntrada(argumentos cola.Cola[string], miPaquete *paquete) bool {

	switch argumentos.Desencolar() {
	case "agregar_archivo":
		ruta_log := argumentos.Desencolar()
		agregar_archivo(ruta_log, miPaquete)
	case "ver_visitantes":
		desde := argumentos.Desencolar()
		hasta := argumentos.Desencolar()
		ver_visitantes(desde, hasta, miPaquete)
	case "ver_mas_visitados":
		n, err := strconv.Atoi(argumentos.Desencolar())
		if err != nil {
			panic("Error de conversion")
		}
		ver_mas_visitados(n, miPaquete)
	default:
		fmt.Fprintf(os.Stderr, "Error en comando: %s\n", argumentos)
		fmt.Println("Error ejecutado")
		return false
	}
	return true
}

func agregar_archivo(ruta_archivo string, miPaquete *paquete) {
	dictDDOS := diccionario.CrearHash[string, []time.Time]()

	iterarArchivoYAplicar(ruta_archivo, func(lineaTexto string) error {

		var lista []time.Time
		linea := procesar_linea(lineaTexto)

		if dictDDOS.Pertenece(linea.IP) {
			lista = append(dictDDOS.Obtener(linea.IP), linea.fecha) //Posiblemente cambiar a utilizar punteros para no inicializar un nuevo array o buscar forma de simplificar codigo
		} else {
			lista = []time.Time{linea.fecha}
		}
		dictDDOS.Guardar(linea.IP, lista)

		if miPaquete.visitados.Pertenece(linea.URL) {
			miPaquete.visitados.Guardar(linea.URL, miPaquete.visitados.Obtener(linea.URL)+1)
		} else {
			miPaquete.visitados.Guardar(linea.URL, 1)
		}

		miPaquete.visitantes.Guardar(linea.IP, 0)

		return nil
	})

	busquedaDOS(dictDDOS)
	fmt.Println("OK")
}

func ver_visitantes(desde string, hasta string, miPaquete *paquete) {
	fmt.Println("Visitantes:")
	miPaquete.visitantes.IterarRango(&desde, &hasta, func(ip string, dato int) bool {
		fmt.Print("\t", ip, "\n")
		return true
	})
	fmt.Println("OK")

}

func ver_mas_visitados(cuantos int, miPaquete *paquete) {
	visitados := make([]duo, miPaquete.visitados.Cantidad())
	contador := 0
	miPaquete.visitados.Iterar(func(clave string, dato uint) bool {
		visitados[contador] = duo{URL: clave, cantidad: dato}
		contador++
		return true
	})
	heapVisitados := cola_prioridad.CrearHeapArr(visitados, compDuo)

	contador = 0
	fmt.Println("Sitios más visitados:")
	for contador < cuantos && !heapVisitados.EstaVacia() {
		duo := heapVisitados.Desencolar()
		fmt.Print("\t", duo.URL, " - ", duo.cantidad, "\n")
		contador++
	}
	fmt.Println("OK")
}

func iterarArchivoYAplicar(ruta_archivo string, funcionAplicada func(cadena string) error) {
	archivo, err := os.Open(ruta_archivo)
	if err != nil {
		panic("Error al abrir el archivo")
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		err := funcionAplicada(s.Text())
		if err != nil {
			panic("Error en interpretacion de una linea")
		}
	}
}

func procesar_linea(linea string) lineaLog {
	palabras := strings.Fields(linea)
	ip := palabras[0]
	fecha := palabras[1]
	URL := palabras[3]
	fecha_parseada, _ := time.Parse("2006-01-02T15:04:05", fecha)
	linea_log := lineaLog{ip, fecha_parseada, URL}
	return linea_log
}

func busquedaDOS(dict diccionario.Diccionario[string, []time.Time]) {
	slice := make([]string, 0)
	dict.Iterar(func(ip string, tiempos []time.Time) bool {
		for i := NUMERO_DOS - 1; i < len(tiempos); i++ {
			if tiempos[i].Sub(tiempos[i-(NUMERO_DOS-1)]) < TIEMPO_DOS { // Chequeo tiempos en grupos de 5
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

func compIp(a, b string, visitar func(a, b int) int) int {
	splitA := separarTokensIp(a)
	splitB := separarTokensIp(b)
	for i := 0; i < len(splitA); i++ {
		if !(splitA[i] == splitB[i]) {
			return visitar(splitA[i], splitB[i])
		}
	}
	return 0
}
func compIpMin(a, b string) int {
	return compIp(a, b, func(a, b int) int { return b - a })
}
func compIpMax(a, b string) int {
	return compIp(a, b, func(a, b int) int { return a - b })
}
func compDuo(a, b duo) int { return int(a.cantidad) - int(b.cantidad) }

func separarTokensIp(ip string) []int {
	slice := make([]int, 4)
	for i, caracter := range strings.Split(ip, ".") {
		num, err := strconv.Atoi(caracter)
		if err != nil {
			panic("Error de conversion") //Posiblemente tenga que manejar el error mas adelante.
		}
		slice[i] = num
	}
	return slice
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

// Aux
/*func procesar_archivo(ruta_archivo string) []lineaLog {
	archivo, err := os.Open(ruta_archivo)
	if err != nil {
		panic("Error al abrir el archivo")
	}
	defer archivo.Close()

	var lineas []lineaLog
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		nueva_entrada := procesar_linea(linea) // Convierte a la linea en el struct
		lineas = append(lineas, nueva_entrada)
	}
	return lineas
}
// Para poder comparar numericamente las IPs
func ipAInt(ip string) uint32 {
	var resultado uint32
	partes := strings.Split(ip, ".")
	for i := 0; i < 4; i++ {
		parte, _ := strconv.Atoi(partes[i])
		resultado = resultado<<8 + uint32(parte)
	}
	return resultado
}

// Para imprimir las IPs en su formato original
func intAIP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

*/
