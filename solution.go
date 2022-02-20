package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type SideNum uint

const SidesTriangle SideNum = 3
const SidesSquare SideNum = 4
const SidesCircle SideNum = 0

func CalcSquare(sideLen float64, sidesNum SideNum) float64 {
	cs := 0.0
	switch sidesNum {
	case SidesSquare:
		cs = math.Pow(sideLen, 2)
	case SidesTriangle:
		cs = math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case SidesCircle:
		cs = math.Pi * math.Pow(sideLen, 2)
	}

	return cs
}
