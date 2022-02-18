package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
	foo()
	fmt.Println("Everything else.")

	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in function foo")
}
