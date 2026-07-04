package main

import (
	"fmt"
	"os"
	"strconv"

	"task-cli/internal/cli"
	"task-cli/internal/storage"
	"task-cli/internal/task"
)

func printUsage() {
	fmt.Println("Usage: task-cli <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add <description>              Add a new task")
	fmt.Println("  update <id> <description>      Update an existing task")
	fmt.Println("  delete <id>                    Delete a task")
	fmt.Println("  mark-in-progress <id>          Mark a task as in-progress")
	fmt.Println("  mark-done <id>                 Mark a task as done")
	fmt.Println("  list                           List all tasks")
	fmt.Println("  list done                      List done tasks")
	fmt.Println("  list todo                      List tasks to do")
	fmt.Println("  list in-progress               List tasks in progress")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	store := storage.NewStorage("tasks.json")
	handler := cli.NewHandler(store)

	command := os.Args[1]
	var err error

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Description is required.")
			printUsage()
			os.Exit(1)
		}
		err = handler.Add(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: ID and Description are required.")
			printUsage()
			os.Exit(1)
		}
		id, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			fmt.Println("Error: ID must be a valid number.")
			os.Exit(1)
		}
		err = handler.Update(id, os.Args[3])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required.")
			printUsage()
			os.Exit(1)
		}
		id, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			fmt.Println("Error: ID must be a valid number.")
			os.Exit(1)
		}
		err = handler.Delete(id)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required.")
			printUsage()
			os.Exit(1)
		}
		id, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			fmt.Println("Error: ID must be a valid number.")
			os.Exit(1)
		}
		err = handler.UpdateStatus(id, task.StatusInProgress)
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required.")
			printUsage()
			os.Exit(1)
		}
		id, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			fmt.Println("Error: ID must be a valid number.")
			os.Exit(1)
		}
		err = handler.UpdateStatus(id, task.StatusDone)
	case "list":
		statusFilter := ""
		if len(os.Args) >= 3 {
			statusFilter = os.Args[2]
			if statusFilter != task.StatusTodo && statusFilter != task.StatusInProgress && statusFilter != task.StatusDone {
				fmt.Printf("Error: Invalid status filter '%s'. Use 'todo', 'in-progress', or 'done'.\n", statusFilter)
				os.Exit(1)
			}
		}
		err = handler.List(statusFilter)
	default:
		fmt.Printf("Error: Unknown command '%s'.\n", command)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
