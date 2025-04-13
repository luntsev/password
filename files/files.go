package files

import "os"

func ReadFile(name string) ([]byte, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}
