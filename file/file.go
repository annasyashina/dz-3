package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	ext := filepath.Ext(name)
	if ext == "json" {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(string(data))
		return data, nil
	}
	return nil, errors.New("Файл не является файлом json")

}

func WriteFile(content []byte, name string) {
	ext := filepath.Ext(name)
	if ext == "json" {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := file.Write(content)
		if err2 != nil {
			defer file.Close()
			fmt.Println(err2)
			return
		}

		fmt.Println("Запись успешна")
	} else {
		fmt.Println("Файл не является файлом json")
	}
}
