package main

import (
	"fmt"

	"news.com/events/src"
)

func main1(){
	s := another.Sphere{
		Radius : 8.0,
	}
	another.CalculateVolume(s,"Sphere")
	fmt.Println(s)

	c := another.Cube{
		Length : 8.0,
	}
	another.CalculateVolume(c,"Cube")
	fmt.Println(c)

	b := another.Box{
		Length : 6.0,
		Width : 7.0,
		Height : 8.0,
	}
	another.CalculateVolume(b,"Box")
	fmt.Println(b)
}