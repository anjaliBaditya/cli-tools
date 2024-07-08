package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "arith",
		Short: "Arithmetic CLI tool",
		Long:  `A CLI tool to perform basic arithmetic operations.`,
	}

	var addCmd = &cobra.Command{
		Use:   "add [num1] [num2]",
		Short: "Add two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			num1, err1 := strconv.ParseFloat(args[0], 64)
			num2, err2 := strconv.ParseFloat(args[1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Please provide valid numbers")
				return
			}
			fmt.Printf("Result: %f\n", num1+num2)
		},
	}

	var subCmd = &cobra.Command{
		Use:   "sub [num1] [num2]",
		Short: "Subtract two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			num1, err1 := strconv.ParseFloat(args[0], 64)
			num2, err2 := strconv.ParseFloat(args[1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Please provide valid numbers")
				return
			}
			fmt.Printf("Result: %f\n", num1-num2)
		},
	}

	var mulCmd = &cobra.Command{
		Use:   "mul [num1] [num2]",
		Short: "Multiply two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			num1, err1 := strconv.ParseFloat(args[0], 64)
			num2, err2 := strconv.ParseFloat(args[1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Please provide valid numbers")
				return
			}
			fmt.Printf("Result: %f\n", num1*num2)
		},
	}

	var divCmd = &cobra.Command{
		Use:   "div [num1] [num2]",
		Short: "Divide two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			num1, err1 := strconv.ParseFloat(args[0], 64)
			num2, err2 := strconv.ParseFloat(args[1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Please provide valid numbers")
				return
			}
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
				return
			}
			fmt.Printf("Result: %f\n", num1/num2)
		},
	}

	rootCmd.AddCommand(addCmd, subCmd, mulCmd, divCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
