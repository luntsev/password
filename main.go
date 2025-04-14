package main

import (
	"fmt"
	"password/account"
	"password/files"
	"strings"

	"github.com/fatih/color"
)

var vault *account.VaultDB

func main() {
	vault, err := account.NewVault(*files.NewJsonDB("accounts.json"))
	if err != nil {
		color.Red(err.Error())
	}
Menu:
	for {
		choise := getMenu()
		switch choise {
		case "1":
			err := createAccount(vault)
			if err != nil {
				color.Red(err.Error())
			}
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault, "accounts.json")
		case "4":
			break Menu
		default:
			color.Red("%s - неизвестная комманда", choise)
		}
	}
}

func promtData(promt string) string {
	var result string
	fmt.Print(promt + ": ")
	fmt.Scanln(&result)
	return result
}

func getMenu() string {
	fmt.Println("-=<Менеджер паролей>=-")
	color.Blue("1 - создать аккаунт")
	color.Green("2 - найти аккаунт")
	color.Red("3 - удалить аккаунт")
	color.Yellow("4 - выход")
	return promtData("Выберите действие")
}

func createAccount(vault *account.VaultDB) error {
	login := promtData("Введите логин (email)")
	password := promtData("Введите пароль (если не задан система сгенерирует свой)")
	url := promtData("Введите URL ресурса")
	newAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		return err
	}
	err = vault.AddAccount(newAccount, "accounts.json")
	return err
}

func findAccount(vault *account.VaultDB) {
	url := promtData("Введите URL искомого ресурса")
	isFinded := false
	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			color.Cyan("Для URL - %s найден следующий акаунт: login: %s, password: %s, URL: %s", url, account.Login, account.Password, account.Url)
			isFinded = true
		}
	}
	if !isFinded {
		color.Red("Ничего не найдено!")
	}
	fmt.Println()
}

func deleteAccount(vault *account.VaultDB, fileName string) {
	url := promtData("Введите URL удаления")
	delCount, err := vault.DelAccount(url, fileName)
	if err == nil {
		color.Red("Удалено %d записей", delCount)
	} else {
		color.Red("Найдено %d записей для удаления. При удалении возникла ошибка: %s", delCount, err.Error())
	}
}
