package main

import "fmt"

func main () {
   // var ages [3] int = [3]int{20, 25, 30} //cannot add more than a fixed length
   var ages = [3]int{20, 25, 30}

    names := [4]string{"yoshi", "mario", "peach", "bowser"}
    names[1] = "luigi"

   fmt.Println(ages, len(ages))
   fmt.Println(names, len(names))

    // slices (use arrays under to hood)
    var scores = []int{100, 50, 60}
    scores[2] = 25
    scores = append(scores, 85) //append function returns a new slice, so we must reassign

    fmt.Println(scores, len(scores))

    // slice ranges
    rangeOne   := names[1:3]//prints indices 1-2
    rangeTwo   := names[2:] //prints indices 2-3
    rangeThree := names[:3] //prints indices 1-2

    fmt.Println(rangeOne, rangeTwo, rangeThree)

    rangeOne = append(rangeOne, "koopa")
    fmt.Println(rangeOne)
}