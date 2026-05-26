package main

import "fmt"

func fizzbuzz() {
	for i := 1; i<=100; i++{
		fizzbuz := ""
		if i % 3 == 0{
			fizzbuz += "fizz"
		}
		if i % 5 == 0 {
			fizzbuz += "buzz"
		}
		if fizzbuz != ""{
			fmt.Println(fizzbuz)
		}else{
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
