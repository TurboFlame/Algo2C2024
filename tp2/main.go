package main

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
	SEPARADOR_ESPACIO                 = ' '
	SEPARADOR_IP                      = '.'
	CANT_CAMPOS_LINEA                 = 4
)

// Creo el struct para manejar mas facilmente las lineas de un archivo log
type lineaLog struct {
	IP    uint32
	fecha time.Time
	URL   string
}

type paquete struct {
	visitados  diccionario.Diccionario[string, uint]
	visitantes diccionario.DiccionarioOrdenado[uint32, uint]
}
type sitioYVisitas struct {
	URL      string
	cantidad uint
}

func crearPaquete() paquete {
	return paquete{visitantes: diccionario.CrearABB[uint32, uint](compIpMax), visitados: diccionario.CrearHash[string, uint]()}
}

func main() {
	miPaquete := crearPaquete()
	scanner := bufio.NewScanner(os.Stdin)
	entradaValida := true
	for scanner.Scan() && entradaValida {
		argumentos, cantArgumentos := separarTokens(scanner.Text(), SEPARADOR_ESPACIO)
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
		visitantes := verVisitantes(desde, hasta, miPaquete)
		imprimirVisitantes(visitantes)
	case entrada == "ver_mas_visitados" && cantArgumentos == CANT_ARGUMENTOS_VER_MAS_VISITADOS:
		n, err := strconv.Atoi(argumentos.Desencolar())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", entrada)
			return false
		}
		visitados := verMasVisitados(n, miPaquete)
		imprimirMasVisitados(visitados)
	default:
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", entrada)
		return false
	}
	return true
}

func procesarLinea(linea string) (lineaLog, error) {
	palabras := strings.Fields(linea)
	if len(palabras) != CANT_CAMPOS_LINEA {
		return lineaLog{}, errors.New("Error, cantidad incorrecta de parametros")
	}
	ip := palabras[0]
	ipNum, err := ipAUint32(ip)
	if err != nil {
		return lineaLog{}, errors.New("Error al procesar IP")
	}

	fecha := palabras[1]
	URL := palabras[3]
	fechaParseada, err := time.Parse("2006-01-02T15:04:05-07:00", fecha)
	lineaLog := lineaLog{ipNum, fechaParseada, URL}
	return lineaLog, err
}

func agregarArchivo(rutaArchivo string, miPaquete *paquete) error {
	dictDOS := diccionario.CrearHash[uint32, []time.Time]()
	err := iterarArchivoYAplicar(rutaArchivo, func(lineaTexto string) error {
		linea, errLinea := procesarLinea(lineaTexto)
		if errLinea != nil {
			return errLinea
		}
		//Si la clave pertenece, carga el arreglo con linea.fecha como unico elemento. Sino, utiliza un append y agrega linea.fecha al final de la lista.
		set[uint32, []time.Time](dictDOS, linea.IP, []time.Time{linea.fecha}, func(lista []time.Time) []time.Time { return append(lista, linea.fecha) })
		set[string, uint](miPaquete.visitados, linea.URL, 1, func(num uint) uint { return num + 1 })
		miPaquete.visitantes.Guardar(linea.IP, 0)
		return nil
	})

	if err != nil {
		return err
	}
	sospechosos := busquedaDOS(dictDOS)
	imprimirDOS(sospechosos)
	return nil
}

func set[K comparable, V any](dict diccionario.Diccionario[K, V], clave K, valorDefault V, visita func(V) V) {
	if dict.Pertenece(clave) {
		dict.Guardar(clave, visita(dict.Obtener(clave)))
	} else {
		dict.Guardar(clave, valorDefault)
	}
}

func verVisitantes(desde string, hasta string, miPaquete *paquete) []uint32 {
	desdeNum, _ := ipAUint32(desde)
	hastaNum, _ := ipAUint32(hasta)
	visitantes := make([]uint32, 0)
	miPaquete.visitantes.IterarRango(&desdeNum, &hastaNum, func(ip uint32, dato uint) bool {
		visitantes = append(visitantes, ip)
		return true
	})
	return visitantes
}

func imprimirVisitantes(visitantes []uint32) {
	fmt.Println("Visitantes:")
	for _, ip := range visitantes {
		fmt.Printf("\t%s\n", uint32AIP(ip))
	}
	fmt.Println("OK")
}

