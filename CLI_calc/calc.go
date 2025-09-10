package main

import (
	"fmt"
	"flag"
)

func main() {

	a := flag.Int64("a",0,"first operand")
	b :=flag.Int64("b",0,"second operand")
	op:=flag.String("op","add","op add|mul|div|sub")
	verbose:=flag.Bool("V",false,"verbose")
    
	flag.Parse()
	var res int64

	switch *op{
	case "add", "+":
		res = *a + *b
	case "sub","-":
		res = *a - *b
	case "mul","*":
		res = *a * *b
	case "div","/":
		if *b == 0 {
			fmt.Println("Error: DivisionByZero")
			return
		}
		res = *a / *b
	default:
		fmt.Println("Unknown input : You input wrong operands")
		return
	}	

	if *verbose {
		fmt.Printf("a=%d b=%d op=%s  res=%d\n", *a, *b, *op, res)
	}else{
		fmt.Println(res)
	}

}