# Notes CLI
## Overview
The notes CLI application is a simple tool for managing text notes. It allows you to add, list, and delete notes. The notes are stored as text files in a directory named notes.

# Features
- Add Note: Create a new note with a specified title and content.
- List Notes: Display a list of all note titles.
- Delete Note: Remove a note by its title.

## Installation
- Clone the repository:
```bash 
git clone https://github.com/anjaliBaditya/cli-tools
cd cli-tools/note-maker
```

- Build the application:
```bash 
go build -o notes
```

- Ensure the notes directory exists in your working directory:
```bash
mkdir -p notes
```

## Usage
The notes CLI provides three main commands: add, list, and delete.

## Add Note
To add a new note, use the add command followed by the note's title and content:
```bash 
./notes add "Note Title" "This is the content of the note."
```

### Example:
```bash 
./notes add "Shopping List" "Milk, Bread, Eggs"
```

## List Notes
To list all notes, use the list command:
```bash
./notes list
```

## Delete Note
To delete a note by its title, use the delete command followed by the note's title:
```bash 
./notes delete "Note Title"
```
### Example:
```bash
./notes delete "Shopping List"
```

# Code Structure
- main.go: The main file containing the CLI commands and their implementations.
- notes: The directory where notes are stored as text files.
# Dependencies
Cobra: A library for creating powerful modern CLI applications.
