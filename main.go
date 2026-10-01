package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

type InvestType struct {
	InvestBalance decimal.Decimal
	LowRisk decimal.Decimal
	MediumRisk decimal.Decimal
	HighRisk decimal.Decimal
}

type UserBankData struct {
	Username string
	Cash decimal.Decimal
	Balance decimal.Decimal
	InvestFund InvestType
}

var userDatas = []UserBankData{
		{
			Username: "Alice",
			Cash: decimal.NewFromFloat(1000),
			Balance: decimal.NewFromFloat(0),
			InvestFund: InvestType{
				InvestBalance: decimal.NewFromFloat(0),
				LowRisk: decimal.NewFromFloat(0),
				MediumRisk: decimal.NewFromFloat(0),
				HighRisk: decimal.NewFromFloat(0),
			},
		},
		{
			Username: "Bob",
			Cash: decimal.NewFromInt(1000),
			Balance: decimal.NewFromFloat(0),
			InvestFund: InvestType{
				InvestBalance: decimal.NewFromFloat(0),
				LowRisk: decimal.NewFromFloat(0),
				MediumRisk: decimal.NewFromFloat(0),
				HighRisk: decimal.NewFromFloat(0),
			},
		},		
		{
			Username: "Charlie",
			Cash: decimal.NewFromInt(1000),
			Balance: decimal.NewFromFloat(0),
			InvestFund: InvestType{
				InvestBalance: decimal.NewFromFloat(0),
				LowRisk: decimal.NewFromFloat(0),
				MediumRisk: decimal.NewFromFloat(0),
				HighRisk: decimal.NewFromFloat(0),
			},
		},
				{
			Username: "Diana",
			Cash: decimal.NewFromInt(1000),
			Balance: decimal.NewFromFloat(0),
			InvestFund: InvestType{
				InvestBalance: decimal.NewFromFloat(0),
				LowRisk: decimal.NewFromFloat(0),
				MediumRisk: decimal.NewFromFloat(0),
				HighRisk: decimal.NewFromFloat(0),
			},
		},
	}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	username := chooseUsername(scanner)
	inOperation(scanner, username)
}

func inOperation(scanner *bufio.Scanner, username string) {
	userData := UserBankData{}
	for _, data := range userDatas {
		if data.Username == username {
			userData = data
		}
	}
	for {
		showBankMenu()
		operation := scanText(scanner)
		switch operation {
		case "1":
			// Show balance
			showBalance(&userData)
		case "2":
			// Deposit money
			depositMoney(&userData, scanner)
		case "3":
			// Withdraw money
			withdrawMoney(&userData, scanner)
		case "4":
			// Send money to a person
			sendMoneytoPpl(&userData, scanner)
		case "5":
			// Invest
		case "6":
			// Transfer between accounts
		case "7":
			// Withdraw all investment
		case "8":
			// Logout
		case "9":
			// Exit
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid operation!. Choose again")
		}
	}
}

func investMoney(userData *UserBankData, scanner *bufio.Scanner) {
	
}

func sendMoneytoPpl(userData *UserBankData, scanner *bufio.Scanner) {
	fmt.Print("Enter the username you want to send money to: ")
	found := false
	OUTERLOOP:
	for{
		user := scanText(scanner) 
		if user == userData.Username {
			fmt.Println("Invalid username! Please try again!")
			continue
		}
		for _, u :=range userDatas {
			if u.Username == user {
				found = true
				fmt.Print("Enter the amount you want to send to " + user + " :")
				amount, err := decimal.NewFromString(scanText(scanner))
				if err!=nil {
					fmt.Println("Invalid withdraw!")
				} else if amount.GreaterThan(userData.Balance) {
					fmt.Println("Insufficient fund! Try again")
				} else {
					userData.Balance = userData.Balance.Sub(amount)
					u.Balance = u.Balance.Add(amount)
					break OUTERLOOP
				}
			} 
		}
		if !found {
			fmt.Println("Invalid username! Please try again!")
		}
	}
}

func withdrawMoney(userData *UserBankData, scanner *bufio.Scanner) {
	for {
		fmt.Print("Enter the amount you want to withdraw: ")
		withdraw, err := decimal.NewFromString(scanText(scanner))
		if err != nil {
			fmt.Println("Invalid withdraw!")
		} else if withdraw.GreaterThan(userData.Balance) {
			fmt.Println("Insufficient fund! Try again")
		} else {
			userData.Balance = userData.Balance.Sub(withdraw)
			userData.Cash = withdraw
			break
		}
	}
}

func depositMoney(userData *UserBankData, scanner *bufio.Scanner) {
	for {
		fmt.Print("Enter the amount you want to deposit: ")
		deposit, err := decimal.NewFromString(scanText(scanner))
		if err != nil {
			fmt.Println("Invalid deposit")
		} else if deposit.GreaterThan(userData.Cash) {
			fmt.Println("Invalid cash fund!")
		} else {
			userData.Balance = deposit
			userData.Cash = userData.Cash.Sub(deposit)
			break
		}
	}
}

func showBalance(userData *UserBankData) {
	interest := userData.Balance.Mul(decimal.NewFromFloat(0.01))
	userData.Balance = userData.Balance.Add(interest)
	fmt.Println("Your balance is: $", userData.Balance)
}

func showBankMenu() {
	fmt.Println(
`
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
		os.Exit(1)
	}
	return input
}

func chooseUsername(scanner *bufio.Scanner) string {
	fmt.Print("Enter your username: ")
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
	fmt.Println("Welcome", username)
	return username
}