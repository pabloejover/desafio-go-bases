package tickets

import "fmt"

// Ticket es una estructura que define un ticket de una reserva
type Ticket struct {
	Id      string
	Nombre  string
	Email   string
	Destino string
	Hora    string
	Precio  string
}

// Listado de tickets
type Storage struct {
	Tickets []Ticket
}

func (s *Storage) PrintInfo() {
	fmt.Printf("%v+", s.Tickets)

}

// ejemplo 1
//func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
//func GetMornings(time string) (int, error) {}

// ejemplo 3
//func AverageDestination(destination string, total int) (int, error) {}
