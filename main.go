package main

import (
	"log"
	"os"
	"strings"

	"github.com/pabloejover/desafio-go-bases.git/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	storage := tickets.Storage{
		Tickets: readFile(filename),
	}

	storage.PrintInfo()

}

// readFile es una funcion creada para leer un archivo
func readFile(filename string) []tickets.Ticket {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	data := strings.Split(string(file), "\n")

	var resultado []tickets.Ticket

	for i := 0; i < len(data); i++ {
		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			ticket := tickets.Ticket{
				Id:      file[0],
				Nombre:  file[1],
				Email:   file[2],
				Destino: file[3],
				Hora:    file[4],
				Precio:  file[5],
			}
			resultado = append(resultado, ticket)
		}

	}
	return resultado

}
