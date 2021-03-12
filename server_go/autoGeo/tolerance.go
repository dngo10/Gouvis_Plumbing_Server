package autoGeo

type Tolerance struct {
	EqualPoint  float64
	EqualVector float64
}

var Global = Tolerance{EqualPoint: 10 ^ -10, EqualVector: 10 ^ -12}
