package main

import (
	"fmt"
)


func main(){

	fmt.Println("This line prints stuff, using Println")
	for i := 0; i < 10; i++ {
		fmt.Println("Hallo Go!");
	}

	somethingUsefull := [5]float64{ 98, 93, 77, 82, 83 }
	for i :0; i < len(somethingUsefull); i++ {
		fmt.Println(somethingUsefull[i])
	}

}
