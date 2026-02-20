package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Client struct {
	ID          string
	CreditScore int
	Income      float64
	Debt        float64
	Overdue     int
}

func generateCSV(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"client_id", "credit_score", "income", "debt", "overdue"})

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		record := []string{
			fmt.Sprintf("C%03d", i+1),
			strconv.Itoa(rand.Intn(550) + 300),              // 300–850
			fmt.Sprintf("%.2f", rand.Float64()*100000+20000), // 20k–120k
			fmt.Sprintf("%.2f", rand.Float64()*50000),        // 0–50k
			strconv.Itoa(rand.Intn(2)),                       // 0 or 1
		}
		writer.Write(record)
	}
}

func main() {
	filename := "clients.csv"
	generateCSV(filename, 100)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var totalScore, totalDebt float64
	var overdueScore float64
	var count, overdueCount float64

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		score, _ := strconv.ParseFloat(row[1], 64)
		debt, _ := strconv.ParseFloat(row[3], 64)
		overdue, _ := strconv.Atoi(row[4])

		totalScore += score
		totalDebt += debt
		count++

		if overdue == 1 {
			overdueScore += score
			overdueCount++
		}
	}

	fmt.Println("=== Credit Risk Analytics Report ===")
	fmt.Printf("Total clients: %.0f\n", count)
	fmt.Printf("Average credit score: %.2f\n", totalScore/count)
	fmt.Printf("Average debt: %.2f\n", totalDebt/count)

	if overdueCount > 0 {
		fmt.Printf("Average credit score (overdue clients): %.2f\n", overdueScore/overdueCount)
	}
}