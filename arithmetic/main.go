package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := createRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "arith",
		Short: "Arithmetic CLI tool",
		Long:  `A CLI tool to perform basic arithmetic operations.`,
	}

	rootCmd.AddCommand(
		createAddCmd(),
		createSubCmd(),
		createMulCmd(),
		createDivCmd(),
	)

	return rootCmd
}

func createAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [num1] [num2]",
		Short: "Add two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run:   runAdd,
	}
}

func createSubCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sub [num1] [num2]",
		Short: "Subtract two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run:   runSub,
	}
}

func createMulCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mul [num1] [num2]",
		Short: "Multiply two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run:   runMul,
	}
}

func createDivCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "div [num1] [num2]",
		Short: "Divide two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run:   runDiv,
	}
}

func runAdd(cmd *cobra.Command, args []string) {
	num1, num2, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %f\n", num1+num2)
}

func runSub(cmd *cobra.Command, args []string) {
	num1, num2, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %f\n", num1-num2)
}

func runMul(cmd *cobra.Command, args []string) {
	num1, num2, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %f\n", num1*num2)
}

func runDiv(cmd *cobra.Command, args []string) {
	num1, num2, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if num2 == 0 {
		fmt.Println("Error: Division by zero")
		return
	}
	fmt.Printf("Result: %f\n", num1/num2)
}

func parseArgs(args []string) (float64, float64, error) {
	num1, err1 := strconv.ParseFloat(args[0], 64)
	num2, err2 := strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("Error: Please provide valid numbers")
	}
	return num1, num2, nil
}
