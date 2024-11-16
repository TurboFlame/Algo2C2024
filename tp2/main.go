package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	comando_ingresado := os.Args[0]
	if comando_ingresado == "agregar_archivo" {
		ruta_log := os.Args[1]
		agregar_archivo(ruta_log)
	} elif comando_ingresado == "ver_visitantes" {
		desde := os.Args[1]
		hasta := os.Args[2]
		ver_visitantes(desde, hasta)
	} elif comando_ingresado == "mas_visitados" {
		n := os.Args[1]
		mas_visitados(n)
	}
	
}

func agregar_archivo(entrada string) {

}

func ver_visitantes(desde string, hasta string) {

}

func mas_visitados(n int) {

}