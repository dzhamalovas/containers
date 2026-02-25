package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	numClients := 1000

	fmt.Println("Генерация данных...")
	err := GenerateData("data/clients.csv", numClients)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Загрузка в SQLite...")
	err = LoadToSQLite("data/clients.csv", "data/clients.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Расчёт аналитики...")
	err = RunAnalytics("data/clients.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Готово.")
	os.Exit(0)
}