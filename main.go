package main

import (
	"fmt"
	"strconv"
	"struct/list/bins"
	"time"
)

func main() {
	binList := bins.BinList{}

	fmt.Println("Приложение для структуры")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBinList(binList)
		case 2:
			binList = addBinList(binList)
		case 3:
			binList = deleteBinList(binList)
		case 4:
			writeBinList(binList)
		case 5:
			binList = readBinList(binList)
		case 6:
			break Menu
		}
	}
}

func readBinList(binList bins.BinList) bins.BinList {
	/*file, err := storage.ReadFile("storage.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return binList
	}

	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println("Ошибка разбора json файла")
		return binList
	}
	return binList*/
	return bins.ReadBinList(binList)
}

func writeBinList(binList bins.BinList) {

	bins.WriteBinList(binList)
	/*data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка преобразования")
		return
	}
	//fmt.Println(string(data))
	//data, err := binList

	storage.WriteFile(data, "storage.json")*/
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть структуру")
	fmt.Println("2. Добавить структуру")
	fmt.Println("3. Удалить структуру")
	fmt.Println("4. Сохранить структуру в файл")
	fmt.Println("5. Прочитать структуру из файла")
	fmt.Println("6. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBinList(binList bins.BinList) {
	if len(binList) == 0 {
		fmt.Println("Пока нет структур")
	}
	//fmt.Println(binList)
	for key, value := range binList {
		fmt.Println(key, ": ", value)
	}
}

func addBinList(binList bins.BinList) bins.BinList {
	var newId string
	var newPrivate bool
	var newCreated time.Time
	var newName string
	var input string
	fmt.Println("Введите id")
	fmt.Scan(&newId)
	fmt.Println("Введите private (true или false)")
	fmt.Scan(&input)

	newPrivate, err := strconv.ParseBool(input)
	if err != nil {
		fmt.Println("Неправильный ввод private")
		return binList
	}

	fmt.Println("Введите время")
	fmt.Scan(&input)

	newCreated, err2 := time.Parse("15:04:05", input)
	if err2 != nil {
		fmt.Println("Неправильный ввод time", err2)
		return binList
	}

	fmt.Println("Введите name")
	fmt.Scan(&newName)

	newBin := bins.Bin{
		Id:        newId,
		Private:   newPrivate,
		CreatedAt: newCreated,
		Name:      newName,
	}
	binList = append(binList, newBin)
	//fmt.Println(binList)

	return binList
}

func deleteBinList(binList bins.BinList) bins.BinList {
	var deleteId string
	fmt.Println("Введите id")
	fmt.Scan(&deleteId)
	for index, value := range binList {
		if value.Id == deleteId {
			binList[index] = binList[len(binList)-1]
			return binList[:len(binList)-1]
		}
	}
	return binList
}