func verMasVisitados(cuantos int, miPaquete *paquete) []sitioYVisitas {
	visitados := make([]sitioYVisitas, miPaquete.visitados.Cantidad())
	contador := 0
	miPaquete.visitados.Iterar(func(clave string, dato uint) bool {
		visitados[contador] = sitioYVisitas{URL: clave, cantidad: dato}
		contador++
		return true
	})
	heapVisitados := cola_prioridad.CrearHeapArr(visitados, compURL) // Ordeno con heapify para mantener la complejidad O(s)

	resultado := make([]sitioYVisitas, 0, cuantos)
	contador = 0
	for contador < cuantos && !heapVisitados.EstaVacia() { // O(k(log(s)). K veces se utiliza desencolar en un heap de s elementos.
		duo := heapVisitados.Desencolar()
		resultado = append(resultado, duo)
		contador++
	}
	return resultado
}

func imprimirMasVisitados(sitios []sitioYVisitas) {
	fmt.Println("Sitios más visitados:")
	for _, sitio := range sitios {
		fmt.Printf("\t%s - %d\n", sitio.URL, sitio.cantidad)
	}
	fmt.Println("OK")
}

func iterarArchivoYAplicar(rutaArchivo string, funcionAplicada func(cadena string) error) error {
	archivo, err := os.Open(rutaArchivo)
	defer archivo.Close()
	if err != nil {
		return errors.New("Error al abrir archivo " + rutaArchivo)
	}
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		err := funcionAplicada(s.Text())
		if err != nil {
			return err
		}
	}
	return nil
}

func busquedaDOS(dict diccionario.Diccionario[uint32, []time.Time]) []uint32 {
	detecciones := make([]uint32, 0)

	// Recibe un diccionario con todas las IPs y una lista de la hora de cada una de sus requests.
	// Cuando encuentra cinco requests hechas en menos de dos segundos, agrega
	dict.Iterar(func(ip uint32, tiempos []time.Time) bool {
		for i := CANT_MAX_REQUESTS - 1; i < len(tiempos); i++ {
			if tiempos[i].Sub(tiempos[i-(CANT_MAX_REQUESTS-1)]) < TIEMPO_MAXIMO_REQUESTS { // Chequeo tiempos en grupos de 5
				detecciones = append(detecciones, ip)
				break
			}
		}
		return true
	})

	// Verificar si detecciones no está vacío antes de ordenar
	if len(detecciones) > 0 {
		radixSortIPs(detecciones)
	}
	return detecciones
}

func radixSortIPs(ips []uint32) {
	var bits = 8
	var base uint32 = 1 << bits
	var filtro uint32 = base - 1

	n := len(ips)
	if n == 0 {
		return
	}

	aux := make([]uint32, n)

	for desplazamiento := 0; desplazamiento < 32; desplazamiento += bits {
		conteo := make([]int, base)
		for _, ip := range ips {
			digito := (ip >> desplazamiento) & filtro
			conteo[digito]++
		}
		if len(conteo) == 0 {
			return
		}
		for i := 1; uint32(i) < base; i++ {
			conteo[i] += conteo[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digito := (ips[i] >> desplazamiento) & filtro
			conteo[digito]--
			aux[conteo[digito]] = ips[i]
		}
		copy(ips, aux)
	}
}

func imprimirDOS(detecciones []uint32) {
	for _, ip := range detecciones {
		fmt.Printf("DoS: %s\n", uint32AIP(ip))
	}
	fmt.Println("OK")
}

func compURL(a, b sitioYVisitas) int { return int(a.cantidad) - int(b.cantidad) }

func compIpMin(a, b uint32) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}

func compIpMax(a, b uint32) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func ipAUint32(ip string) (uint32, error) {
	partes := strings.Split(ip, ".")
	if len(partes) != 4 {
		return 0, errors.New("IP no valida")
	}
	var ipNum uint32
	for i := 0; i < 4; i++ {
		parte, err := strconv.Atoi(partes[i])
		if err != nil || parte < 0 || parte > 255 {
			return 0, errors.New("IP no valida")
		}
		ipNum = ipNum<<8 + uint32(parte)
	}
	return ipNum, nil
}

func uint32AIP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
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

func ipEsValida(ip string) bool {
	_, err := ipAUint32(ip)
	return err == nil
}
