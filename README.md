# CLI To-Do Application

Hello! This is a simple CLI To-Do application built with Go and the Cobra library. It allows you to manage a list of tasks directly from the command line.

Tasks can be:
- Added
- Marked as complete
- Deleted
- Listed for easy tracking

All tasks are stored in a JSON file to persist data between sessions.

## Installation

To get started, clone this repository and navigate into the project directory. You can then run the application with the command below.

```shell
go run main.go
```

## Usage
To see all available commands, use the --help flag:
```shell
go run main.go --help
```
Each command also supports --help to provide specific information. For example:
```shell
go run main.go add --help
```
### Available Commands

| Command               | Description                                                |
|-----------------------|------------------------------------------------------------|
| `add [description]`   | Adds a new task with a description.                        |
| `list`                | Lists all tasks.                                           |
| `complete [id]`       | Marks the task with the specified ID as complete.          |
| `delete [id]`         | Deletes the task with the specified ID.                    |

