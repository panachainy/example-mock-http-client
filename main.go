package main

import (
	"fmt"

	"example-mock-http-client/client"
)

func main() {
	exampleClient := client.NewExampleClient()
	resp, err := exampleClient.GetName("1")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	fmt.Println("Resp ", resp)
}
