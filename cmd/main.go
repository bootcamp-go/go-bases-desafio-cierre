package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/service"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/storage"
)

func main() {
	// env
	// ...

	// dependencies
	st := storage.NewStorageTicketCSVFile("./tickets.csv")
	sv := service.NewServiceTicketDefault(st)
	
	// app
	// -> tickets to Japan
	total, err := sv.GetTotalTickets("Japan")
	if err != nil {
		panic(err)
	}
	fmt.Println("Total tickets to Japan:", total)

	// -> tickets during the early morning
	total, err = sv.GetCountByPeriod(service.PeriodEarlyMorning)
	if err != nil {
		panic(err)
	}
	fmt.Println("Total tickets during the early morning:", total)

	// -> percentage of tickets to Japan
	percentage, err := sv.PercentageByDestination("Japan")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Percentage of tickets to Japan: %.2f%%\n", percentage)
}