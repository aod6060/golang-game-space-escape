package vmath



type Mat2 struct {
	M00, M01 float32
	M10, M11 float32
}

func Mat2Create(m00 float32, m01 float32, m10 float32, m11 float32) Mat2 {
	var temp Mat2
	temp.M00 = m00
	temp.M01 = m01
	temp.M10 = m10
	temp.M11 = m11
	return temp
}

func Mat2Identiy() Mat2 {
	return Mat2Create(
		1.0, 0.0,
		0.0, 1.0)
}

func Mat2Add(a *Mat2, b *Mat2) Mat2 {
	var temp Mat2
	temp.M00 = a.M00 + b.M00
	temp.M01 = a.M01 + b.M01
	temp.M10 = a.M10 + b.M10
	temp.M11 = a.M11 + b.M11
	return temp
}

func Mat2Sub(a *Mat2, b *Mat2) Mat2 {
	var temp Mat2
	temp.M00 = a.M00 - b.M00
	temp.M01 = a.M01 - b.M01
	temp.M10 = a.M10 - b.M10
	temp.M11 = a.M11 - b.M11
	return temp
}

func Mat2MulScalar(a *Mat2, s float32) Mat2 {
	var temp Mat2
	temp.M00 = a.M00 * s
	temp.M01 = a.M01 * s
	temp.M10 = a.M10 * s
	temp.M11 = a.M11 * s
	return temp
}

func Mat2DivScalar(a *Mat2, s float32) Mat2 {
	var temp Mat2
	temp.M00 = a.M00 / s
	temp.M01 = a.M01 / s
	temp.M10 = a.M10 / s
	temp.M11 = a.M11 / s
	return temp	
}

func Mat2MulMatrix(a *Mat2, b *Mat2) Mat2 {
	var temp Mat2
	temp.M00 = a.M00 * b.M00 + a.M01 * b.M10
	temp.M01 = a.M00 * b.M01 + a.M01 * b.M11
	temp.M10 = a.M10 * b.M00 + a.M11 * b.M10
	temp.M11 = a.M10 * b.M01 + a.M11 * b.M11
	return temp		
}

func Mat2Transpose(a *Mat2) Mat2 {
	var temp Mat2
	temp.M00 = a.M00
	temp.M01 = a.M10
	temp.M10 = a.M01
	temp.M11 = a.M11
	return temp
}

func Mat2Trace(a *Mat2) float32 {
	return a.M00 + a.M11
}

func Mat2Det(a *Mat2) float32 {
	return a.M00 * a.M11 - a.M01 * a.M10
}

func Mat2Inverse(a *Mat2) Mat2 {
	var temp Mat2

	temp.M00 = a.M11
	temp.M01 = -a.M01
	temp.M10 = -a.M10
	temp.M11 = a.M00

	return Mat2DivScalar(&temp, Mat2Det(&temp))
}