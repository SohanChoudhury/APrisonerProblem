package main

import (
    "fmt"
    "math/rand"
)


func main() {

    solve(create())

}

func create() []int {

    var i int  // declares i as int
    fmt.Print("For how many prisoners would you like to test the strategy? ")
    fmt.Scan(&i)

    r := make([]int, i)
    numOfPrisoners := len(r)


    for i, _ := range r{
        r[i] = rand.Intn(numOfPrisoners+1)
    }

    fmt.Println(r)
    return r
}


func solve(r []int) { // not sure what output will exactly be

    c := false

    for i, t := range r{
        x := guess(r, i)
        if x == t {
            fmt.Printf("The prisoners live. Prisoner %d guessed their number, %d, correctly.\n", i, x)
            c = true
        }
    }

    if c == false { fmt.Println("The prisoners die. Your strategy does not work.")}

}

func guess(r []int, i int) int {
    sum := 0
    for _, v := range r{
        sum += v
    }

    h := sum - r[i]

    numOfPrisoners := len(r)
    x := i - h

    for x < 0 {
        x += numOfPrisoners
    }

    return x
}