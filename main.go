package main

import (
	"fmt"
	"strconv"
	"./myLib"
)

func plus(a int, b int) int {
	return a+b
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("1.1 + 1.2 = ", 1.1+1.2)

	var a string = "I am a string"
	fmt.Println("var a = " + a)

	fmt.Println("function calling test : " + strconv.Itoa(plus(1, 2))) // also convert int to string

	lines := myLib.StringToLines(`
		line1
		line2
		line3
	`)
	fmt.Println("String2Lines : ", lines)

	lines = myLib.File2Lines("./myLib/Files2Lines.txt")
	fmt.Println("File2Lines : ", lines)

	for i := 1; i <= 10; i++ {
		if i == 1 {
			fmt.Println("Starting loop")
		}

		fmt.Println(i);
	}
}