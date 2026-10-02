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
	for {
		username := chooseUsername(scanner)
		inOperation(scanner, &username)
	}
	
}

func inOperation(scanner *bufio.Scanner, username *string) {
	userData := UserBankData{}
	for _, data := range userDatas {
		if data.Username == *username {
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
			investMoney(&userData, scanner)
		case "6":
			// Transfer between accounts
			transferAccount(&userData, scanner)
		case "7":
			// Withdraw all investment
			withdrawInvestment(&userData, scanner)
		case "8":
			// Logout
			return
		case "9":
			// Exit
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid operation!. Choose again")
		}
	}
}

func withdrawInvestment(userData *UserBankData, scanner *bufio.Scanner) {
	for {
		fmt.Println("Do you want to withdraw all the investment?")
		fmt.Println("1. Yes\n2. No")
		input := scanText(scanner)
		switch input {
		case "1":
			userData.InvestFund.InvestBalance = userData.InvestFund.InvestBalance.Add(userData.InvestFund.LowRisk).Add(userData.InvestFund.MediumRisk).Add(userData.InvestFund.HighRisk)
			userData.InvestFund.LowRisk = decimal.NewFromFloat(0)
			userData.InvestFund.MediumRisk = decimal.NewFromFloat(0)
			userData.InvestFund.HighRisk = decimal.NewFromFloat(0)
		case "2":
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func transferAccount(userData *UserBankData, scanner *bufio.Scanner) {
	for {
		fmt.Println("Choose what you want to transfer:")
		fmt.Println("1. Transfer from saving account to investment account")
		fmt.Println("2. Transfer from investment account to saving account")
		fmt.Println("3. Exit")
		transferType := scanText(scanner)
		switch transferType {
		case "1":
			transferAmountBwtAccount(&userData.Balance, &userData.InvestFund.InvestBalance, scanner)
		case "2":
			transferAmountBwtAccount(&userData.InvestFund.InvestBalance, &userData.Balance, scanner)
		case "3":
			return
		default:
			fmt.Println("Invalid! Please try again")
		}
	}
}

func transferAmountBwtAccount(from *decimal.Decimal, to *decimal.Decimal, scanner *bufio.Scanner) {
	for {
		fmt.Print("Enter the amount:")
		amount, err:=decimal.NewFromString(scanText(scanner))
		if err != nil {
			fmt.Println("Invalid input")
		} else if amount.GreaterThan(*from) {
			fmt.Println("Insufficient fund")
		} else {
			*from = from.Sub(amount)
			*to = to.Add(amount)
			break
		}
	}
}

func investMoney(userData *UserBankData, scanner *bufio.Scanner) {
	fmt.Println("Investment account:", userData.InvestFund.InvestBalance)
	fmt.Print(`
1. LOW_RISK
2. MEDIUM_RISK
3. HIGH_RISK
4. Exit
`)
	OUTERLOOP:
	for {
		fmt.Print("Choose the investment: ")
		invest:=scanText(scanner)
		switch invest {
		case "1":
			investFund(scanner, &userData.InvestFund.InvestBalance, &userData.InvestFund.LowRisk)
			break OUTERLOOP
		case "2" :
			investFund(scanner, &userData.InvestFund.InvestBalance, &userData.InvestFund.MediumRisk)
			break OUTERLOOP
		case "3" :
			investFund(scanner, &userData.InvestFund.InvestBalance, &userData.InvestFund.HighRisk)
			break OUTERLOOP
		case "4":
			return
		default:
			fmt.Println("Invalid fund! Please try again!")
		}
	}
}

func investFund(scanner *bufio.Scanner, investBalance *decimal.Decimal, investFund *decimal.Decimal) {
	for {
		fmt.Print("Enter the amount you want to invest:") 
		amount, err:= decimal.NewFromString(scanText(scanner))
		if err != nil {
			fmt.Println("Invalid input")
		} else if amount.GreaterThan(*investBalance) {
			fmt.Println("Insufficient fund")
		} else {
			*investBalance = investBalance.Sub(amount)
			*investFund = investFund.Add(amount) 
			break
		}
	}
	
}

func sendMoneytoPpl(userData *UserBankData, scanner *bufio.Scanner) {
	fmt.Print("Enter the username you want to send money to or type 0 to return to menu: ")
	found := false
	OUTERLOOP:
	for{
		user := scanText(scanner) 
		if user == "0" {
			return
		} else if user == userData.Username {
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
		fmt.Print("Enter the amount you want to withdraw or type 0 to return: ")
		withdraw, err := decimal.NewFromString(scanText(scanner))
		if withdraw.Equal(decimal.NewFromFloat(0)){
			return
		}
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
	fmt.Print(`
1. Saving account
2. Investment account
0. Exit
Enter the account to deposit: `)
	account := scanText(scanner)
	if account == "0" {
		return
	}

	for {
		fmt.Print("Enter the amount you want to deposit: ")
		deposit, err := decimal.NewFromString(scanText(scanner))
		if err != nil {
			fmt.Println("Invalid deposit")
		} else if deposit.GreaterThan(userData.Cash) {
			fmt.Println("Invalid cash fund!")
		} else {	
			if account == "1" {
				userData.Balance = deposit
				userData.Cash = userData.Cash.Sub(deposit)
				break
			} else if account == "2" {
				userData.InvestFund.InvestBalance = deposit
				userData.Cash = userData.Cash.Sub(deposit)
				break
			}
		}
	}
}

func showBalance(userData *UserBankData) {
	// saving balance
	interestBalance := userData.Balance.Mul(decimal.NewFromFloat(0.01))
	userData.Balance = userData.Balance.Add(interestBalance)
	// low risk investment balance
	lowRiskBalance := userData.InvestFund.LowRisk.Mul(decimal.NewFromFloat(0.02))
	userData.InvestFund.LowRisk = userData.InvestFund.LowRisk.Add(lowRiskBalance)
	// medium risk investment balance
	mediumRiskBalance := userData.InvestFund.MediumRisk.Mul(decimal.NewFromFloat(0.05))
	userData.InvestFund.MediumRisk = userData.InvestFund.MediumRisk.Add(mediumRiskBalance)
	// high risk investment balance
	highRiskBalance := userData.InvestFund.HighRisk.Mul(decimal.NewFromFloat(0.1))
	userData.InvestFund.HighRisk = userData.InvestFund.HighRisk.Add(highRiskBalance)

	fmt.Println("Your saving balance is: $", userData.Balance)
	fmt.Println("Your investment balance is: $", userData.InvestFund.InvestBalance)
	fmt.Println("Low_Risk investment fund: $", userData.InvestFund.LowRisk)
	fmt.Println("Medium_Risk investment fund: $", userData.InvestFund.MediumRisk)
	fmt.Println("High_Risk investment fund: $", userData.InvestFund.HighRisk)
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