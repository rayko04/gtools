# gtools

A lightweight command-line utility written in Go that implements a collection of common filesystem operations. This project was built to learn Go through practical development while following a modular and maintainable project structure.

## Features

- Print the current working directory
- List directory contents
- Display file contents
- Create files
- Create directories
- Copy files
- Move or rename files
- Remove files and empty directories
- Built-in help command
- Version information

## Project Structure

```text
gtools/
├── commands/
│   ├── cat.go
│   ├── cp.go
│   ├── ls.go
│   ├── mkdir.go
│   ├── mv.go
│   ├── pwd.go
│   ├── rm.go
│   └── touch.go
├── registry/
│   ├── command.go
│   ├── help.go
│   ├── registry.go
│   └── version.go
├── main.go
├── go.mod
└── README.md
```

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd gtools
```

Build the project:

```bash
go build -o gtools
```

Or run it directly:

```bash
go run .
```

## Usage

```bash
gtools <command> [arguments]
```

## Available Commands

| Command | Description |
|---------|-------------|
| `pwd` | Print the current working directory |
| `ls [directory]` | List the contents of a directory |
| `cat <file>` | Display the contents of a file |
| `touch <file>` | Create a file if it does not exist |
| `mkdir <directory>` | Create a directory |
| `rm <path>` | Remove a file or an empty directory |
| `mv <source> <destination>` | Move or rename a file or directory |
| `cp <source> <destination>` | Copy a file |
| `help [command]` | Display general or command-specific help |
| `version` | Display the current version |

## Examples

Print the current directory:

```bash
gtools pwd
```

List the current directory:

```bash
gtools ls
```

List another directory:

```bash
gtools ls ~/Documents
```

Display a file:

```bash
gtools cat notes.txt
```

Create a file:

```bash
gtools touch notes.txt
```

Create a directory:

```bash
gtools mkdir projects
```

Copy a file:

```bash
gtools cp source.txt backup.txt
```

Rename a file:

```bash
gtools mv old.txt new.txt
```

Remove a file:

```bash
gtools rm old.txt
```

Show help:

```bash
gtools help
```

Show help for a specific command:

```bash
gtools help cp
```

Show version:

```bash
gtools version
```

## Technologies

- Go
- Go Standard Library

## Learning Objectives

This project focuses on learning practical Go development, including:

- Project organization with multiple packages
- Command dispatch using a registry
- File and directory operations
- Error handling
- Resource management with `defer`
- Reading and writing files
- Working with the Go standard library
- Building maintainable command-line applications

## License

This project is licensed under the MIT License.
