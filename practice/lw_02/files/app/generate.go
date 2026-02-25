package main

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func clamp(val, min, max float64) float64 {
	return math.Max(min, math.Min(max, val))
}

func GenerateData(filename string, n int) error {

	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"id", "age", "income", "debt", "credit_score", "delinquent"})

	for i := 1; i <= n; i++ {

		age := rand.Intn(42) + 18

		income := clamp(rand.NormFloat64()*10000+40000, 20000, 100000)
		debt := clamp(rand.NormFloat64()*7000+15000, 0, income*0.9)

		creditScore := clamp(rand.NormFloat64()*80+650, 300, 850)

		delinquent := 0
		if debt/income > 0.6 || creditScore < 500 {
			if rand.Float64() < 0.6 {
				delinquent = 1
			}
		} else if rand.Float64() < 0.1 {
			delinquent = 1
		}

		writer.Write([]string{
			strconv.Itoa(i),
			strconv.Itoa(age),
			strconv.FormatFloat(income, 'f', 2, 64),
			strconv.FormatFloat(debt, 'f', 2, 64),
			strconv.FormatFloat(creditScore, 'f', 0, 64),
			strconv.Itoa(delinquent),
		})
	}

	return nil
}