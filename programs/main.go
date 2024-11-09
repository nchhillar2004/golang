package main

import "fmt"

func main(){
    fmt.Println("-- Welcome to my go programs --")
    fmt.Println(`Select a program to run: 
    1. Age validator
    2. Calculator
    3. Web
    3. Exit`)

    var choice uint8

    fmt.Print("Enter your choice (1-3): ")
    fmt.Scanln(&choice)

    switch choice{
    case 1:
        ageValidator()
    case 2:
        calculator()
    case 3:
        web()
    default:
        fmt.Println("Exiting...")
    }
}
