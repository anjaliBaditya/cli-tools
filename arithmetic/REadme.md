# Arithmetic CLI Tool
## Overview
The Arithmetic CLI Tool is a command-line utility built in Go (Golang) using the Cobra library. It allows users to perform basic arithmetic operations such as addition, subtraction, multiplication, and division from the command line.

# Features
- Addition (add): Adds two numbers together.
- Subtraction (sub): Subtracts one number from another.
- Multiplication (mul): Multiplies two numbers.
- Division (div): Divides one number by another.


# Requirements
Go (Golang) installed on your machine
Dependencies managed using Go modules
Installation
To install the Arithmetic CLI Tool, clone the repository and build the executable:
```bash
git clone <repository_url>
cd arithmetic-cli
go build -o arith
```

# Usage
After building the executable, you can use the CLI tool to perform arithmetic operations from your terminal. Here are some examples:

```bash 
# Addition
./arith add 5 3   # Output: Result: 8.000000

# Subtraction
./arith sub 5 3   # Output: Result: 2.000000

# Multiplication
./arith mul 5 3   # Output: Result: 15.000000

# Division
./arith div 5 3   # Output: Result: 1.666667

```

# Command Syntax
```bash 
arith command [num1] [num2]
```
Replace command with one of the following:

- add: Adds two numbers
- sub: Subtracts two numbers
- mul: Multiplies two numbers
- div: Divides two numbers
## Replace [num1] and [num2] with the numbers you want to operate on.
