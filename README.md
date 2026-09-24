# GreenDayBank
Green Day Bank
You will implement a simple command-line banking application where users can perform basic banking operations.

The Situation
How nice it would be if your savings and investments in your bank would only grow and never shrink. In this task you'll have an opportunity to build such a magical bank.

Functional Requirements
User Interaction
The application uses the `bufio.Scanner` from the `bufio` package for all user input. Handle EOF (Ctrl+D) gracefully by breaking out of input loops when `scanner.Scan()` returns false. The system must handle invalid inputs gracefully, displaying appropriate error messages for scenarios like non-existent usernames or insufficient funds.

Login System: Users must log in by entering their names. The system supports four users: Alice, Bob, Charlie, and Diana, who can all log in and send money to each other.

Menu System: After logging in, users are presented with a menu of options in this exact order:
```
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
```
Users select options by entering the corresponding number and pressing Enter. After each operation completes, the menu is displayed again for the next action.

Account Management
Each user starts with $1000 in cash that is not yet deposited into the bank. Users have two accounts: a savings account and an investment account, both initially empty. All account balances must be represented using the `decimal.Decimal` type from the `shopspring/decimal` package for precise financial calculations.

Interest and Balance Viewing: The bank provides a 1% interest rate on savings accounts. Every time users view their account balance, the interest is automatically calculated and added to their savings account balance.

Account Operations: Users can deposit money from their cash into their savings account (cash can only be deposited to the savings account). They can withdraw money from their savings account back to cash, send money from their savings account to another user, or transfer money between their savings and investment accounts.

Investment Fund Management
Users can invest in predefined funds using money from their investment account. The bank offers three funds:

LOW_RISK (appreciates at 2%)
MEDIUM_RISK (appreciates at 5%)
HIGH_RISK (appreciates at 10%).
The gain is calculated and added to the money in the fund every time the user views their account balance.

Users can withdraw all their investments at any time, which moves all fund money back to their investment account balance.

Session Management
Persistent User State: Once a user logs in, their session remains active until they explicitly log out or exit the application.
Logout: Users can log out and log in with a different user.
Exit: Users can exit the application.
NB! All the exact implementation details are not revealed in the task description. The complete requirements will be revealed by automated tests and the video below.

Possible Project Structure
```
simpleBank
├── main.go           // Entry point, initializes banking service
├── model/            // Contains data structures and core logic
│   ├── account.go    // Account interface and base struct
│   ├── fund.go       // Investment fund types and growth factors
│   ├── investment_account.go
│   ├── savings_account.go
│   └── user.go       // User entity and methods
└── service/          // Handles banking operations and user interaction
    └── banking_service.go // Manages CLI and banking functionality
```
`main.go` needs to be located at the root of your `simpleBank` repository. This is the entry point of the application, initializing and starting the banking service. Your Go module name should match your repository name: `simpleBank`.

Video Example
Video: Banking app demo

Bonus Functionality
You are welcome to implement additional bonus features after successfully completing the task.

Useful Links
Go Tour
Go bufio Package
Go Error Handling
shopspring/decimal Package
Go Structs, Methods and Receivers
Go Maps
