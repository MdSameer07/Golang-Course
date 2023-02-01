package another

import (
	"fmt"
	"math"
)

type Cube struct {
	Length float64
}

func (c Cube) String() string{
	return fmt.Sprintf("The length of Cube is %v",c.Length)
}

type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (b Box) String() string{
	return fmt.Sprintf("The length,width and height of Box are %v,%v and %v",b.Length,b.Width,b.Height)
}

type Sphere struct {
	Radius float64
}

func (s Sphere) String() string{
	return fmt.Sprintf("The radius of the Sphere is %v",s.Radius)
}

type OfStructure interface {
	Volume() float64
	String() string
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func CalculateVolume(kind OfStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f\n", called, kind.Volume())
	fmt.Printf(kind.String())
	fmt.Printf("\n")
}

func main() {

	c := Cube{
		Length: 7,
	}

	b := Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := Sphere{
		Radius: 7.14,
	}

	CalculateVolume(c, "Cube")
	CalculateVolume(b, "Box")
	CalculateVolume(s, "Sphere")
}