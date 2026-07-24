package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

func main() {
	fmt.Println("=================================")
	fmt.Println("🚀 Welcome to Must Task CLI")
	fmt.Println("=================================")

	tasks := []Task{
		{ID: 1, Title: "Set up Go environment", CreatedAt: time.Now()},
		{ID: 2, Title: "Learn Go structs and functions", CreatedAt: time.Now()},
	}

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
			fmt.Println("\n📋 Current Tasks:")
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
			fmt.Printf("✅ Task '%s' added successfully!\n", cleanTitle)

		case "3":
			fmt.Println("\n👋 Exiting. Happy Coding!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please choose 1, 2, or 3.")
		}
	}
}
