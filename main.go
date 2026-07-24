package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Task structure with IsCompleted field
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

const fileName = "tasks.json"

// Load tasks from file
func loadTasks() []Task {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Println("⚠️ Warning: Could not parse tasks.json")
		return []Task{}
	}

	return tasks
}

// Save tasks to file
func saveTasks(tasks []Task) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("❌ Error saving tasks:", err)
		return
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("❌ Error writing file:", err)
	}
}

func main() {
	fmt.Println("=================================")
	fmt.Println("🚀 Welcome to Must Task CLI v2")
	fmt.Println("=================================")

	tasks := loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu Options:")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Mark Task as Complete")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter option (1-5): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n📋 Your Tasks:")
			if len(tasks) == 0 {
				fmt.Println("   No tasks found!")
			}
			for _, task := range tasks {
				status := "[ ]"
				if task.IsCompleted {
					status = "[X]"
				}
				fmt.Printf("   %s %d. %s (Created: %s)\n",
					status,
					task.ID,
					task.Title,
					task.CreatedAt.Format("15:04:05"),
				)
			}

		case "2":
			fmt.Print("Enter new task title: ")
			titleInput, _ := reader.ReadString('\n')
			cleanTitle := strings.TrimSpace(titleInput)

			if cleanTitle == "" {
				fmt.Println("⚠️ Task title cannot be empty!")
				continue
			}

			newTask := Task{
				ID:          len(tasks) + 1,
				Title:       cleanTitle,
				IsCompleted: false,
				CreatedAt:   time.Now(),
			}

			tasks = append(tasks, newTask)
			saveTasks(tasks)
			fmt.Printf("✅ Task '%s' added!\n", cleanTitle)

		case "3":
			fmt.Print("Enter task ID to mark complete: ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("⚠️ Invalid ID number!")
				continue
			}

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].IsCompleted = true
					found = true
					break
				}
			}

			if found {
				saveTasks(tasks)
				fmt.Printf("🎉 Task #%d marked as complete!\n", id)
			} else {
				fmt.Printf("❌ Task #%d not found!\n", id)
			}

		case "4":
			fmt.Print("Enter task ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("⚠️ Invalid ID number!")
				continue
			}

			updatedTasks := []Task{}
			found := false

			for _, task := range tasks {
				if task.ID == id {
					found = true
					continue // Skip adding this task to delete it
				}
				updatedTasks = append(updatedTasks, task)
			}

			if found {
				tasks = updatedTasks
				saveTasks(tasks)
				fmt.Printf("🗑️ Task #%d deleted!\n", id)
			} else {
				fmt.Printf("❌ Task #%d not found!\n", id)
			}

		case "5":
			fmt.Println("\n👋 Exiting. Happy Coding!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please choose between 1 and 5.")
		}
	}
}
