package vmath

import "math"

type Vec3 struct {
	X, Y, Z float32
}

func Vec3Create(x float32, y float32, z float32) Vec3 {
	var temp Vec3
	temp.X = x
	temp.Y = y
	temp.Z = z
	return temp
}

func Vec3Add(a *Vec3, b *Vec3) Vec3 {
	var temp Vec3
	temp.X = a.X + b.X
	temp.Y = a.Y + b.Y
	temp.Z = a.Z + b.Z
	return temp
}

func Vec3Sub(a *Vec3, b *Vec3) Vec3 {
	var temp Vec3
	temp.X = a.X - b.X
	temp.Y = a.Y - b.Y
	temp.Z = a.Z - b.Z
	return temp
}

func Vec3Mul(a *Vec3, b float32) Vec3 {
	var temp Vec3
	temp.X = a.X * b
	temp.Y = a.Y * b
	temp.Z = a.Z * b
	return temp
}

func Vec3Div(a *Vec3, b float32) Vec3 {
	var temp Vec3
	temp.X = a.X / b
	temp.Y = a.Y / b
	temp.Z = a.Z / b
	return temp	
}

func Vec3Length(a *Vec3) float32 {
	return float32(math.Sqrt(float64(a.X * a.X + a.Y * a.Y + a.Z * a.Z)))
}

func Vec3Normalize(a *Vec3) Vec3 {
	return Vec3Div(a, Vec3Length(a))
}

func Vec3Dot(a *Vec3, b *Vec3) float32 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func Vec3Angle(a *Vec3, b *Vec3) float32 {
	var d float32 = Vec3Dot(a, b)
	var la float32 = Vec3Length(a)
	var lb float32 = Vec3Length(b)
	return float32(math.Acos(float64(d / (la * lb))))
}

func Vec3Cross(a *Vec3, b *Vec3) Vec3 {
	var temp Vec3

	temp.X = a.Y * b.Z - a.Z * b.Y
	temp.Y = a.Z * b.X - a.X * b.Z
	temp.Z = a.X * b.Y - a.Y * b.X
	
	return temp
}

func Vec3ToArray(a *Vec3) []float32 {
	var temp []float32
	temp = append(temp, a.X)
	temp = append(temp, a.Y)
	temp = append(temp, a.Z)
	return temp
}