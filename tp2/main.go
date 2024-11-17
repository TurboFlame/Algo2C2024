package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/lista"
	"time"
)

const (
	TIEMPO_DOS = 2
	NUMERO_DOS = 5
)

// Creo el struct para manejar mas facilmente las lineas de un archivo log
type lineaLog struct {
	IP    string
	fecha time.Time
	URL   string
}

func main() {
	comando_ingresado := os.Args[0]
	if comando_ingresado == "agregar_archivo" {
		ruta_log := os.Args[1]
		agregar_archivo(ruta_log)
	} else if comando_ingresado == "ver_visitantes" {
		desde := os.Args[1]
		hasta := os.Args[2]
		ver_visitantes(desde, hasta)
	} else if comando_ingresado == "ver_mas_visitados" {
		n := os.Args[1]
		ver_mas_visitados(n)
	} else {
		os.FPrintf(os.Stderr, "Error en comando", comando_ingresado)
	}
}

func agregar_archivo(ruta_archivo string) {
	lineas := procesar_archivo(ruta_archivo)
	busquedaDOS(lineas)
	fmt.Println("OK")
}

func ver_visitantes(desde string, hasta string) {
	
}

func ver_mas_visitados(n int) {
	//Implementar despues
}

// Aux
func procesar_archivo(ruta_archivo string) []lineaLog {
	archivo, err := os.Open(ruta_archivo)
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

func procesar_linea(linea string) lineaLog {
	palabras := strings.Fields(linea)
	ip := palabras[0]
	fecha := palabras[1]
	URL := palabras[3]
	fecha_parseada, _ := time.Parse("2006-01-02T15:04:05", fecha)
	linea_log := lineaLog{ip, fecha_parseada, URL}
	return linea_log
}

func busquedaDOS(lineas []lineaLog) {
	abb := diccionario.CrearABB[uint32, []time.Time](func(a, b uint32) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for _, linea := range lineas {
		ip := ipAInt(linea.IP)
		if abb.Pertenece(ip) {
			lista := abb.Obtener(ip)
			lista = append(lista, linea.fecha)
			abb.Guardar(ip, lista)
		} else {
			lista := []time.Time{linea.fecha}
			abb.Guardar(ip, lista)
		}
	}

	abb.Iterar(func(ip uint32, tiempos []time.Time) bool {
		for i := 0; i <= len(tiempos)-NUMERO_DOS; i++ {
			if tiempos[i+(NUMERO_DOS-1)].Sub(tiempos[i]) < TIEMPO_DOS { // Chequeo tiempos en grupos de 5
				fmt.Printf("DoS: %s\n", ip)
				break
			}
		}
		return true
	})
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
