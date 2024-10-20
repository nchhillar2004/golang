package main

import "fmt"

func ageValidator(){
    fmt.Print("Enter your age: ")
    var age int8 // better to use uint8 because age can't be negative

    fmt.Scanln(&age)

    if age>=18{
        fmt.Println("You can vote")
    } else if age<18 && age>0{
        fmt.Println("Your are under age")
    } else{
        fmt.Println("Enter a valid age")
    }
}
