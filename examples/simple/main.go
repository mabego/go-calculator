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
