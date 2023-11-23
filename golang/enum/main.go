package main

import (
	"fmt"
	"log"

	"github.com/candy12t/sandbox/golang/enum/video"
)

func main() {
	{
		v, err := video.New("Netflix")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)
		fmt.Printf("%T\n", v)
	}
	{
		v, err := video.New("ABC")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)
		fmt.Printf("%T\n", v)
	}
}
