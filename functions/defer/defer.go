package main

import "fmt"

func cleanUp(resource string) {
	fmt.Printf("The resource %s is cleared.\n",resource)
}

func worker() {
	defer cleanUp("A")
	defer cleanUp("B")

	fmt.Println("Worker running")
}

func main() {
	worker()
}