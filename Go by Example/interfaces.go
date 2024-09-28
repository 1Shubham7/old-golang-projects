package main
import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

type geometry interface{
	area() int
	parameter() int
}

func (a rectangle) area() int{
	return a.length * a.width
}

func (a rectangle) parameter() int{
	return ((a.length + a.width)*2)
}

func main() {
	myrect := rectangle{
		length: 10,
		width: 10,
	}
	fmt.Println(myrect.area())
	fmt.Println(myrect.parameter())
}