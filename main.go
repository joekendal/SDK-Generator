package main

import (
	"fmt"
	"log"
)

func main() {
	c := getConfig()
	if &c.aws != (&awsConfig{}) {
		fmt.Println("Grabbing Swagger from AWS")
		err := getSwagger(c)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

}
