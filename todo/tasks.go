package todo

import (
	"encoding/json"
	"log"
	"os"
)

type Task struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
}

var tasks []Task //Инициализируем слайс структур Task
var nextID = 1   // Инициализация ID

func LoadTasks() error { // Функция для загрузки задач из JSON
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	return nil
}

func SaveTasks() error { //Функция для сохранения слайса задач
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func AddTask(description string) { //Функция для добавления новой задачи
	log.Printf("Добавление задачи с описанием %s", description)
	task := Task{ID: nextID, Description: description, Completed: false} //Создаем новый экземпляр
	tasks = append(tasks, task)                                          //Добавляем его к слайсу
	nextID++                                                             //Обновляем счетчик для следующей задачи

	err := SaveTasks()
	if err != nil {
		log.Printf("Ошибка сохранения задач %v", err)
	}

	log.Printf("Задачи добавлена с ID %d", task.ID)
}

func GetAllTasks() []Task { //Возвращаем слайс с задачами
	log.Printf("Возвращение списка задач: %d задач", len(tasks))

	return tasks
}

func CompleteTask(id int) bool {
	for index, task := range tasks { //Пробегаемся по слайсу задач
		if task.ID == id { //Если id задачи совпадает с той, которой мы хотим завершить
			tasks[index].Completed = true //То делаем её завершенной
			err := SaveTasks()            //Сохраняем слайс задач
			if err != nil {
				log.Printf("Ошибка выполнения задачи %v: %v", task.ID, err)
			}
			return true
		}
	}
	return false //Если не совпадает id, сразу же возвращаем false
}

func DeleteTask(id int) bool {
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			err := SaveTasks()
			if err != nil {
				log.Printf("Ошибка удаления задач %v", err)
			}
			return true
		}
	}
	return false
}
