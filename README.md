# A Prisoner Problem

### An interesting logic problem, the solution to which I coded using Go.

In my second semester at Georgia Tech, I'm taking CS 2051, a discrete math course. Our professor regularly gives us puzzles to work through, and this logic problem was one I found particularly challenging. As I'm currently learning [Go](https://golang.org), I decided to write a simple program to test whether my solution would work for larger test cases. 


### The problem.

Suppose there are **x** prisoners, each of whom are simultaneously given a hat with a number on it. The numbers are assigned at random, so long as the number is ∈ [1, **x**]. The numbers can repeat amongst the prisoners.

Simultaenously, all prisoners will be asked to guess the number that is on their own hat. If every prisoner guesses incorrectly, all the prisoners are executed (grim, I know). If at least one of the prisoners guesses their number correctly, then all of the prisoners live.

Prior to being given the hats with their numbers, the prisoners are allowed to communicate a strategy as to how to guess their own number. Once the prisoners are assigned a number, there can no longer be any communication between prisoners. What is a strategy that the prisoners can use to ensure that they will all survive?


### The solution, kinda.

I won't disclose the full solution here. You can easily download the [program](https://github.com/SohanChoudhury/APrisonerProblem/blob/master/APrisonerProblem.go) to see the specifics. Below, I'll discuss the framework of my code.

As you may notice, my ```main()``` function, which is run from the console when the program is executed, is quite simple. I did this intentionally, as to showcase the benefits of Go as a function-based language.

```go
func main() {

    solve(create())

}
```

The semantics are quite straightforwards here. First, I'm using the ```create()``` function to initialize a slice that contains the desired **x** number of prisoners. I then give each of the prisoners a random number within the appropriate range.

```go
func create() []int {

    var i int  // declares i as int
    fmt.Print("For how many prisoners would you like to test the strategy? ")
    fmt.Scan(&i) // takes user input from console
    
    .
    .
    .
    
    // after initlaizing the array and giving the prisoners random values,
    // I print the slice to the console and return it
    
    fmt.Println(r)
    return r
}
```

Next, this returned slice is used as a parameter when executing ```solve()```, the entirety of which is shown below.

```go
func solve(r []int) {

    c := false  // c checks whether or not the strategy works

    for i, t := range r{
        x := guess(r, i)  // guess returns an integer value, x, which is prisoner i's guess 
        if x == t {
            fmt.Printf("The prisoners live. Prisoner %d guessed their number, %d, correctly.\n", i, x)
            c = true  // the strategy works!
        }
    }

    if c == false { fmt.Println("The prisoners die. Your strategy does not work.")}

}
```

The ```guess()``` function contains the true strategy for the problem. It takes in as parameters the slice of all prisoners and the integer specifying the index of the prisoner whose guess we must determine. Finally, the guess is returned as an integer.

```go
func guess(r []int, i int) int {

    .
    .
    .
  
    return x
}
```

This solution works whenever the number of prisoners is ≥ 4, so long as this number is within 32 or 64 bits, depending on the system. Download the program and try it out!
