# Command-Line To-Do Application

A simple command-line To-Do application written in Go. Manage your tasks using easy commands and a local JSON file for storage.

## Features
- Add tasks with a title
- List pending or all tasks
- Mark tasks as done
- Stores tasks in `tasks.json`

## Usage

1. Add a new task:
   ```sh
   go run todo.go add -title="Buy milk"
   ```
   or
   ```sh
   go run todo.go add "Pay rent"
   ```
   Output:
   Using tasks file: C:\Users\Abhishek\OneDrive\Desktop\go_project\ToDo\tasks.json
   added: #1 Buy milk

2. List tasks:
   ```sh
   go run todo.go list
   ```
   Output:
   [ ] #1 Buy milk
   [ ] #2 Pay rent

3. Mark a task as done:
   ```sh
   go run todo.go do -id=1
   ```
   Output:
   marked done: #1

4. List all tasks (including completed):
   ```sh
   go run todo.go list -all
   ```
   Output:
   [x] #1 Buy milk
   [ ] #2 Pay rent

## Requirements
- Go 1.16 or higher


