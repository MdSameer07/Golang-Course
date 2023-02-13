package main

import (
	"news.com/events/src")

func main(){
	s := another.Sphere{
		Radius : 8.0,
	}
	another.CalculateVolume(s,"Sphere")
	another.StringerFunction(s)

	c := another.Cube{
		Length : 8.0,
	}
	another.CalculateVolume(c,"Cube")
	another.StringerFunction(c)

	b := another.Box{
		Length : 6.0,
		Width : 7.0,
		Height : 8.0,
	}
	another.CalculateVolume(b,"Box")
	another.StringerFunction(b)
}