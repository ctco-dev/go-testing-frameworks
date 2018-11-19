package main

import "errors"

// import 	"log"

func main() {
}

func Sum(a, b uint) uint {
	return sum(a, b)
}

func sum(a, b uint) uint {
	return a + b
}

func Divide(a, b uint) (uint, error) {
	if b == 0 {
		return 0, errors.New("Division by zero.")
	}
	return a / b, nil
}

func DivideUnsafe(a, b uint) uint {
	return a / b
}
