package files

import (
	"os"
	"password/output"
)

type JsonDB struct {
	fileName string
}

func (db *JsonDB) Read() ([]byte, error) {
	file, err := os.ReadFile(db.fileName)
	if err != nil {
		output.PrintError(err, "Не удалось произвести чтение из файла")
		return nil, err
	}
	return file, nil
}

func (db *JsonDB) Write(content []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		output.PrintError(err, "Не удалось создать файл")
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err, "Не удалось записать в файл")
		return err
	}
	return nil
}

func NewJsonDB(fileName string) *JsonDB {
	return &JsonDB{fileName: fileName}
}
