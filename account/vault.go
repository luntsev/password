package account

import (
	"encoding/json"
	"password/encrypter"
	"password/output"
	"time"
)

type DataBase interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func (v *Vault) FindAccount(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range v.Accounts {
		if checker(account, str) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

type VaultDB struct {
	Vault
	db  DataBase
	enc encrypter.Encrypter
}

func (v *VaultDB) AddAccount(acc *Account) error {
	v.Accounts = append(v.Accounts, *acc)
	v.UpdateAt = time.Now()
	return v.saveData()
}

func (v *VaultDB) DelAccount(url string) (int, error) {
	count := 0
	var newAccounts []Account
	for _, account := range v.Vault.Accounts {
		if account.Url != url {
			newAccounts = append(newAccounts, account)
		} else {
			count++
		}
	}
	if count != 0 {
		v.Vault.Accounts = newAccounts
		v.Vault.UpdateAt = time.Now()
	}
	err := v.saveData()
	return count, err
}

func (v *VaultDB) saveData() error {
	plaintStr, err := json.Marshal(v.Vault)
	if err != nil {
		output.PrintError(err, "Неудалось закодировать данные в JSON")
		return err
	}
	encryptedStr := v.enc.Encrypt(plaintStr)
	err = v.db.Write(encryptedStr)
	return err
}

func NewVault(db DataBase, enc encrypter.Encrypter) (*VaultDB, error) {
	file, err := db.Read()

	newVaultDB := VaultDB{
		Vault: Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		},
		db:  db,
		enc: enc,
	}

	if err != nil {
		output.PrintError(err, "Не удалось прочесть файл")
		return &newVaultDB, err
	}

	plaintStr := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(plaintStr, &vault)
	if err != nil {
		output.PrintError(err, "Не удалось разобрать JSON")
		return &newVaultDB, err
	}

	newVaultDB.Vault = vault
	return &newVaultDB, nil
}
