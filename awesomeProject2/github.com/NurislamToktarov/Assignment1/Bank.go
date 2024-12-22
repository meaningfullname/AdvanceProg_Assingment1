package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (account *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		account.Balance += amount
		fmt.Printf("Deposited: %.2f\n", amount)
	} else {
		fmt.Println("Deposit amount must be positive.")
	}
}

func (account *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive.")
		return
	}
	if account.Balance >= amount {
		account.Balance -= amount
		fmt.Printf("Withdrew: %.2f\n", amount)
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (account *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", account.Balance)
}

func main() {
	account := BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       0,
	}

	var choice int
	var amount float64
	//hello
	fmt.Println("Welcome to the Bank Account System!")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exiting. Thank you for using the system!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
