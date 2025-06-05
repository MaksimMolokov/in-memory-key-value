package main

import (
	"bufio"
	"fmt"
	"os"

	"db/internal/compute"
	"db/internal/storage"
	computePkg "db/pkg/compute"

	"go.uber.org/zap"
)

// Домашнее задание: комментарий 1
func main() {
	// Инициализация логгера
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("ошибка инициализации логгера: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Инициализация компонентов
	engine := storage.NewEngine(logger)
	parser := compute.NewParser(logger)

	// Создание сканера для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в in-memory key-value базу данных!")
	fmt.Println("Доступные команды: SET, GET, DEL")
	fmt.Println("Для выхода введите 'exit'")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "exit" {
			break
		}

		// Парсинг команды
		cmd, err := parser.Parse(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		// Выполнение команды
		switch c := cmd.(type) {
		case *computePkg.SetCommand:
			if err := engine.Set(c.Key, c.Value); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Println("OK")
			}
		case *computePkg.GetCommand:
			if value, err := engine.Get(c.Key); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Println(value)
			}
		case *computePkg.DelCommand:
			if err := engine.Delete(c.Key); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Println("OK")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения ввода: %v\n", err)
	}
}

// Домашнее задание: комментарий 2
