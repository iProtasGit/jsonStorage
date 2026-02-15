package controller

import (
	"fmt"
	
	"jsonStorage/config"
	"jsonStorage/internal/usecase"
	"jsonStorage/internal/adapter"
)

func StartCLI(cfg *config.Config, storage *adapter.Storage) {
	uc := usecase.NewUseCase(storage)
	for {
		fmt.Print("Welcome to Json Storage by iProtas!\n\n\n")
		fmt.Println("Please select an option:")
		fmt.Println("1. Create a new bin")
		fmt.Println("2. View existing bins")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			id, private, name, err := GetDataFromUser()
			if err != nil {
				fmt.Println(err)
				return
			}
			uc.CreateBin(id, private, name)
		case 2:
			uc.GetBins()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func GetDataFromUser() (string, bool, string, error) {
	var id, name string
	var private bool

	fmt.Print("Enter bin ID: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return "", false, "", err
	}

	fmt.Print("Is the bin private? (true/false): ")
	_, err = fmt.Scanln(&private)
	if err != nil {
		return "", false, "", err
	}

	fmt.Print("Enter bin name: ")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return "", false, "", err
	}

	return id, private, name, nil
}