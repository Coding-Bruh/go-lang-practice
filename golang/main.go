package main

import "fmt"

var someName = "hello"

func main () {

    // string 
    var nameOne string = "mario"
    var nameTwo = "luigi"
    var nameThree string

    fmt.Println(nameOne, nameTwo, nameThree)

    nameOne = "peach"
    nameThree = "bowser"

    fmt.Println(nameOne, nameTwo, nameThree)

    nameFour := "yoshi"
    fmt.Println(nameFour)

    // ints
    var ageOne int = 29
    var ageTwo = 30
    ageThree := 40

    fmt.Println(ageOne, ageTwo, ageThree)

    // bits & memory
    var numOne int8 = 25
    var numTwo int8 = -128
    var numThree uint8 = 255 //255 max valid number for uint8
    var numFour uint16 = 256

    fmt.Println(numOne, numTwo, numThree, numFour)

    // floating point
    var scoreOne float32 = 25.98
    var scoreTwo float64 = 9899867587385364566574.7
    
    fmt.Println(scoreOne, scoreTwo)
}