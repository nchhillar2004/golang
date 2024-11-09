package main
import "fmt"

func main(){
    fmt.Println("Hello world from Go!")

    var myVar int = 128 // declare a variable, better practise to use var and data_type
    myVar2 := 200 // can use : instead of var, GOATed way
    // := (walrus operator) declares a new variable and assigns a value

    fmt.Println(myVar)
    fmt.Println(myVar2) // can't leave unused variables so printing is necessary

    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Welcome ", name)

    fmt.Println("-- Data Types --")
    dataTypes()

    fmt.Println("-- Functions --")
    functions()

    fmt.Println("-- Control flow --")
    controlFlow()

    fmt.Println("-- Formatting string --")
    formattingString()
}
