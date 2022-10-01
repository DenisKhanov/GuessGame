package main

import (
	"fmt"
	"math/rand"
	"time"
)

var answer = "\nСожалею, но ты не отгадал(а)! В"
var numbers []int

func main() {
	rand.Seed(time.Now().Unix())
	randomNumber := rand.Intn(100) + 1
	fmt.Println("У тебя есть 7 попыток на то, чтобы отгадать случайно генерируемое число :)\nВведи число от 1 до 100, которое ты предполагаешь!")
	var guess int
	for i := 7; i > 0; i-- {
		for {
			_, err := fmt.Scanf("%d\n", &guess)
			if err != nil {
				fmt.Println("Введи пожалeйста именно число!")
			} else {
				break
			}
		}
		for _, i := range numbers {
			if guess == i {
				fmt.Printf("Не очень умный ход, ведь ты уже пробовал число %d!\n", guess)
			}
		}
		numbers = append(numbers, guess)

		if guess > randomNumber {
			fmt.Printf("Твое число, %d больше загаданного!", guess)
			fmt.Printf(" Осталось %d попыток ;)\n", i-1)
		} else if guess < randomNumber {
			fmt.Printf("Твое число, %d меньше загаданного!", guess)
			fmt.Printf(" Осталось %d попыток ;)\n", i-1)
		} else {
			answer = "\nПоздравляю, ты отгадал(а) в"
			break
		}

	}
	fmt.Print(answer)
	fmt.Printf("ерное число %d!\n\n", randomNumber)
	fmt.Println("Жамкай Enter для выхода!")
	fmt.Scanf("%d\n", &guess)
}
