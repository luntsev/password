package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"password/output"
	"regexp"
	"time"
)

type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createDate"`
	UpdateAt time.Time `json:"updateDate"`
}

func (acc *Account) GetLogin() string {
	return acc.Login
}

func (acc *Account) GetPassword() string {
	return acc.Password
}

func (acc *Account) GetURL() string {
	return acc.Url
}

func (acc *Account) GetCreateDate() time.Time {
	return acc.CreateAt
}

func (acc *Account) GetUpdateDate() time.Time {
	return acc.UpdateAt
}

func NewAccount(accLogin, accPassword, accUrl string) (*Account, error) {
	emailRe := regexp.MustCompile(`[A-z0-9-_]+@[A-z0-9-_]+\.[A-z]+`)
	if !emailRe.MatchString(accLogin) {
		err := errors.New("invalid login")
		output.PrintError(err, "Введен некорректный логин")
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(accUrl)
	if err != nil {
		output.PrintError(err, "Введен некорректный URL")
		return nil, err
	}

	if accPassword == "" {
		accPassword = generatePassword(12)
	}
	newAcc := Account{
		Login:    accLogin,
		Password: accPassword,
		Url:      accUrl,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	return &newAcc, nil
}

var validChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!@#$%&*-_+^")

func generatePassword(length int) string {
	password := make([]rune, length)
	for i := range password {
		password[i] = validChars[rand.IntN(len(validChars))]
	}
	return string(password)
}
