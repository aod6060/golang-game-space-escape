package vmath

import "math"

type Vec2 struct {
	X, Y float32
}


// I'll just treat it like C
func Vec2Create(x float32, y float32) Vec2 {
	var temp Vec2
	temp.X = x
	temp.Y = y
	return temp
}

func Vec2Add(a *Vec2, b *Vec2) Vec2 {
	var temp Vec2
	temp.X = a.X + b.X
	temp.Y = a.Y + b.Y
	return temp
}

func Vec2Sub(a *Vec2, b *Vec2) Vec2 {
	var temp Vec2
	temp.X = a.X - b.X
	temp.Y = a.Y - b.Y
	return temp
}

func Vec2Mul(a *Vec2, b float32) Vec2 {
	var temp Vec2
	temp.X = a.X * b
	temp.Y = a.Y * b
	return temp
}

func Vec2Div(a *Vec2, b float32) Vec2 {
	var temp Vec2
	temp.X = a.X / b
	temp.Y = a.Y / b
	return temp
}

func Vec2Length(a *Vec2) float32 {
	return float32(math.Sqrt(float64(a.X * a.X + a.Y * a.Y)))
}

func Vec2Normalize(a *Vec2) Vec2 {
	return Vec2Div(a, Vec2Length(a))
}

func Vec2Dot(a *Vec2, b *Vec2) float32 {
	return a.X * b.X + a.Y * b.Y
}

func Vec2Angle(a *Vec2, b *Vec2) float32 {
	var d float32 = Vec2Dot(a, b)
	var la float32 = Vec2Length(a)
	var lb float32 = Vec2Length(b)
	return float32(math.Acos(float64(d / (la * lb))))
}

func Vec2ToArray(a *Vec2) []float32 {
	var temp []float32
	temp = append(temp, a.X)
	temp = append(temp, a.Y)
	return temp
}