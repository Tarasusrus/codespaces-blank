package main

import (
	"fmt"
	"log"
	"test/go/src/github.com/headfirstgo/homework"
)

func main() {
	coordinates := homework.Coordinates{}
	err:= coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(1122.42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latityde())
	fmt.Println(coordinates.Longitude())
}