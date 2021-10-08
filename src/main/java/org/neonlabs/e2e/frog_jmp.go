package a03frogjmp

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func SolutionFrogJmp(X int, Y int, D int) int {
	// write your code in Go 1.4
    if X == Y {
        return 0
    }
    var candidate = (Y - X) / D
    if X + candidate * D >= Y {
        return candidate
    } else {
        return candidate + 1
    }
}
