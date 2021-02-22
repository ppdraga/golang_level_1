package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func outOfRange() {
	defer func() {
		if v := recover(); v != nil {
			t := time.Now()
			fmt.Printf("recovered at %v!\n", t)
		}
	}()
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[6])
}

func main() {
	// 1. Напишите программу, в которой неявно будет срабатывать паника.
	// Сделайте отложенную функцию, которая будет обрабатывать эту панику
	// и печатать предупреждение в консоль. Критерий выполнения
	// задания — программа не завершается аварийно.
	// 2. Дополните программу собственной ошибкой, хранящей время её возникновения.
	outOfRange()
	fmt.Println("program is going on!")

	// 3. Напишите функцию, которая создаёт файл в файловой системе и
	// использует отложенный вызов функций для безопасного закрытия файла.
	nf, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer nf.Close()
}
