package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	rnd "math/rand/v2"
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

func (enc *Encrypter) generateKey() (string, error) {
	var validChar = []rune("1234567890abcdef")
	key := make([]rune, 32)
	for i := range key {
		key[i] = validChar[rnd.IntN(16)]
	}
	enc.Key = string(key)

	file, err := os.Create(".env")
	if err != nil {
		output.PrintError(err, "Не удалось создать файл .env")
		return "", err
	}

	content := []byte(fmt.Sprintf("KEY=%s", enc.Key))
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err, "Не удалось записать в файл .env")
		return "", err
	}
	return enc.Key, nil
}

func NewEncrypter() *Encrypter {
	enc := &Encrypter{}
	key := os.Getenv("KEY")
	if key == "" {
		err := errors.New("no \"KEY\" in environment variables")
		output.PrintError(err, "В переменных окружения нет ключа шифрования")
		key, _ = enc.generateKey()
	}
	enc.Key = key
	return enc
}
