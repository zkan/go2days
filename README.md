# Golang Presentations

## How to run interactive presentation

https://godoc.org/golang.org/x/tools/present

    go get golang.org/x/tools/present

    export GOPATH=$(go env GOPATH)
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:$GOBIN

## How to run presentation

> marp -s .

## Instruction

### Elementary (3 days)

Day1

- Why Go

- Installation

- Simple Syntax (var, func, if, for)

- TDD
  - FizzBuzz
  - Diamond Star
  - Diamond Number

- Package
  - Show how package works
  - Expose things
  - Exercise main, diamond, fizzbuzz package

- Var

- Type
  - new type Int

- API FizzBuzz

Day2

- Array, Slice
  - split double

- Map
  - oscar male who got more than one
  - JWT custom with map

- Struct
  - API login

- Method
  - Captcha
  - API Captcha

Day3

- Interface

- API connect to db

### Intermediate (2 days)

Day1

- Exercise
  - moretypes/18

- First-Class

- HOF

- Type Function

- Go Routine

- Channel

Day2

- WTF

- Context

- Testable pattern

- Versioning

### Advance (2 days)

Day1

- S.O.L.I.D.
- Reader,Writer

Day2

- Benchmark testing

- Race Condition

- Game Of Life

### Ultimate

- Profiling

```
Outline : Golang Tutorial for 2 Day 

	•	Introduction to Golang
	•	Project Structure 
	•	Installation and configurations
	•	go root 
	•	go path 
	•	go env
	•	Setup Editor (VS Code) - - > extension 
	•	Write  first Go program 
	•	go run 
	•	go build
	•	Package
	•	Introduction to Go module 
	•	Basic of Go language - 1
	•	Variable & Basic type 
	•	Default Zero Value
	•	Input Standard
	•	Workshop 1 – Triangle Area 
	•	 Basic of Go language – 2
	•	If statement: condition{ }
	•	Switch condition
	•	Workshop 2 – Grade 
	•	Basic of GO language  - 3
	•	For (loop)
	•	Array
	•	Slice [] T
	•	Map Map [T] T
	•	Workshop 3 – Multiplication Table 
	•	Basic of Go language – 4 
	•	Struct 
	•	Struct with JSON
	•	Function 
	•	Custom Type 
	•	Method
	•	Method on primitive type 
	•	Method on function type
	•	Interface


	•	Basic of GO language – 5
	•	Goroutine
	•	Channel
	•	Defer
	•	Pointer
	•	Unit Test
	•	Go test
	•	Test Coverage 
	•	Test Coverage Report HTML
	•	API (GIN web framework+Error Handling) 
	•	GET
	•	POST
	•	PUT
	•	Delete
	•	Error Handling
	•	Workshop 4 – Hello API
	•	MongoDB 
	•	Connect 
	•	Insert
	•	Update
	•	Delete
	•	Find
	•	Workshop 5 – MongoDB 
	•	Swagger (API Document)
	•	Setup Swagger Library in GO
	•	Swagger Init
	•	Workshop 6 - Swagger
```