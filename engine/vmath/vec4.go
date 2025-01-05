package vmath

import "math"

type Vec4 struct {
	X, Y, Z, W float32
}

func Vec4Create(x float32, y float32, z float32, w float32) Vec4 {
	var temp Vec4
	temp.X = x
	temp.Y = y
	temp.Z = z
	temp.W = w
	return temp
}

func Vec4Add(a *Vec4, b *Vec4) Vec4 {
	var temp Vec4
	temp.X = a.X + b.X
	temp.Y = a.Y + b.Y
	temp.Z = a.Z + b.Z
	temp.W = a.W + b.W
	return temp
}

func Vec4Sub(a *Vec4, b *Vec4) Vec4 {
	var temp Vec4
	temp.X = a.X - b.X
	temp.Y = a.Y - b.Y
	temp.Z = a.Z - b.Z
	temp.W = a.W - b.W
	return temp
}

func Vec4Mul(a *Vec4, b float32) Vec4 {
	var temp Vec4
	temp.X = a.X * b
	temp.Y = a.Y * b
	temp.Z = a.Z * b
	temp.W = a.W * b
	return temp
}

func Vec4Div(a *Vec4, b float32) Vec4 {
	var temp Vec4
	temp.X = a.X / b
	temp.Y = a.Y / b
	temp.Z = a.Z / b
	temp.W = a.W / b
	return temp
}

func Vec4Length(a *Vec4) float32 {
	return float32(math.Sqrt(float64(a.X * a.X + a.Y * a.Y + a.Z * a.Z + a.W * a.W)))
}

func Vec4Normalize(a *Vec4) Vec4 {
	return Vec4Div(a, Vec4Length(a))
}

func Vec4Dot(a *Vec4, b *Vec4) float32 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z + a.W * b.W
}

func Vec4Angle(a *Vec4, b *Vec4) float32 {
	var d float32 = Vec4Dot(a, b)
	var la float32 = Vec4Length(a)
	var lb float32 = Vec4Length(b)
	return float32(math.Acos(float64(d / (la * lb))))
}

func Vec4ToArray(a *Vec4) []float32 {
	var temp []float32
	temp = append(temp, a.X)
	temp = append(temp, a.Y)
	temp = append(temp, a.Z)
	temp = append(temp, a.W)
	return temp
}