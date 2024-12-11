# Haze Programming Language

Haze is an interpreted programming language implemented in Go with a Pratt parser. This project was developed by following Thorsten Ball's book "Writing an Interpreter in Go".

## Features

- Dynamic typing
- First-class and higher-order functions
- Built-in functions
- Integer, boolean, string data types
- Arrays and hash maps
- Arithmetic operations and comparisons
- Modulo operation support
- Closures
- REPL environment

## Prerequisites

- Go 1.16 or later

## Installation

```bash
# Clone the repository
git clone https://github.com/tharun0120/haze.git

# Navigate to project directory
cd haze

# Build the interpreter
go build -o haze
```

## Usage

### REPL Mode

Start the interactive REPL:

```bash
./haze
```

You'll see:

```
Hey [username], Welcome to the Haze Interpreter!
This is the Haze REPL, Type away!
To exit, type in "exit"!
>>
```

### Execute .haze Files

Run a Haze source file:

```bash
./haze path/to/file.haze
```

## Language Examples

### Hello World

```haze
print("Hello World");
```

### Variables and Arithmetic

```haze
let x = 10;
let y = 5;
let sum = x + y;
let product = x * y;
let quotient = x / y;
let remainder = x % y;

print(sum);      // 15
print(product);  // 50
print(quotient); // 2
print(remainder);// 0
```

### Functions

```haze
let add = fn(a, b) {
    return a + b;
};

let result = add(5, 3);
print(result); // 8
```

### Arrays

```haze
let arr = [1, 2, 3];
print(arr[0]); // 1
print(len(arr)); // 3
print(first(arr)); // 1
print(last(arr)); // 3
```

### Hash Maps

```haze
let hash = {
    "name": "Haze",
    "type": "Language",
    "cool": true
};

print(hash["name"]); // "Haze"
```

### Control Flow

```haze
let max = fn(x, y) {
    if (x > y) {
        return x;
    } else {
        return y;
    }
};

print(max(5, 10)); // 10
```

### Built-in Functions

- `len()`: Get length of arrays and strings
- `first()`: Get first element of array
- `last()`: Get last element of array
- `rest()`: Get all elements except first
- `push()`: Add element to array
- `print()`: Print to console

## Project Structure

```
ast/          - Abstract Syntax Tree definitions
evaluator/    - Evaluation logic
examples/     - Example Haze programs
lexer/        - Lexical analysis
object/       - Object system implementation
parser/       - Pratt parser implementation
repl/         - REPL environment
token/        - Token definitions
main.go       - Entry point
```

## Acknowledgments

- Thorsten Ball for his book "Writing an Interpreter in Go"
