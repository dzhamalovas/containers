package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func RunAnalytics(dbFile string) error {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	var avgScore float64
	var avgIncome float64
	var avgDebt float64
	var delinquencyRate float64
	var highRiskCount int
	var totalClients int

	// Средние значения
	err = db.QueryRow("SELECT AVG(credit_score) FROM clients").Scan(&avgScore)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT AVG(income) FROM clients").Scan(&avgIncome)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT AVG(debt) FROM clients").Scan(&avgDebt)
	if err != nil {
		return err
	}

	// Доля просрочек
	err = db.QueryRow("SELECT AVG(delinquent) FROM clients").Scan(&delinquencyRate)
	if err != nil {
		return err
	}

	// Количество клиентов
	err = db.QueryRow("SELECT COUNT(*) FROM clients").Scan(&totalClients)
	if err != nil {
		return err
	}

	// Высокий риск (рейтинг < 500 или debt/income > 0.6)
	queryHighRisk := `
		SELECT COUNT(*) FROM clients
		WHERE credit_score < 500
		OR (debt / income) > 0.6
	`
	err = db.QueryRow(queryHighRisk).Scan(&highRiskCount)
	if err != nil {
		return err
	}

	fmt.Println("========== АНАЛИТИКА CREDIT RISK ==========")
	fmt.Printf("Всего клиентов: %d\n", totalClients)
	fmt.Printf("Средний кредитный рейтинг: %.2f\n", avgScore)
	fmt.Printf("Средний доход: %.2f\n", avgIncome)
	fmt.Printf("Средний долг: %.2f\n", avgDebt)
	fmt.Printf("Доля просрочек: %.2f\n", delinquencyRate)
	fmt.Printf("Клиенты высокого риска: %d\n", highRiskCount)

	fmt.Println("============================================")

	log.Println("Аналитика успешно рассчитана")

	return nil
}