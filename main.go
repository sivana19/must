package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Task structure
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

const fileName = "tasks.json"

// Load tasks from tasks.json file
func loadTasks() []Task {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		// If file doesn't exist yet, return an empty list
		return []Task{}
	}

	var tasks []Task
	// Turn JSON text into Go struct data
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Println("⚠️ Warning: Could not parse tasks.json")
		return []Task{}
	}

	return tasks
}

// Save tasks to tasks.json file
func saveTasks(tasks []Task) {
	// Turn Go struct data into clean JSON text
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("❌ Error saving tasks:", err)
		return
	}

	// Write data to the file
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("❌ Error writing file:", err)
	}
}

func main() {
	fmt.Println("=================================")
	fmt.Println("🚀 Welcome to Must Task CLI")
	fmt.Println("=================================")

	// 1. Load existing tasks from tasks.json when app starts
	tasks := loadTasks()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu Options:")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Exit")
		fmt.Print("Enter option (1-3): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n📋 Your Current Tasks:")
			if len(tasks) == 0 {
				fmt.Println("   No tasks found!")
			}
			for _, task := range tasks {
				fmt.Printf("   [%d] %s (Created: %s)\n",
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
				ID:        len(tasks) + 1,
				Title:     cleanTitle,
				CreatedAt: time.Now(),
			}

			tasks = append(tasks, newTask)

			// 2. Save updated list to file immediately
			saveTasks(tasks)

			fmt.Printf("✅ Task '%s' added and saved!\n", cleanTitle)

		case "3":
			fmt.Println("\n👋 Exiting. Happy Coding!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please choose 1, 2, or 3.")
		}
	}
}
