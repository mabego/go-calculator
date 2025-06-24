# go-calculator

## Installation

```shell
go install github.com/mabego/go-calculator@latest
```

## Usage

### REPL

```
$ go-calculator
calculator> (2.5 - 1.35) * 2.0
2.3
calculator> -sin((-1+2.5)*pi)
1
calculator> 180*atan2(log(e), log10(10))/pi
45
calculator> exit
```

### Command mode

```
$ go-calculator "(2.5 - 1.35) * 2.0"
2.3
$ go-calculator "-sin((-1+2.5)*pi)"
1
$ go-calculator "180*atan2(log(e), log10(10))/pi"
45
```

### Library

```go
package main

import (
	"fmt"
	"log"

	"github.com/mabego/go-calculator"
)

func main() {
	calc := calculator.New()
	
	val, err := calc.Calculate("(2.5 - 1.35) * 2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 2.3

	val, err = calc.Calculate("-sin((-1+2.5)*pi)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 1

	val, err = calc.Calculate("180*atan2(log(e), log10(10))/pi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 45
}
```

## Acknowledgements

This project is a fork of [mnogu/go-calculator](https://github.com/mnogu/go-calculator).