package cmd

import (
	"cli-todo/todo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var addTaskCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Добавить новую задачу",
	Long:  "Добавление новой задачи с описанием",
	Args:  cobra.ExactArgs(1), //Всего один аргумент
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		todo.AddTask(description)
		fmt.Println("Задача добавлена:", description)
	},
}

var completeTask = &cobra.Command{
	Use:   "complete [id]",
	Short: "Завершить задачу по ID",
	Long:  "Завершение ранее добавленной задачи",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Неверный ID!")
			return
		}

		if todo.CompleteTask(ID) {
			fmt.Printf("Задача с ID %d завершена!", ID)
		} else {
			fmt.Println("Задача с таким ID не найдена")
		}
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete {id}",
	Short: "Удалить задачу по ID",
	Long:  "Удаление существующей задачи по ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Неверный ID!")
			return
		}

		if todo.DeleteTask(ID) {
			fmt.Printf("Задача с ID %d удалена!", ID)
		} else {
			fmt.Println("Задача с таким ID не найдена")
		}

	},
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список задач",
	Long:  "Получение списка добавленные задач",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := todo.GetAllTasks()
		if len(tasks) == 0 {
			fmt.Println("Задач нет")
			return
		}

		fmt.Printf("%-5s %-20s %-15s\n", "ID", "Описание", "Статус")
		fmt.Println(strings.Repeat("-", 45))
		for _, task := range tasks {
			status := "не завершена"
			if task.Completed {
				status = "завершена"
			}
			fmt.Printf("%-5d %-20s %-15s\n", task.ID, task.Description, status)

		}
	},
}

func init() {
	err := todo.LoadTasks()
	if err != nil {
		log.Printf("Ошибка загрузки задач %v", err)
	}
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
	rootCmd.AddCommand(listTasksCmd)
	rootCmd.AddCommand(completeTask)
}
