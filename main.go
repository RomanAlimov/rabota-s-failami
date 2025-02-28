package main

import (
	"fmt" //fmt для печати текста
	"os"  // работа с файлами и созданием.
)

func main() {
	file, err := os.Create("text.txt") // создал файл текст тxt.

	os.Rename("hello golang", file.Name()) //переименовал файл

	if err != nil { // проверка на ошибку
		fmt.Print("error uncouwn - ", nil) // если ошибка найдена вывести текст и ошибку.
		os.Exit(1)                         // если ошибка - то завершение рпботы терминала.
	}
	defer file.Close() // чтобы закрыть соединение файла.

	dannie := "Text to file\nFrom\nGolang\n\tTabulate" //  текст в файл

	file.WriteString(dannie) // передача текста в файл

	file_data, _ := os.ReadFile(file.Name()) // передача текста в вывод терминала.

	fmt.Println(string(file_data)) // печатает весь текст в добавок на терминал
}
