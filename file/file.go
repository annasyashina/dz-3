package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type FileDb struct {
	filename string
}

func (f *FileDb) Read() ([]byte, error) {
	ext := filepath.Ext(f.filename)
	if ext == ".json" {
		data, err := os.ReadFile(f.filename)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(string(data))
		return data, nil
	}
	return nil, errors.New("Файл не является файлом json")
}

func (f *FileDb) Write(content []byte) {

	ext := filepath.Ext(f.filename)
	if ext == ".json" {
		file, err := os.Create(f.filename)
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
func NewFileDb(filename string) *FileDb {
	return &FileDb{
		filename: filename,
	}
}

/*func ReadFile(name string) ([]byte, error) {
	ext := filepath.Ext(name)
	if ext == ".json" {
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
	if ext == ".json" {
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
*/
