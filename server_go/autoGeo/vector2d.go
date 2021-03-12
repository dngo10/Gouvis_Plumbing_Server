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

func (v_origin *Vector2d) Add(v *Vector2d) *Vector2d {
	var x = v_origin.X + v.X
	var y = v_origin.Y + v.Y

	return NewVector2d(x, y)
}

func (v_origin *Vector2d) DivideBy(value float64) *Vector2d {
	var x = v_origin.X / value
	var y = v_origin.Y / value
	return NewVector2d(x, y)
}

func (v_origin *Vector2d) DotProduct(v *Vector2d) float64 {
	return v_origin.X*v.X + v_origin.Y*v.Y
}

func (v_origin *Vector2d) IsZeroLength() bool {
	return v_origin.Length < Global.EqualVector
}

func (v_origin *Vector2d) GetAngleTo(v *Vector2d) float64 {
	if !v_origin.IsZeroLength() && !v.IsZeroLength() {
		return math.Acos(v_origin.DotProduct(v) / (v_origin.Length * v.Length))
	} else {
		return math.NaN()
	}
}

func (v_origin *Vector2d) GetUnitVector() *Vector2d {
	if v_origin.IsZeroLength() {
		return nil
	}

	var x = v_origin.X / v_origin.Length
	var y = v_origin.Y / v_origin.Length

	return NewVector2d(x, y)
}

func (v_origin *Vector2d) GetUnitVectorT(tolerlance *Tolerance) *Vector2d {
	return nil
}
