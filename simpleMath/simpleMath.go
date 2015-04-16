package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){
	rand.Seed( time.Now().UTC().UnixNano())

	fmt.Println(randInt(0,400))
	fmt.Println(plus(20,40))
	fmt.Println(minus(40,20))
	fmt.Println(math.Pi)
	fmt.Println(doSomerandomMath())
}


func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func plus(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func doSomerandomMath() int{
	return 1
}

