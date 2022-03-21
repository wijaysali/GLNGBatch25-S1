package main

import "fmt"

func main(){
	// const full_name string = "Golang"
	// fmt.Printf("Hello %s", full_name)

	// var value = (5 + 2)*3
	// fmt.Println(value)

	var firstCondition bool= 3 < 5
	var secondCondition bool = "golang" == "Golang"
	var thirdCondition bool = 10 != 3.8
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition", firstCondition)
	fmt.Println("second condition", secondCondition)
	fmt.Println("third condition", thirdCondition)
	fmt.Println("fourth condition", fourthCondition)

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("left && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("left || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!left \t(%t\n", wrongReverse)




}