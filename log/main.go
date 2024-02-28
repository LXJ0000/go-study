package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("log 1")
	fmt.Println(1)
	//log.Fatalln("log 2")
	//fmt.Println(2)
	log.Panicln("log 3")
	fmt.Println(3)
	fmt.Println(4)

}
