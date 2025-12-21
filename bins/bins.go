package bins

import (
	"encoding/json"
	"fmt"

	//"struct/list/bins"

	"struct/list/file"

	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList = []Bin

func WriteBinList(binList BinList, db file.Db) {

	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка преобразования")
		return
	}
	//var d file.Db = file.NewFileDb("storage.json")
	db.Write(data)
	//file.WriteFile(data, "storage.json")
}

func ReadBinList(binList BinList, db file.Db) BinList {
	//var d file.Db = file.NewFileDb("storage.json")
	data, err := db.Read()

	//file, err := file.ReadFile("storage.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return binList
	}

	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println("Ошибка разбора json файла")
		return binList
	}
	return binList
}
