package files

import "os"

type JsonDB struct {
	fileName string
}

func (db *JsonDB) Read() ([]byte, error) {
	file, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (db *JsonDB) Write(content []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func NewJsonDB(fileName string) *JsonDB {
	return &JsonDB{fileName: fileName}
}
