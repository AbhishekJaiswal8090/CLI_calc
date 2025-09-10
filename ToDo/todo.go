// todo.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Task holds one todo item
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const tasksFile = "tasks.json"

func loadTasks() ([]Task, error) {
	// if file doesn't exist, return empty slice
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return []Task{}, nil
	}
	f, err := os.ReadFile(tasksFile)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(f, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func main() {
	// ensure tasks file is in current working directory
	dir, _ := os.Getwd()
	fmt.Printf("Using tasks file: %s\n\n", filepath.Join(dir, tasksFile))

	if len(os.Args) < 2 {
		fmt.Println("usage: todo <command> [options]")
		fmt.Println("commands: add, list, do")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "task title (or pass as positional arg)")
		addCmd.Parse(os.Args[2:])

		// allow title as positional arg if not set via flag
		if *title == "" {
			if addCmd.NArg() == 0 {
				fmt.Println("add: provide -title or pass title as argument")
				return
			}
			*title = addCmd.Arg(0)
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("error loading tasks:", err)
			return
		}
		t := Task{ID: nextID(tasks), Title: *title}
		tasks = append(tasks, t)
		if err := saveTasks(tasks); err != nil {
			fmt.Println("error saving tasks:", err)
			return
		}
		fmt.Printf("added: #%d %s\n", t.ID, t.Title)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		showAll := listCmd.Bool("all", false, "show completed tasks as well")
		listCmd.Parse(os.Args[2:])

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("error loading tasks:", err)
			return
		}
		for _, t := range tasks {
			if t.Done && !*showAll {
				continue
			}
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] #%d %s\n", status, t.ID, t.Title)
		}

	case "do":
		doCmd := flag.NewFlagSet("do", flag.ExitOnError)
		id := doCmd.Int("id", -1, "task id to mark done (or pass id as positional arg)")
		doCmd.Parse(os.Args[2:])
		if *id == -1 {
			if doCmd.NArg() == 0 {
				fmt.Println("do: provide -id or pass id as positional arg")
				return
			}
			// parse positional arg to int
			var n int
			_, err := fmt.Sscanf(doCmd.Arg(0), "%d", &n)
			if err != nil {
				fmt.Println("invalid id:", doCmd.Arg(0))
				return
			}
			*id = n
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("error loading tasks:", err)
			return
		}
		found := false
		for i := range tasks {
			if tasks[i].ID == *id {
				tasks[i].Done = true
				found = true
				break
			}
		}
		if !found {
			fmt.Println("task id not found:", *id)
			return
		}
		if err := saveTasks(tasks); err != nil {
			fmt.Println("error saving tasks:", err)
			return
		}
		fmt.Printf("marked done: #%d\n", *id)

	default:
		fmt.Println("unknown command:", cmd)
		fmt.Println("commands: add, list, do")
	}
}
