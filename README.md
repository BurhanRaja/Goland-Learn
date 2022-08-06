# Golang

### Go is somewhere between C++ and Javascript when it comes to `Fast and fun for humans` and `Fast and efficient for computers`. 

## Compiler vs Interpreter

### Go is a Compiled Language

- Compiler is used to convert the human readable code to machine code to run and Interpreter is used to run code in run-time.
- The languages that uses compiler catches error in compile time only and the languages that uses iterpreter catches error in runtime. 
- Compiler is faster than Iterpreter.

### Go is a strong typed language i.e. you can not change the string "Hello World!!" to int or int to string once declared.

### Memory Management

There are two methods for allocating memory :-

#### new() :- 

- Allocates Memory but no initialization
- You will get memory address
- Zeroed Storage, will not able to store anyhting

#### make() :-

- Allocates Memory and initialization takes place
- You will get memory address
- Non-zeroed Storage, will able to store anyhting 