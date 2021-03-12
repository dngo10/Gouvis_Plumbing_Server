package main

import (
	"autocad1/autoGeo"
	"fmt"
)

func main() {
	fmt.Println("Hello WOrld")

	//
	var t autoGeo.Tolerance = autoGeo.Tolerance{EqualPoint: 1, EqualVector: 3}
	fmt.Println(t.EqualPoint)
	fmt.Println(t.EqualVector)

	v := autoGeo.NewVector2d(23, 23)

	fmt.Println(v.X)
}
