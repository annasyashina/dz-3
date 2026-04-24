package storage

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}

func WriteFile(content []byte, name string) {
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
}
