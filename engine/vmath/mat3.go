package vmath

type Mat3 struct {
	M00, M01, M02 float32
	M10, M11, M12 float32
	M20, M21, M22 float32
}

func Mat3Create(
	m00 float32, m01 float32, m02 float32, 
	m10 float32, m11 float32, m12 float32, 
	m20 float32, m21 float32, m22 float32) Mat3 {
	
	var temp Mat3
	temp.M00 = m00
	temp.M01 = m01
	temp.M02 = m02

	temp.M10 = m10
	temp.M11 = m11
	temp.M12 = m12

	temp.M20 = m20
	temp.M21 = m21
	temp.M22 = m22

	return temp
}

func Mat3Identity() Mat3 {
	return Mat3Create(
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0)
}

func Mat3Add(a *Mat3, b *Mat3) Mat3 {
	var temp Mat3
	temp.M00 = a.M00 + b.M00
	temp.M01 = a.M01 + b.M01
	temp.M02 = a.M02 + b.M02

	temp.M10 = a.M10 + b.M10
	temp.M11 = a.M11 + b.M11
	temp.M12 = a.M12 + b.M12

	temp.M20 = a.M20 + b.M20
	temp.M21 = a.M21 + b.M21
	temp.M22 = a.M22 + b.M22

	return temp
}

func Mat3Sub(a *Mat3, b *Mat3) Mat3 {
	var temp Mat3
	temp.M00 = a.M00 - b.M00
	temp.M01 = a.M01 - b.M01
	temp.M02 = a.M02 - b.M02

	temp.M10 = a.M10 - b.M10
	temp.M11 = a.M11 - b.M11
	temp.M12 = a.M12 - b.M12

	temp.M20 = a.M20 - b.M20
	temp.M21 = a.M21 - b.M21
	temp.M22 = a.M22 - b.M22

	return temp
}

func Mat3MulScalar(a *Mat3, s float32) Mat3 {
	var temp Mat3
	temp.M00 = a.M00 * s
	temp.M01 = a.M01 * s
	temp.M02 = a.M02 * s

	temp.M10 = a.M10 * s
	temp.M11 = a.M11 * s
	temp.M12 = a.M12 * s

	temp.M20 = a.M20 * s
	temp.M21 = a.M21 * s
	temp.M22 = a.M22 * s

	return temp	
}

func Mat3DivScalar(a *Mat3, s float32) Mat3 {
	var temp Mat3
	temp.M00 = a.M00 / s
	temp.M01 = a.M01 / s
	temp.M02 = a.M02 / s

	temp.M10 = a.M10 / s
	temp.M11 = a.M11 / s
	temp.M12 = a.M12 / s

	temp.M20 = a.M20 / s
	temp.M21 = a.M21 / s
	temp.M22 = a.M22 / s

	return temp	
}

func Mat3MulMatrix(a *Mat3, b *Mat3) Mat3 {
	var temp Mat3

	temp.M00 = a.M00 * b.M00 + a.M01 * b.M10 + a.M02 * b.M20
	temp.M01 = a.M00 * b.M01 + a.M01 * b.M11 + a.M02 * b.M21
	temp.M02 = a.M00 * b.M02 + a.M01 * b.M12 + a.M02 * b.M22

	temp.M10 = a.M10 * b.M00 + a.M11 * b.M10 + a.M12 * b.M20
	temp.M11 = a.M10 * b.M01 + a.M11 * b.M11 + a.M12 * b.M21
	temp.M12 = a.M10 * b.M02 + a.M11 * b.M12 + a.M12 * b.M22

	temp.M20 = a.M20 * b.M00 + a.M21 * b.M10 + a.M22 * b.M20
	temp.M21 = a.M20 * b.M01 + a.M21 * b.M11 + a.M22 * b.M21
	temp.M22 = a.M20 * b.M02 + a.M21 * b.M12 + a.M22 * b.M22

	return temp
}

func Mat3Transpose(a *Mat3) Mat3 {
	var temp Mat3
	temp.M00 = a.M00
	temp.M01 = a.M10
	temp.M02 = a.M20

	temp.M10 = a.M01
	temp.M11 = a.M11
	temp.M12 = a.M21

	temp.M20 = a.M02
	temp.M21 = a.M21
	temp.M22 = a.M22

	return temp
}

func Mat3Trace(a *Mat3) float32 {
	return a.M00 + a.M11 + a.M22
}

func Mat3Det(a *Mat3) float32 {
	var ma, mb, mc Mat2

	ma = Mat2Create(
		a.M11, a.M12,
		a.M21, a.M22)

	mb = Mat2Create(
		a.M10, a.M12,
		a.M20, a.M22)

	mc = Mat2Create(
		a.M10, a.M11,
		a.M20, a.M21)

	return a.M00 * Mat2Det(&ma) - a.M01 * Mat2Det(&mb) + a.M02 * Mat2Det(&mc)
}

func Mat3Inverse(m *Mat3) Mat3 {
	var a, b, c Mat2
	var d, e, f Mat2
	var g, h, i Mat2

	// a
	a = Mat2Create(
		m.M11, m.M12,
		m.M21, m.M22)

	// b
	b = Mat2Create(
		m.M10, m.M12,
		m.M20, m.M22)

	// c
	c = Mat2Create(
		m.M10, m.M11,
		m.M20, m.M21)

	// d
	d = Mat2Create(
		m.M01, m.M02,
		m.M21, m.M22)

	// e
	e = Mat2Create(
		m.M00, m.M02,
		m.M20, m.M22)

	// f
	f = Mat2Create(
		m.M00, m.M01,
		m.M20, m.M21)

	// g
	g = Mat2Create(
		m.M01, m.M02,
		m.M11, m.M12)

	// h
	h = Mat2Create(
		m.M00, m.M02,
		m.M10, m.M12)

	// i
	i = Mat2Create(
		m.M00, m.M01,
		m.M10, m.M11)

	
	var temp Mat3

	temp = Mat3Create(
		Mat2Det(&a), -Mat2Det(&b), Mat2Det(&c),
		-Mat2Det(&d), Mat2Det(&e), -Mat2Det(&f),
		Mat2Det(&g), -Mat2Det(&h), Mat2Det(&i))
		
	var temp2 = Mat3Transpose(&temp)

	var det = Mat3Det(m)

	return Mat3DivScalar(&temp2, det)
}

func Mat3MulVec3(m *Mat3, v *Vec3) Vec3 {
	var temp Vec3

	temp.X = m.M00 * v.X + m.M01 * v.Y + m.M02 * v.Z
	temp.Y = m.M10 * v.X + m.M11 * v.Y + m.M12 * v.Z
	temp.Z = m.M20 * v.X + m.M21 * v.Y + m.M22 * v.Z

	return temp
}

func Mat3ToArray(m *Mat3) []float32 {
	var temp []float32
	temp = append(temp, m.M00)
	temp = append(temp, m.M01)
	temp = append(temp, m.M02)

	temp = append(temp, m.M10)
	temp = append(temp, m.M11)
	temp = append(temp, m.M12)

	temp = append(temp, m.M20)
	temp = append(temp, m.M21)
	temp = append(temp, m.M22)
	return temp
}