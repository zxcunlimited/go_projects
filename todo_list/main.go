package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var list []string
	var choice string
	fmt.Println("Welcome to the TO DO List CLI app!")
	fmt.Println()
	for {
		fmt.Println("Enter your command (create, read, update, delete):")
		choice, _ = bufio.NewReader(os.Stdin).ReadString('\r')
		choice = strings.Replace(choice, "\r", "", 1)
		switch choice {
		case "create":
			fmt.Println("Enter task name:")
			task, _ := bufio.NewReader(os.Stdin).ReadString('\r')
			task = strings.Replace(task, "\r", "", 1)
			list = append(list, task)
		case "read":
			for i, cur_task := range list {
				fmt.Printf("%d. %s\n", i+1, cur_task)
			}
		case "update":
			flag := true
			for flag {
				fmt.Println("Enter task name to update:")
				task, _ := bufio.NewReader(os.Stdin).ReadString('\r')
				task = strings.Replace(task, "\r", "", 1)
				for i, list_task := range list {
					if list_task == task {
						flag = false
						fmt.Println("Enter new task name:")
						new_task, _ := bufio.NewReader(os.Stdin).ReadString('\r')
						new_task = strings.Replace(new_task, "\r", "", 1)
						list[i] = new_task
						fmt.Printf("Updated task #%d with name \"%s\" successfully!\n", i, new_task)
						break
					}
				}
			}
		case "delete":
			flag := true
			for flag {
				fmt.Println("Enter task name to remove:")
				task, _ := bufio.NewReader(os.Stdin).ReadString('\r')
				task = strings.Replace(task, "\r", "", 1)
				for i, list_task := range list {
					if list_task == task {
						flag = false
						fmt.Printf("Removed task #%d with name \"%s\" successfully!\n", i, list_task)
						// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
						copy(list[i:], list[i+1:])

						// 2. Удалить последний элемент (записать нулевое значение).
						list[len(list)-1] = ""

						// 3. Усечь срез.
						list = list[:len(list)-1]
						break
					}
				}
			}
		default:
			fmt.Println("Invalid command! Please, try again!")
		}
	}
}
