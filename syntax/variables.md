## Variables

- string => "Hello World"
- int (int8, int16, int32, int64) => 24
- float (float32, float64) => 3.14
- rune => 'A'
- bool => true
-

## Variables declaration

2 types declaration

- long declaration
- short declaration

Long Declaration

```go
func main() {
  var string name = "Jonh"
  var int age = 24
}
```

Short Declaration

```go
func main() {
  city := "Tashkent"
  score := 96.8
}
```

Multiple variable declaration

```go
var (
  rune A = 'A'
  name string = "John"
  number int = 24
)
```

Const Declaration

```go
const PI float64 = 3.14
```

Zero Values

```go
var x int // Default: 0
var y string // Default: "" empty string
var z bool // Default: false
```
