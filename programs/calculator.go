package main

import "fmt"

func calculator(){
    var num1, num2 int
    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)

    var sum int = num1+num2
    var sub int = num1-num2
    var mult int = num1*num2
    var div int
    if num2 == 0 {
        fmt.Println("cannot divide by zero")
    } else{
        div = num1/num2
        fmt.Printf("Sum = %v\n Sub = %v\n Mult = %v\n Div = %v\n", sum, sub, mult, div)
    }
}
