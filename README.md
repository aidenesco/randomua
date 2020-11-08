# RandomUA [![GoDoc](https://godoc.org/github.com/aidenesco/randomua?status.svg)](https://godoc.org/github.com/aidenesco/randomua) [![Go Report Card](https://goreportcard.com/badge/github.com/aidenesco/randomua)](https://goreportcard.com/report/github.com/aidenesco/randomua)
This package generates a random user agent from a preloaded, weighted set.

Thank you to intoli.com for releasing their dataset.

## Installation
```sh
go get -u github.com/aidenesco/randomua
```

## Usage

```go
import "github.com/aidenesco/randomua"

func main() {
    ua := randomua.Random()

    fmt.Println(ua) // Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36
}
```

