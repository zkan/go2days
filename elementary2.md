# go

## elementary class day 2

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

# Array

var name [n]T

```go
    var array [5]string
```

---

# Array (for)

var name [n]T

```go
    var array [5]string

    for i := range array {
        fmt.Println(array[i])
    }
```

---

# Array (for)

var name [n]T

```go
    var array [5]string

    for i, v := range array {
        fmt.Println(i, v)
    }
```

---

# Slice

var name []T

```go
    var slice []int

    primes := [...]int{2, 3, 5, 7, 11, 13}    
    slice = primes[1:4]

    for i := range slice {
        fmt.Println(slice[i])
    }
```

---

# Zero value of slice is nil

```go
    var s []int

    if s == nil {
        fmt.Println("it's nil")
    }

// make it first
    s := make([]int, 5)
```

---

# append slice

```go
    s := []int{1, 2, 3, 4, 5}
    s = append(s, 6, 7, 8)
```

---

# Exercise

https://tour.golang.org/moretypes/18 https://tour.golang.org/moretypes/18
https://go-tour-th.appspot.com/moretypes/18 https://go-tour-th.appspot.com/moretypes/18

---

# map[T]T

```go
    var m map[string]string

    if m == nil {
        fmt.Println("it's nil")
    }
```

---

# map[T]T

make it first

```go
    m := make(map[string]string)

    if m == nil {
        fmt.Println("it's nil")
    }

    m["a"] = "apple"
    m["b"] = "banana"
    m["c"] = "coconut"
    m["d"] = "durian"
    m["e"] = "elderberry"
    m["f"] = "fig"
    m["g"] = "guava"
```

---

# map[T]T

```go
    m := map[string]string{
        "a" : "apple",
        "b" : "banana",
        "c" : "coconut",
        "d" : "durian",
        "e" : "elderberry",
        "f" : "fig",
        "g" : "guava",
    }

    delete(m, "d")

    for k, v := range m {
        fmt.Println(k, v)
    }
```

---

# Exercise

open a file oscar_age_male.csv
print any actors name who got the oscar more than one time

    Marlon Brando
    Daniel Day-Lewis
    Sean Penn
    Tom Hanks
    Fredric March
    Spencer Tracy
    Gary Cooper
    Jack Nicholson
    Dustin Hoffman

---

# len()

```go
    s := []primes{2, 3, 5, 7, 11, 13}
    m := map[string]string{
        "a" : "apple",
        "b" : "banana",
        "c" : "coconut",
        "d" : "durian",
        "e" : "elderberry",
        "f" : "fig",
        "g" : "guava",
    }

    fmt.Println(len(s), len(m))
```

---

# Pointer

Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

    var p *int
The & operator generates a pointer to its operand.

    i := 42
    p = &i
The # operator denotes the pointer's underlying value.

    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

---

# Type

    type newType T

    type text string

    type aliasType = T

    type text = string

---


# method on primitive type

```go
    type text string

    func (t text) split(sep string) []string{
        return strings.Split(string(t), sep)
    }
```

---

# Exercise

New Int Type with Method

.String() to convert integer to string
.Set(n int) to set new value
.Int() int

---

# struct type

```go
    type rectangle struct {
        w,l float64
    }
```

---

# struct usage

```go
    rec := rectangle{w: 10, l: 12}

    area := rec.w # rec.l
```

---

# Exercise

pick a Web Framework to make the Login API

request:

```json
    {
        "email": "abc@mail.com",
        "password": "12345678"
    }
```

response:

```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
```

---

# method on struct

```go
    type rectangle struct {
        width float64
        length float64
    }

    func (r rectangle) area() float64 {
        return r.width * r.length
    }
```

---

# method with pointer receiver

```go
    type rectangle struct {
        width float64
        length float64
    }

    func (r *rectangle) area() float64 {
        return r.width * r.length
    }
```

---

# Exercise

Captcha: (pattern,leftOperand,operator,rightOperand)

operator: 1=+,2=-,3=*

[left/right]operand: [0,1,2,3,4,5,6,7,8,9]/["zero","one","two","three","four","five","six","seven","eight","nine"]

patter: 1=number+operation+number in word
ex: pattern=1,leftOperand=1,operation=1,rightOperand=1 is `1 + one`
patter: 2=number in word+operation+number
ex: pattern=2,leftOperand=1,operation=1,rightOperand=1 is `one + 1`

