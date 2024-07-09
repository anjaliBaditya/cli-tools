package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define flags
	rootDir := flag.String("root-dir", "", "Root directory to operate on")
	list := flag.Bool("list", false, "List files and directories")
	create := flag.String("create", "", "Create a new file or directory")
	copyFrom := flag.String("copy-from", "", "Copy file from")
	copyTo := flag.String("copy-to", "", "Copy file to")
	moveFrom := flag.String("move-from", "", "Move file from")
	moveTo := flag.String("move-to", "", "Move file to")
	delete := flag.String("delete", "", "Delete file or directory")
	search := flag.String("search", "", "Search for files by name or content")
	flag.Parse()

	// Check required flags
	if *rootDir == "" {
		fmt.Println("Error: required flag -root-dir not provided")
		flag.Usage()
		os.Exit(1)
	}

	// Perform operations
	if *list {
		listFilesAndDirs(*rootDir)
	} else if *create!= "" {
		createFileOrDir(*rootDir, *create)
	} else if *copyFrom!= "" && *copyTo!= "" {
		copyFile(*copyFrom, *copyTo)
	} else if *moveFrom!= "" && *moveTo!= "" {
		moveFile(*moveFrom, *moveTo)
	} else if *delete!= "" {
		deleteFileOrDir(*rootDir, *delete)
	} else if *search!= "" {
		searchFiles(*rootDir, *search)
	}
}

func listFilesAndDirs(rootDir string) {
	// List files and directories in the root directory
	files, err := os.ReadDir(rootDir)
	if err!= nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createFileOrDir(rootDir string, name string) {
	// Create a new file or directory
	path := filepath.Join(rootDir, name)
	if err := os.MkdirAll(path, os.ModePerm); err!= nil {
		log.Fatal(err)
	}
}

func copyFile(from string, to string) {
	// Copy a file from one location to another
	fromFile, err := os.Open(from)
	if err!= nil {
		log.Fatal(err)
	}
	defer fromFile.Close()
	toFile, err := os.Create(to)
	if err!= nil {
		log.Fatal(err)
	}
	defer toFile.Close()
	_, err = io.Copy(toFile, fromFile)
	if err!= nil {
		log.Fatal(err)
	}
}

func moveFile(from string, to string) {
	// Move a file from one location to another
	err := os.Rename(from, to)
	if err!= nil {
		log.Fatal(err)
	}
}

func deleteFileOrDir(rootDir string, name string) {
	// Delete a file or directory
	path := filepath.Join(rootDir, name)
	err := os.RemoveAll(path)
	if err!= nil {
		log.Fatal(err)
	}
}

func searchFiles(rootDir string, query string) {
	// Search for files by name or content
	files, err := os.ReadDir(rootDir)
	if err!= nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), query) {
			fmt.Println(file.Name())
		}
	}
}
