# go

## elementary class day 3

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

# First-Class Function

```go
var add = func(a, b int) int {
    return a + b
}

fmt.Println(add(1, 2))
```

---

# Higher-Order Function

```go
func hof(fn func(string) string) {
    ...
}

func hof() func(string) string {
    ...
}
```

---

# Higher-Order Function

https://dev.to/pallat/hof-in-go-18mm

---

# Closure Function

```go
func main() {
    fn1, fn2 := factory()
    fn1()
    fn1()
    fmt.Println(fn2())

    fn1()
    fmt.Println(fn2())
}

func factory() (func(), func() int) {
    var i int
    return func() {
            i++
        },
        func() int {
            return i
        }
}
```

---

# Exercise

> Make a Wrapper for FizzBuzz handler to validate JWT

---

# func type

```go
type areaFunc func(float64, float64) float64 
```

---

# Exercise

Validate JWT in FizzBuzzHandler

```http
GET http://localhost:8080/fizzbuzz/3

Header["Authorization"]: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
```

---

# method on function

```go
    type IntnFunc func(int) int

    func (fn IntnFunc) Intn(n int) int {
        return fn(n)
    }
```

---

# goroutine

```go
func main() {
    total := 10
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i)
    }
    fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
    fmt.Println(i)
}

```

---

# goroutine

```go
var wg = sync.WaitGroup{}

func main() {
    total := 10
    wg.Add(total)
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i)
    }
    wg.Wait()
    fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
    fmt.Println(i)
    wg.Done()
}
```

---

# channel 

keyword `chan`

- no buffered channel
- buffered channel

---

# buffered channel

```go
total := 10
ch := make(chan int, total)
for i := total; i > 0; i-- {
    ch <- i
}
close(ch)

for i := range ch {
    fmt.Println(i)
}
```

---

# no buffered channel

```go
func main() {
    total := 10
    ch := make(chan struct{})
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i, ch)
    }
    for i := 0; i < total; i++ {
        <-ch
    }
    fmt.Println(time.Now().Sub(now))
}

func printout(i int, ch chan struct{}) {
    fmt.Println(i)
    ch <- struct{}{}
}
```

---

# Exercise

> generate 10 random 16 bits and use 10 go-routine to encrypt it

---

# Closure Function

```go
func main() {
    fn := factory()
    fmt.Println(fn())
    fmt.Println(fn())
    fmt.Println(fn())
}

func factory() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

---

