package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

type InvestType struct {
	LowRisk decimal.Decimal
	MediumRisk decimal.Decimal
	HighRisk decimal.Decimal
}

type UserBankData struct {
	Username string
	Cash decimal.Decimal
	Deposit decimal.Decimal
	InvestFund InvestType
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	chooseUsername(scanner)
}

func showBankMenu() {
	fmt.Println(`
		--- Banking App Menu ---
		1. Show balance
		2. Deposit money
		3. Withdraw money
		4. Send money to a person
		5. Invest in funds
		6. Transfer between accounts
		7. Withdraw all investments
		8. Logout
		9. Exit
	`)
}

func scanText(scanner *bufio.Scanner) string {
	input:=""
	if scanner.Scan(){
		input= strings.TrimSpace(scanner.Text())
	} else{
		fmt.Println("EOF: Exiting")
		os.Exit(0)
	}
	return input
}

func chooseUsername(scanner *bufio.Scanner) {
	fmt.Println("Enter your username:")
	username:=scanText(scanner)
	switch username {
	case "Alice":
	case "Bob":
	case "Charlie":
	case "Diana":
	default:
		fmt.Println("Invalid username. Please try again!")
		chooseUsername(scanner)
	}
}