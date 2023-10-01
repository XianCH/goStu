package main

type Shape interface {
	Area() float64
}

type Cricle struct {
	Radius float64
}

func (c *Cricle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Ractangle struct {
	Width float64
	Hight float64
	
}

func (r *Ractangle) Area() float64 {
	return r.Hight * r.Width
}
.
func main() {

}
