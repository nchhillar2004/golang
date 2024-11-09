package main

import "fmt"

func formattingString(){
    var name = "Saul goodman"
    var rating = 9.423

    // put the data (string, int, float) in a string and format it
    msg := fmt.Sprintf("Hi %s, your rating is %.1f out of 10\n", name, rating) // Sprintf returns the formatted string

    // now we have a string 'msg' with name and other data like rating i.e type float

    fmt.Print(msg)
}
