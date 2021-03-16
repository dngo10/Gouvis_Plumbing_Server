package autoGeo

import "math"

type Vector2d struct {
	X      float64
	Y      float64
	Angle  float64
	Length float64
}

func NewVector2d(X float64, Y float64) *Vector2d {
	v := Vector2d{X: X, Y: Y}
	v.Length = math.Sqrt(math.Pow(X, 2) + math.Pow(Y, 2))
	v.Angle = math.Atan2(v.Y, v.X)
	return &v
}

func (Vself *Vector2d) Add(v *Vector2d) *Vector2d {
	var x = Vself.X + v.X
	var y = Vself.Y + v.Y

	return NewVector2d(x, y)
}

func (Vself *Vector2d) DivideBy(value float64) *Vector2d {
	var x = Vself.X / value
	var y = Vself.Y / value
	return NewVector2d(x, y)
}

func (Vself *Vector2d) DotProduct(v *Vector2d) float64 {
	return Vself.X*v.X + Vself.Y*v.Y
}

func (Vself *Vector2d) IsZeroLength() bool {
	return Vself.Length < Global.EqualVector
}

func (Vself *Vector2d) GetAngleTo(v *Vector2d) float64 {
	if !Vself.IsZeroLength() && !v.IsZeroLength() {
		return math.Acos(Vself.DotProduct(v) / (Vself.Length * v.Length))
	} else {
		return math.NaN()
	}
}

func (Vself *Vector2d) GetUnitVector() *Vector2d {
	if Vself.IsZeroLength() {
		return nil
	}
	return Vself._getUnitVector()
}

func (Vself *Vector2d) GetUnitVectorT(tolerlance *Tolerance) *Vector2d {
	if math.Abs(Vself.Length) < tolerlance.EqualVector {
		return nil
	}
	return Vself._getUnitVector()
}

func (Vself *Vector2d) GetPerpendicularVector() *Vector2d {
	return NewVector2d(-Vself.Y, Vself.X)
}

func (Vself *Vector2d) IsCodirectionalTo(v *Vector2d) *Vector2d {
	var v1, v2 = v._getTwoSameLengthVector(Vself)

	if v1 == nil || v2 == nil {
		return nil
	}

}

//PRIVATE FUNCTIONS, DO NOT USE IT OUTSIDE OF STRUCT.

func (Vself *Vector2d) _getUnitVector() *Vector2d {
	var x = Vself.X / Vself.Length
	var y = Vself.Y / Vself.Length

	return NewVector2d(x, y)
}

func (Vself *Vector2d) _getTwoSameLengthVector(v *Vector2d) (*Vector2d, *Vector2d) {
	var v1 *Vector2d = NewVector2d(Vself.X, Vself.Y)
	var v2 *Vector2d = NewVector2d(v.X, v.Y)

	if v1.Length > v2.Length {
		v2._scaleVector(v2.Length / v1.Length)
	} else {
		v1._scaleVector(v1.Length / v1.Length)
	}

	return v1, v2
}

func (Vself *Vector2d) _scaleVector(factor float64) {
	if math.Abs(factor) < Global.EqualPoint {

	}
	Vself = Vself.DivideBy(1.0 / factor)
}
