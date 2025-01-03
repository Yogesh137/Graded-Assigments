package main

import (
	"errors"
	"fmt"
)

type Bank struct {
	Name     string
	Address  string
	Accounts []Account
}

type Account struct {
	ID       int
	Name     string
	Balance  float64
	TransLog []string
}

const (
	OptionDeposit     = 1
	OptionWithdraw    = 2
	OptionViewBalance = 3
	OptionHistory     = 4
	OptionExit        = 5
)

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	a.TransLog = append(a.TransLog, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw deducts money from the account balance.
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal Amount should be greater than zero")
	}
	if amount > a.Balance {
		return errors.New("Insufficient Balance")
	}
	a.Balance -= amount
	a.TransLog = append(a.TransLog, fmt.Sprintf("Withdrawn: %.2f", amount))
	return nil
}

func (a *Account) ViewTransactionHistory() {
	if len(a.TransLog) == 0 {
		fmt.Println("No Transactions")
		return
	}
	fmt.Println("Transaction History:")
	for _, transaction := range a.TransLog {
		fmt.Println(transaction)
	}
}

func main() {
	var bank Bank

	fmt.Print("Do you have a bank set up? (yes/no): ")
	var response string
	fmt.Scan(&response)

	if response == "no" {
		fmt.Print("Enter bank name: ")
		fmt.Scan(&bank.Name)
		fmt.Print("Enter bank address: ")
		fmt.Scan(&bank.Address)
		fmt.Println("Bank created successfully!")
	} else {
		bank = Bank{
			Name:    "My Bank",
			Address: "143 My Lane , My City",
			Accounts: []Account{
				{ID: 1, Name: "Alice", Balance: 2000.0},
				{ID: 2, Name: "Bob", Balance: 1000.0},
			},
		}
	}

	fmt.Printf("Welcome to %s!\nAddress: %s\n\n", bank.Name, bank.Address)

	fmt.Print("Do you have an account? (yes/no): ")
	fmt.Scan(&response)

	var currentAccount *Account
	if response == "no" {
		var newAccount Account
		newAccount.ID = len(bank.Accounts) + 1
		fmt.Print("Enter your name: ")
		fmt.Scan(&newAccount.Name)
		fmt.Print("Enter initial deposit amount: ")
		fmt.Scan(&newAccount.Balance)
		newAccount.TransLog = append(newAccount.TransLog, fmt.Sprintf("Account created with initial deposit: %.2f", newAccount.Balance))
		bank.Accounts = append(bank.Accounts, newAccount)
		currentAccount = &bank.Accounts[len(bank.Accounts)-1]
		fmt.Println("Account created successfully! Your account ID is:", newAccount.ID)
	} else {
		var accountID int
		fmt.Print("Enter your account ID: ")
		fmt.Scan(&accountID)
		for i := range bank.Accounts {
			if bank.Accounts[i].ID == accountID {
				currentAccount = &bank.Accounts[i]
				break
			}
		}
		if currentAccount == nil {
			fmt.Println("Invalid account ID.")
			return
		}
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case OptionDeposit:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := currentAccount.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}
		case OptionWithdraw:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := currentAccount.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}
		case OptionViewBalance:
			fmt.Printf("Current Balance: %.2f\n", currentAccount.Balance)
		case OptionHistory:
			currentAccount.ViewTransactionHistory()
		case OptionExit:
			fmt.Println("Thank you for using the Bank Transaction System. Goodbye!")
			break
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		if choice == OptionExit {
			break
		}
	}
}
