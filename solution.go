package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type MyType struct {
	sides int
}

func CalcSquare(sideLen float64, sidesNum MyType) float64 {
	square := 0.0
	switch sidesNum.sides {
	case 0:
		square = math.Pi * math.Pow(sideLen, 2)
		fmt.Printf("Square 0 %v",square)
	case 3:
		square = math.Pow(sideLen, 2) * math.Sqrt(3)/4
		fmt.Printf("Square 3%v",square)
	case 4:
		square = math.Pow(sideLen, 2)
		fmt.Printf("Square 4%v",square)
	default:
		square = 0
		fmt.Printf("Square default%v",square)
	}
	return square
}


