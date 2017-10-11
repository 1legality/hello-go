package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("1.1 + 1.2 = ", 1.1+1.2)

	var a string = "I am a string"
	fmt.Println(a)

	fmt.Println(plus(1, 2))

	for i := 1; i <= 10; i++ {
		if i == 1 {
			fmt.Println("Starting loop")
		}

		fmt.Println(i);
	}

}