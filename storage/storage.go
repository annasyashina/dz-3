package storage

import (
	"fmt"
	"os"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type StorageDb struct {
	filename string
}

func (f *StorageDb) Read() ([]byte, error) {
	//	return []byte("file content"), nil
	data, err := os.ReadFile(f.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}

func (f *StorageDb) Write(content []byte) {

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
}
func NewStorageDb(filename string) *StorageDb {
	return &StorageDb{
		filename: filename,
	}
}

/*func ReadFile(name string) ([]byte, error) {
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
*/
