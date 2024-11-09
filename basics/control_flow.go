package main

import (
	"errors"
	"fmt"
) // import multiple packages

func controlFlow(){
    var divisionOutput, remainder, err = intDivision(482, 5)
    
    // conditional statements
    if err!=nil {
        fmt.Println(err.Error())
    } else if remainder == 0{
        fmt.Println("Result of division: ", divisionOutput)
    } else{
        fmt.Printf("Result of division: %v, Remainder: %v\n", divisionOutput, remainder)
    }

    // switch case 
    switch remainder{
    case 0:
        fmt.Println("The division was exact")
    default:
        fmt.Println("We have a remainder")
    }
}

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error 
    // error by default is nil
    if denominator==0 {
        err = errors.New("Cannot divide by zero")
        return 0, 0, err
    }
    var result = numerator/denominator
    var remainder = numerator%denominator
    return result, remainder, err
}
