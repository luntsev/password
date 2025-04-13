package account

import (
	"encoding/json"
	"password/files"
	"time"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func (v *Vault) loadDate(fileName string) error {
	byteBuf, err := files.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteBuf, v)
	return err
}

func (v *Vault) saveData(fileName string) error {
	byteBuf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = files.WriteFile(byteBuf, fileName)
	return err
}

func (v *Vault) AddAccount(acc *Account, fileName string) error {
	v.Accounts = append(v.Accounts, *acc)
	v.UpdateAt = time.Now()
	return v.saveData(fileName)
}

func (vault *Vault) DelAccount(url, fileName string) (int, error) {
	count := 0
	var newVault Vault
	for _, account := range vault.Accounts {
		if account.Url != url {
			newVault.Accounts = append(newVault.Accounts, account)
		} else {
			count++
		}
	}
	newVault.UpdateAt = time.Now()
	if count != 0 {
		vault = &newVault
	}
	err := vault.saveData(fileName)
	return count, err
}

func NewVault(fileName string) (*Vault, error) {
	newVault := Vault{
		Accounts: []Account{},
		UpdateAt: time.Now(),
	}

	err := newVault.loadDate(fileName)
	if err != nil {
		return &newVault, err
	}

	return &newVault, nil
}
