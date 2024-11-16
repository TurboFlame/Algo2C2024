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
	comando_ingresado := os.Args[0]
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

