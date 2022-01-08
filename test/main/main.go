package main

import (
	"encoding/json"
	"fmt"
	// "test/secondary"
	// "test/third"
)

func main() {
	// fmt.Println("Test 2")
	// fmt.Println(secondary.Clock)
	// third.Thirdly()

	// helloWorld := "hello world"

	// pointerToHW := &helloWorld
	// fmt.Println(pointerToHW)


	// reciever functions/methods
	// create person instance

	type Person struct {
		Hello string `json:"hello"`
	}

	myJson := map[string]string{"hello": "iamjson"}
	data, _ := json.Marshal(myJson)
	fmt.Println(data)
	fmt.Print("&t", data, "\n")
	// print new instance
	// fmt.Println(myPerson)
	// // Call welcome method/reciever function to welcome the new person
	// fmt.Println(myPerson.Welcome())
	// // Update the persons first name
	// myPerson.UpdateName("Sally")
	// myPerson.UpdateAge(32)
	// fmt.Println(myPerson)


	

}
