package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/heap"
	"tdas/lista"
	"time"
)

type lineaLog struct {
	ip    string
	fecha time.Time
	URL   string
}

func main() {
	comando_ingresad := os.Args[0]
	if comando_ingresado == "agregar_archivo" {
		ruta_log := os.Args[1]
		agregar_archivo(ruta_log)
	} else if comando_ingresado == "ver_visitantes" {
		desde := os.Args[1]
		hasta := os.Args[2]
		ver_visitantes(desde, hasta)
	} else if comando_ingresado == "mas_visitados" {
		n := os.Args[1]
		mas_visitados(n)
	}
}

func agregar_archivo(ruta_archivo string) {
	archivo, err := os.Open(ruta_archivo)
	defer archivo.Close()
	var lineas []lineaLog
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		nueva_entrada := procesar_linea(linea) // Convierte a la linea en el struct
		lineas = append(lineas, nueva_entrada)
	}
	busquedaDOS(lineas)
}

func ver_visitantes(desde string, hasta string) {

}

func mas_visitados(n int) {

}

// Aux

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

	for _, linea := range lineas {

	}
}
