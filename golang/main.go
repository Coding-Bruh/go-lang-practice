package main

import "fmt"

var someName = "hello"

func main () {
    age := 35
    name := "Abhinav Chennapareddy" 

    //Print
    fmt.Print("hello, ")
    fmt.Print("world! \n")
    fmt.Print("new line \n")

    fmt.Println("hello ninjas!")
    fmt.Println("Goodbye ninjas!")
    fmt.Println("my age is", age , "and my name is", name)

    //Printf(format string) &_ = format specifier
    fmt.Printf("my age is %d and my name is %s\n", age, name)

    fmt.Printf("age is of type %t \n", age)
    fmt.Printf("you scored %f points \n", 255.55)
    fmt.Printf("you scored %0.1f points \n", 255.55)

    // Sprintf (save formatted stirngs)
    var str = fmt.Sprintf("my age is %d and my name is %s\n", age, name)
    fmt.Println("the saved string is:", str)
}