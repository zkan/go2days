---
marp: true
---

# go

## elementary class day 1

Pallat Anchaleechamaikorn
golang developer
Technical Coach

yod.pallat@gmail.com
https://github.com/pallat
https://github.com/focusive
https://github.com/gophernment

https://tour.golang.org/
https://github.com/uber-go/guide

---

# Getting Started

## Official Golang Page

https://golang.org/ https://golang.org/
https://go.dev/ https://go.dev/

---

# Go Users

https://github.com/golang/go/wiki/GoUsers

---

# Installation

https://dev.to/pallat/install-go-4a1a https://dev.to/pallat/install-go-4a1a

---

# Initial a project

## linux/Macbook

    mkdir -p ~/projects/hello
    cd hello

## windows

    md projects
    cd projects
    md hello
    cd hello

---

# Open VS Code

    code .

---

# Initial go module

    go mod init hello

go.mod

```go
    module hello

    go 1.14
```

---

# Hello World

run

    go run main.go

---

# Basic syntax

Variable declaration

```go
var s string
var s string = "Hello World"
var s = "Hello World"

var i int
var i int = 9
var i = 9

// only in function
s := "Hello World"
i := 9
```

---

# Basic syntax

Function

```go
    func add(a int, b int) int {
        return a + b
    }

    func add(a, b int) int {
        return a + b
    }
```

---

# Basic syntax

if/else

```go
    if a != b {
        println("a not equal to b")
    } else if a < b {
        println("a less than b")
    } else {
        println("ok")
    }
```

---

# Basic syntax

for

```go
    for i := 0; i < 10; i++ {

    }

    for i <= 10 {

    }

    for {
        
    }
```

---

# Prime factor

print prime number in 1..100

---

# package

rules

    only one package in any directory except testing file can plus suffix _test in there

    exposed name begins with capital charactor

---

# unit testing in go

*filename_test.go*

```go
    import "testing"

    func TestACase(t *testing.T) {

    }

    func Test_a_case(t *testing.T) {

    }

    func TestCases(t *testing.T) {
        t.Run("case number 1", func(t *testing.T) {
            
        })
        t.Run("case number 2", func(t *testing.T) {

        })
    }
```

---

# TDD

Test Driven Development

![TDD](./images/classic_tdd.png)

---

# TDD with the classic problem

FizzBuzz

    given 1 then "1"
    given 2 then "2"
    given 3 then "Fizz"
    given 4 then "4"
    given 5 then "Buzz"

---

# TDD with the classic problem

FizzBuzz

    given 1 then "1"
    given 2 then "2"
    given 3 then "Fizz"
    given 4 then "4"
    given 5 then "Buzz"
    given 6 then "Fizz"
    given 7 then "7"
    given 8 then "8"
    given 9 then "Fizz"
    given 10 then "Buzz"

---

# TDD with the classic problem

FizzBuzz

    given 1 then "1"
    given 2 then "2"
    given 3 then "Fizz"
    given 4 then "4"
    given 5 then "Buzz"
    given 6 then "Fizz"
    given 7 then "7"
    given 8 then "8"
    given 9 then "Fizz"
    given 10 then "Buzz"
    given 12 then "Fizz"
    given 15 then "FizzBuzz"
    given 18 then "Fizz"
    given 20 then "Buzz"
    given 30 then "FizzBuzz"

---

# TDD with the classic problem

FizzBuzz

any number divisible by three is replaced by the word fizz and any number divisible by five by the word buzz. Numbers divisible by 15 become fizz buzz

    1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14,
    Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26,
    Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, ...

---

# Exercise

## RESTful API of FizzBuzz

    try net/http with gorilla/mux

---

# Variable

```go
    var c, python, java bool

    var i, j int = 1, 2

    var c, python, java = true, false, "no!"

    c, python, java := true, false, "no!"
```

---

# Type

basic type

    bool

    string

    int  int8  int16  int32  int64

    uint uint8 uint16 uint32 uint64 uintptr

    byte  // alias for uint8

    rune  // alias for int32
          // represents a Unicode code point

    float32 float64

    complex64 complex128

---

# Zero Value

- *0#  _for_ _numeric_ _types_
- *false# _for_ _the_ _boolean_ _type_
- *""# (the empty string) for strings
