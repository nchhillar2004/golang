package main

import "fmt"

func functions(){
    printMe("Hello world !!")
    
    var output int = addition(12, 48)
    fmt.Println("Sum = ", output)
    
    var divisionOutput, remainderValue int = division(128, 4) // specify miltiple variables for multiple return
    fmt.Printf("Division output = %v, Remainder = %v\n", divisionOutput, remainderValue)
}

func printMe(myString string){
    fmt.Println("Function called.")
    fmt.Println(myString)
}

// function returning an int value
func addition(num1 int, num2 int) int {
    var sum = num1 + num2
    return sum
}

// returning multiple values
func division(num1 int, num2 int) (int, int) {
    var result = num1/num2
    var remainder = num1%num2
    return result, remainder
}
