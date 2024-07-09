package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const notesDir = "notes"

func main() {
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.Mkdir(notesDir, 0755)
	}

	var rootCmd = &cobra.Command{Use: "notes"}

	var addCmd = &cobra.Command{
		Use:   "add [title] [content]",
		Short: "Add a new note",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			content := strings.Join(args[1:], " ")
			filename := filepath.Join(notesDir, title+".txt")
			err := ioutil.WriteFile(filename, []byte(content), 0644)
			if err != nil {
				fmt.Printf("Error adding note: %v\n", err)
				return
			}
			fmt.Printf("Added note: %s\n", title)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all notes",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(notesDir)
			if err != nil {
				fmt.Printf("Error listing notes: %v\n", err)
				return
			}
			fmt.Println("Notes:")
			for _, file := range files {
				fmt.Println(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
			}
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [title]",
		Short: "Delete a note by title",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			filename := filepath.Join(notesDir, title+".txt")
			err := os.Remove(filename)
			if err != nil {
				fmt.Printf("Error deleting note: %v\n", err)
				return
			}
			fmt.Printf("Deleted note: %s\n", title)
		},
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.Execute()
}
