package account

import (
	"encoding/json"
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

func (v *VaultDB) AddAccount(acc *Account, fileName string) error {
	v.Accounts = append(v.Accounts, *acc)
	v.UpdateAt = time.Now()
	return v.saveData()
}

func (v *VaultDB) DelAccount(url, fileName string) (int, error) {
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

type VaultDB struct {
	Vault
	db DataBase
}

func (v *VaultDB) loadDate() error {
	byteBuf, err := v.db.Read()
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteBuf, &v.Vault)
	return err
}

func (v *VaultDB) saveData() error {
	byteBuf, err := json.Marshal(v.Vault)
	if err != nil {
		return err
	}

	err = v.db.Write(byteBuf)
	return err
}

func NewVault(db DataBase) (*VaultDB, error) {
	newVault := VaultDB{
		Vault: Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		},
		db: db,
	}

	err := newVault.loadDate()
	if err != nil {
		return &newVault, err
	}

	return &newVault, nil
}
