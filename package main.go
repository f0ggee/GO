package main

import (
	"bufio"
	"math/rand"
	"strconv"

	"fmt"
	"os"
	"path/filepath"

	"strings"
	"time"
)

func password() {

	right := "123445"
	for attempts := 0; attempts < 100; attempts++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter your password:")
		scanner.Scan()
		answer := scanner.Text()

		if strings.ToLower(answer) == right {
			fmt.Println("Your right")
			time.Sleep(1 * time.Second)
			return
		} else {
			fmt.Println("Your lose")

		}

	}
	fmt.Println("your lose ")
}

func reade() string {
	var name string

	fmt.Println("Hi thise your note for your data")
	time.Sleep(2 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type your login please")
	scanner.Scan()
	name = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка:", err)
	}
	return string(name)

}

func passworde() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Do you want to generate a password? yes/no")
		scanner.Scan()
		text := strings.ToLower(strings.TrimSpace(scanner.Text())) // Приводим ввод к нижнему регистру и убираем лишние пробелы

		// Генерация пароля
		if text == "yes" {
			fmt.Println("What length do you want?")
			scanner.Scan()
			texte := scanner.Text()

			textint, err := strconv.Atoi(texte)
			if err != nil || textint <= 0 {
				fmt.Println("Error: please enter a valid positive number.")
				return "" // Возвращаемся к запросу
			}

			rand.Seed(time.Now().UnixNano())
			chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()"
			var password []rune

			for i := 0; i < textint; i++ {
				password = append(password, rune(chars[rand.Intn(len(chars))]))
			}

			return string(password)
		}

		// Пользователь вводит свой пароль
		if text == "no" {
			fmt.Println("Type your password:")
			scanner.Scan()
			password := scanner.Text()

			// Возвращаем введённый пароль
			return password
		}

		// Если пользователь ввёл что-то другое
		fmt.Println("Please type 'yes' or 'no'.")
	}
}

func main() {
	password()
	///passworde := passworde()

	name := reade()

	password := passworde()

	fmt.Println("Your name and age will be written in the file soon")

	time.Sleep(1 * time.Second)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filePath := filepath.Join(homeDir, "Desktop", "Your_data.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("thise your login:" + name))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = file.Write([]byte("\n"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = file.Write([]byte("Thise yout password:" + password))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Successfully wrote the name and age to the file.\n")
}
