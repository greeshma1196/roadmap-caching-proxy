package main

import (
	"flag"
	"fmt"
)

func main() {

	portPtr := flag.Int("port", 8080, "an int")
	originPtr := flag.String("origin", "http://dummyjson.com", "a string")
	flag.Parse()

	fmt.Printf("port: %d\n", *portPtr)
	fmt.Printf("origin: %s\n", *originPtr)
}
