package main

import (
	"fmt"
	"password/account"
	"password/files"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultDB){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	vault, err := account.NewVault(files.NewJsonDB("accounts.json"))
	if err != nil {
		color.Red(err.Error())
	}

	for {
		choise := promtData([]string{
			"-=<Менеджер паролей>=-",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите действие",
		})
		currFunc := menu[choise]
		if currFunc == nil {
			break
		}
		currFunc(vault)
	}
}

func promtData[T any](promt []T) string {
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
	login := promtData([]string{"Введите логин (email)"})
	password := promtData([]string{"Введите пароль (если не задан система сгенерирует свой)"})
	url := promtData([]string{"Введите URL ресурса"})
	newAccount, _ := account.NewAccount(login, password, url)
	vault.AddAccount(newAccount)
}

func findAccount(vault *account.VaultDB) {
	url := promtData([]string{"Введите URL искомого ресурса"})
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

func deleteAccount(vault *account.VaultDB) {
	url := promtData([]string{"Введите URL удаления"})
	delCount, err := vault.DelAccount(url)
	if err == nil {
		color.Red("Удалено %d записей", delCount)
	} else {
		color.Red("Найдено %d записей для удаления. При удалении возникла ошибка: %s", delCount, err.Error())
	}
}
