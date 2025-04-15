package main

import (
	"fmt"
	"password/account"
	"password/encrypter"
	"password/files"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultDB){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		color.Red("Не удалось загрузить переменные окружения из файла .env")
	}
	vault, err := account.NewVault(files.NewJsonDB("accounts.vault"), *encrypter.NewEncrypter())
	if err != nil {
		color.Red(err.Error())
	}

	for {
		choise := promtData(
			"-=<Менеджер паролей>=-",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите действие",
		)
		currFunc := menu[choise]
		if currFunc == nil {
			break
		}
		currFunc(vault)
	}
}

func promtData(promt ...any) string {
	var result string
	for i, val := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", val)
		} else {
			fmt.Println(val)
		}
	}
	fmt.Scanln(&result)
	return result
}

func createAccount(vault *account.VaultDB) {
	login := promtData("Введите логин (email)")
	password := promtData("Введите пароль (если не задан система сгенерирует свой)")
	url := promtData("Введите URL ресурса")
	newAccount, _ := account.NewAccount(login, password, url)
	vault.AddAccount(newAccount)
}

func findAccount(vault *account.VaultDB) {
	choise := promtData(
		"1. Поиск по полю URL",
		"2. Поиск по полю login",
		"По какому полю искать",
	)
	var accounts []account.Account
	switch choise {
	case "1":
		url := promtData("Введите URL")
		accounts = vault.FindAccount(url, func(acc account.Account, str string) bool {
			return strings.Contains(acc.Url, str)
		})
	case "2":
		url := promtData("Введите login")
		accounts = vault.FindAccount(url, func(acc account.Account, str string) bool {
			return strings.Contains(acc.Login, str)
		})
	default:
		color.Red("Недопустимый выбор!")
	}

	if len(accounts) > 0 {
		for _, account := range accounts {
			color.Cyan("Найден акаунт: login: %s, password: %s, URL: %s", account.Login, account.Password, account.Url)
		}
	} else {
		color.Red("Ничего не найдено!")
	}
	fmt.Println()
}

func deleteAccount(vault *account.VaultDB) {
	url := promtData("Введите URL удаления")
	delCount, err := vault.DelAccount(url)
	if err == nil {
		color.Red("Удалено %d записей", delCount)
	} else {
		color.Red("Найдено %d записей для удаления. При удалении возникла ошибка: %s", delCount, err.Error())
	}
}
