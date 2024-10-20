package main

import "fmt"
import "unicode/utf8"

func dataTypes(){
    var num int = 4392
    fmt.Println(num)
    // int - depends on platform 32bit/ 64bit
    // int8 - 8bits/ 1byte -128 to 127
    // int16 - range 32,000
    // int32
    // int64 64 bits/8byte

    // uint - unsigned integer (only stores non-negative values)
    // uint8  0 to 255
    // uint16 0 to 65,000
    // uint32 and uint64
    var x uint = 500
    fmt.Printf("Type: %T, value: %v \n", x, x)
    // Printf to format, here %T tell the type and %v value & \n nextline

    var empty int
    fmt.Println(empty) // default value of all int's, unit's(int8, int16,... uint64) and float is 0

    var floatNum float32 = 546537.98000 // this will return 546538
    // float32 small decimals, don't return the precise value
    var floatNum2 float64 = 734678.8200909
    // float64 large numbers, takes more space, return the exact value

    fmt.Println("Float32: ", floatNum)
    fmt.Println("Float64: ", floatNum2)

    var myString string = "Hello world" // single line string
    // default value of string is an empty string = ''
    var myMultiLineString string = `Hello
    now I am in second line
    !`

    fmt.Println(myString)
    fmt.Println(myMultiLineString)

    fmt.Println(len(myString)) // len() function return the no. of Bytes instead of the length. some symbols may give length = 2

    fmt.Println(utf8.RuneCountInString(myString)) // count no. of character, length of the string

    // rune data type represent a character
    var myRune rune = 'a'
    fmt.Println(myRune)

    var myBoolean bool = false // true or false
    fmt.Println(myBoolean)
    // default value of boolean in false

}
