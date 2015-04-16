package main


import ("fmt"; "math"; "math/rand"; "time")

func main(){
	rand.Seed( time.Now().UTC().UnixNano())

	fmt.Println(math.Pi)
	fmt.Println(randInt(0,400))
}


func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
