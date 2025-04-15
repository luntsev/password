package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"password/output"
)

type Encrypter struct {
	Key string
}

func (enc *Encrypter) Encrypt(plaintStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		output.PrintError(err, "Неудалось сформировать блок шифрования")
		panic(err)
	}

	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		output.PrintError(err, "Неудалось сформировать шифр")
		panic(err)
	}

	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		output.PrintError(err, "Неудалось сформировать случайный nonce")
		panic(err)
	}

	return aesGSM.Seal(nonce, nonce, plaintStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		output.PrintError(err, "Неудалось сформировать блок шифрования")
		panic(err)
	}

	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		output.PrintError(err, "Неудалось сформировать шифр")
		panic(err)
	}

	nonceSize := aesGSM.NonceSize()
	nonce, cipherStr := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plaintStr, err := aesGSM.Open(nil, nonce, cipherStr, nil)
	if err != nil {
		output.PrintError(err, "Неудалось расшифровать данные")
		panic(err)
	}
	return plaintStr
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("В переменных окружения не задан ключ шифрования")
	}
	return &Encrypter{Key: key}
}
