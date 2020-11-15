---
marp: true
---

# Go

![bg left height:600px](golang-gopher.svg)

---

# Code Structure

WTF

---

# WTF

Works = works well
Testable = unit testing, test double
Fit = Readable, Maintainable, Simple

---

How's your organization

![width:800px](black-and-white-blackboard-business-chalkboard-356043.jpg)

---

How about your meal

![width:800px](flat-lay-photography-of-three-tray-of-foods-1640775.jpg)

---

# Testable pattern

---

## First Class Replacing

action.go

```go
var Action = func(a,b int) int {
    return a+b
}

func handler(w http.ResponseWriter, r *http.Request) {
    Action(1,2)
    ...
}
```

action_test.go

```go
func TestAction(t *testing.T) {
    Action = func(a,b int) int {
        return 0
    }
}
```

---

## First Class Field

action.go

```go
type actionFunc func(int,int) int

type handler struct {
    action actionFunc
}
```

action_test.go

```go
func TestAction(t *testing.T) {
    action = func(a,b int) int {
        return 0
    }

    h := handler{
        action: action,
    }
}
```

---

## HOF and Closure

action.go

```go
type actionFunc func(a,b int) int

func handler(action actionFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

    }
}
```

action_test.go

```go
func TestAction(t *testing.T) {
    h := handler(func(a,b int) int {
        return 0
    })
}
```

---

## Interface

action.go

```go
type actioner interface {
    Action(int,int) int
}

func handler(feature actioner) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        feature.Action(1, 2)
    }
}
```

---

## Interface and function implementation

action.go

```go
type actionFunc func(int,int) int
    
func (fn actionFunc) Action(a, b int) int {
    return fn(a,b)
}

func handler(feature actioner) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        feature.Action(1, 2)
    }
}
```

---

## Interceptor Middleware

```go
{
    data, err := base64.StdEncoding.DecodeString(string(b))
    if err != nil {
    }

    buff := bytes.NewBuffer(data)
    r.Body = ioutil.NopCloser(buff)
}

type EncodeWriter struct {
    http.ResponseWriter
}

func (w *EncodeWriter) Write(b []byte) (int, error) {
    str := base64.StdEncoding.EncodeToString(b)
    return w.ResponseWriter.Write([]byte(str))
}

```

---

## Context

![bg left width:400px](totem.jpg)

---

## S.O.L.I.D. in go

By Robert C. Martin

> Sigle Responsibility (SRP)
> Open/closed principle (OCP)
> Liskov substitution (LSP)
> Interface segregation (ISP)
> Dependency inversion (DIP)

---

## Sigle Responsibility (SRP)

"In any well-designed system, objects should only have a single responsibility."

---

## Open/closed principle (OCP)
"A software module should be open for extension but closed for modification."
go using the Composition

example: io.ReadCloser

```go
type MyReader struct {
    io.ReadCloser
}

func (r *MyReader) Read(p []byte) (n int, err error) {
    log.Println(string(b))
    return w.ResponseWriter.Write(b)
}
```

---

## Liskov substitution (LSP)

If, for each object, O1 of type S there is an object O2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when O1 is substituted for O2, then S is a subtype of T.

---

## Interface segregation (ISP)

"The bigger the interface, the weaker the abstraction."

---

## Dependency inversion (DIP)

"High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions."
