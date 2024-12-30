# `iris`
A simple package to convert hex colors to ANSI256.

## Usage
`go get github.com/aboxofsox/iris`

### Setting Colors
```go
package main

import (
        "fmt"
        "github.com/aboxofsox/iris"
)

func main() {
        text := iris.SetColor("hello world", "#010101", "#FF0000")
        fmt.Println(text)
} 
```

### Setting Foreground Text Color

```go
package main

import (
        "fmt"
        "github.com/aboxofsox/iris"
)

func main() {
        text := iris.SetFgColor("hello world", "#FF000")
        fmt.Println(text)
}

```

### Setting Background Text Color
```go
package main

import (
        "fmt"
        "github.com/aboxofsox/iris"
)

func main() {
        text := iris.SetBgColor("hello world" "#FF0000")
}
```

### Stripping ANSI

```go
package main

import (
        "fmt"
        "github.com/aboxofsox/iris"
)

func main() {
        text := iris.SetHex("hello world", "#FF0000")
        stripped := iris.Strip(text)
        fmt.Println(text, stripped)
}
```
