package main

import 	(
		"fmt"
		"net/http"
		"log"
		"io/ioutil"
)

func main (){
	fmt.Println("this is friggin gosome")
	resp, err := http.Get("http://duxilio.nl");
	
 	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}
	
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("\n%v\n\n", string(body))
}
