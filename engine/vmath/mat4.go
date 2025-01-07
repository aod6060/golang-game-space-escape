package vmath

type Mat4 struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

func Mat4Create(
	m00 float32, m01 float32, m02 float32, m03 float32,
	m10 float32, m11 float32, m12 float32, m13 float32,
	m20 float32, m21 float32, m22 float32, m23 float32,
	m30 float32, m31 float32, m32 float32, m33 float32) Mat4 {
	
	var temp Mat4
	temp.M00 = m00
	temp.M01 = m01
	temp.M02 = m02
	temp.M03 = m03

	temp.M10 = m10
	temp.M11 = m11
	temp.M12 = m12
	temp.M13 = m13

	temp.M20 = m20
	temp.M21 = m21
	temp.M22 = m22
	temp.M23 = m23

	temp.M30 = m30
	temp.M31 = m31
	temp.M32 = m32
	temp.M33 = m33

	return temp
}

func Mat4Identity() Mat4 {
	return Mat4Create(
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0)
}

func Mat4Add(a *Mat4, b *Mat4) Mat4 {
	var temp Mat4

	temp.M00 = a.M00 + b.M00
	temp.M01 = a.M01 + b.M01
	temp.M02 = a.M02 + b.M02
	temp.M03 = a.M03 + b.M03

	temp.M10 = a.M10 + b.M10
	temp.M11 = a.M11 + b.M11
	temp.M12 = a.M12 + b.M12
	temp.M13 = a.M13 + b.M13

	temp.M20 = a.M20 + b.M20
	temp.M21 = a.M21 + b.M21
	temp.M22 = a.M22 + b.M22
	temp.M23 = a.M23 + b.M23

	temp.M30 = a.M30 + b.M30
	temp.M31 = a.M31 + b.M31
	temp.M32 = a.M32 + b.M32
	temp.M33 = a.M33 + b.M33

	return temp
}

func Mat4Sub(a *Mat4, b *Mat4) Mat4 {
	var temp Mat4

	temp.M00 = a.M00 - b.M00
	temp.M01 = a.M01 - b.M01
	temp.M02 = a.M02 - b.M02
	temp.M03 = a.M03 - b.M03

	temp.M10 = a.M10 - b.M10
	temp.M11 = a.M11 - b.M11
	temp.M12 = a.M12 - b.M12
	temp.M13 = a.M13 - b.M13

	temp.M20 = a.M20 - b.M20
	temp.M21 = a.M21 - b.M21
	temp.M22 = a.M22 - b.M22
	temp.M23 = a.M23 - b.M23

	temp.M30 = a.M30 - b.M30
	temp.M31 = a.M31 - b.M31
	temp.M32 = a.M32 - b.M32
	temp.M33 = a.M33 - b.M33

	return temp
}

func Mat4MulScalar(a *Mat4, s float32) Mat4 {
	var temp Mat4

	temp.M00 = a.M00 * s
	temp.M01 = a.M01 * s
	temp.M02 = a.M02 * s
	temp.M03 = a.M03 * s

	temp.M10 = a.M10 * s
	temp.M11 = a.M11 * s
	temp.M12 = a.M12 * s
	temp.M13 = a.M13 * s

	temp.M20 = a.M20 * s
	temp.M21 = a.M21 * s
	temp.M22 = a.M22 * s
	temp.M23 = a.M23 * s

	temp.M30 = a.M30 * s
	temp.M31 = a.M31 * s
	temp.M32 = a.M32 * s
	temp.M33 = a.M33 * s

	return temp
}

func Mat4DivScalar(a *Mat4, s float32) Mat4 {
	var temp Mat4

	temp.M00 = a.M00 / s
	temp.M01 = a.M01 / s
	temp.M02 = a.M02 / s
	temp.M03 = a.M03 / s

	temp.M10 = a.M10 / s
	temp.M11 = a.M11 / s
	temp.M12 = a.M12 / s
	temp.M13 = a.M13 / s

	temp.M20 = a.M20 / s
	temp.M21 = a.M21 / s
	temp.M22 = a.M22 / s
	temp.M23 = a.M23 / s

	temp.M30 = a.M30 / s
	temp.M31 = a.M31 / s
	temp.M32 = a.M32 / s
	temp.M33 = a.M33 / s

	return temp
}

func Mat4MulMatrix(a *Mat4, b *Mat4) Mat4 {
	var temp Mat4

	// Row 1
	temp.M00 = a.M00 * b.M00 + a.M01 * b.M10 + a.M02 * b.M20 + a.M03 * b.M30
	temp.M01 = a.M00 * b.M01 + a.M01 * b.M11 + a.M02 * b.M21 + a.M03 * b.M31
	temp.M02 = a.M00 * b.M02 + a.M01 * b.M12 + a.M02 * b.M22 + a.M03 * b.M32
	temp.M03 = a.M00 * b.M03 + a.M01 * b.M13 + a.M02 * b.M23 + a.M03 * b.M33

	// Row 2
	temp.M10 = a.M10 * b.M00 + a.M11 * b.M10 + a.M12 * b.M20 + a.M13 * b.M30
	temp.M11 = a.M10 * b.M01 + a.M11 * b.M11 + a.M12 * b.M21 + a.M13 * b.M31
	temp.M12 = a.M10 * b.M02 + a.M11 * b.M12 + a.M12 * b.M22 + a.M13 * b.M32
	temp.M13 = a.M10 * b.M03 + a.M11 * b.M13 + a.M12 * b.M23 + a.M13 * b.M33

	// Row 3
	temp.M20 = a.M20 * b.M00 + a.M21 * b.M10 + a.M22 * b.M20 + a.M23 * b.M30
	temp.M21 = a.M20 * b.M01 + a.M21 * b.M11 + a.M22 * b.M21 + a.M23 * b.M31
	temp.M22 = a.M20 * b.M02 + a.M21 * b.M12 + a.M22 * b.M22 + a.M23 * b.M32
	temp.M23 = a.M20 * b.M03 + a.M21 * b.M13 + a.M22 * b.M23 + a.M23 * b.M33

	// Row 4
	temp.M30 = a.M30 * b.M00 + a.M31 * b.M10 + a.M32 * b.M20 + a.M33 * b.M30
	temp.M31 = a.M30 * b.M01 + a.M31 * b.M11 + a.M32 * b.M21 + a.M33 * b.M31
	temp.M32 = a.M30 * b.M02 + a.M31 * b.M12 + a.M32 * b.M22 + a.M33 * b.M32
	temp.M33 = a.M30 * b.M03 + a.M31 * b.M13 + a.M32 * b.M23 + a.M33 * b.M33

	return temp
}

func Mat4Transpose(a *Mat4) Mat4 {
	var temp Mat4
	
	// Row 1
	temp.M00 = a.M00
	temp.M01 = a.M10
	temp.M02 = a.M20
	temp.M03 = a.M30

	// Row 2
	temp.M10 = a.M01
	temp.M11 = a.M11
	temp.M12 = a.M21
	temp.M13 = a.M31

	// Row 3
	temp.M20 = a.M02
	temp.M21 = a.M12
	temp.M22 = a.M22
	temp.M23 = a.M32

	// Row 4
	temp.M30 = a.M03
	temp.M31 = a.M13
	temp.M32 = a.M23
	temp.M33 = a.M33

	return temp
}

func Mat4Trace(a *Mat4) float32 {
	return a.M00 + a.M11 + a.M22 + a.M33
}

func Mat4Det(a *Mat4) float32 {
	/*
		M00 M01 M02 M03
		M10 M11 M12 M13
		M20 M21 M22 M23
		M30 M31 M32 M33
	*/

	var ma, mb, mc, md Mat3

	ma = Mat3Create(
		a.M11, a.M12, a.M13,
		a.M21, a.M22, a.M23,
		a.M31, a.M32, a.M33)

	mb = Mat3Create(
		a.M10, a.M12, a.M13,
		a.M20, a.M22, a.M23,
		a.M30, a.M32, a.M33)

	mc = Mat3Create(
		a.M10, a.M11, a.M13,
		a.M20, a.M21, a.M23,
		a.M30, a.M31, a.M33)

	md = Mat3Create(
		a.M10, a.M11, a.M12,
		a.M20, a.M21, a.M22,
		a.M30, a.M31, a.M32)

	return a.M00 * Mat3Det(&ma) - a.M01 * Mat3Det(&mb) + a.M02 * Mat3Det(&mc) - a.M03 * Mat3Det(&md)
}

func Mat4Inverse(mat *Mat4) Mat4 {
	/*
		M00 M01 M02 M03
		M10 M11 M12 M13
		M20 M21 M22 M23
		M30 M31 M32 M33
	*/

	var det float32 = Mat4Det(mat)

	var a, b, c, d Mat3
	var e, f, g, h Mat3
	var i, j, k, l Mat3
	var m, n, o, p Mat3

	// Row 1
	a = Mat3Create(
		mat.M11, mat.M12, mat.M13,
		mat.M21, mat.M22, mat.M23,
		mat.M31, mat.M32, mat.M33,
	)

	b = Mat3Create(
		mat.M10, mat.M12, mat.M13,
		mat.M20, mat.M22, mat.M23,
		mat.M30, mat.M32, mat.M33,
	)

	c = Mat3Create(
		mat.M10, mat.M11, mat.M13,
		mat.M20, mat.M21, mat.M23,
		mat.M30, mat.M31, mat.M33,
	)

	d = Mat3Create(
		mat.M10, mat.M11, mat.M12,
		mat.M20, mat.M21, mat.M22,
		mat.M30, mat.M31, mat.M32,
	)

	// Row 2
	e = Mat3Create(
		mat.M01, mat.M02, mat.M03,
		mat.M21, mat.M22, mat.M23,
		mat.M31, mat.M32, mat.M33,
	)

	f = Mat3Create(
		mat.M00, mat.M02, mat.M03,
		mat.M20, mat.M22, mat.M23,
		mat.M30, mat.M32, mat.M33,
	)

	g = Mat3Create(
		mat.M00, mat.M01, mat.M03,
		mat.M20, mat.M21, mat.M23,
		mat.M30, mat.M31, mat.M33,
	)

	h = Mat3Create(
		mat.M00, mat.M01, mat.M02,
		mat.M20, mat.M21, mat.M22,
		mat.M30, mat.M31, mat.M32,
	)

	// Row 3
	i = Mat3Create(
		mat.M01, mat.M02, mat.M03,
		mat.M11, mat.M12, mat.M13,
		mat.M31, mat.M32, mat.M33,
	)

	j = Mat3Create(
		mat.M00, mat.M02, mat.M03,
		mat.M10, mat.M12, mat.M13,
		mat.M30, mat.M32, mat.M33,
	)

	k = Mat3Create(
		mat.M00, mat.M01, mat.M03,
		mat.M10, mat.M11, mat.M13,
		mat.M30, mat.M31, mat.M33,
	)

	l = Mat3Create(
		mat.M00, mat.M01, mat.M02,
		mat.M10, mat.M11, mat.M12,
		mat.M30, mat.M31, mat.M32,
	)

	// Row 4
	m = Mat3Create(
		mat.M01, mat.M02, mat.M03,
		mat.M11, mat.M12, mat.M13,
		mat.M21, mat.M22, mat.M23,
	)

	n = Mat3Create(
		mat.M00, mat.M02, mat.M03,
		mat.M10, mat.M12, mat.M13,
		mat.M20, mat.M22, mat.M23,
	)

	o = Mat3Create(
		mat.M00, mat.M01, mat.M03,
		mat.M10, mat.M11, mat.M13,
		mat.M20, mat.M21, mat.M23,
	)

	p = Mat3Create(
		mat.M00, mat.M01, mat.M02,
		mat.M10, mat.M11, mat.M12,
		mat.M20, mat.M21, mat.M22,
	)

	/*
		M00 -M01 M02 -M03
		-M10 M11 -M12 M13
		M20 -M21 M22 -M23
		-M30 M31 -M32 M33
	*/

	var temp Mat4 = Mat4Create(
		Mat3Det(&a), -Mat3Det(&b), Mat3Det(&c), -Mat3Det(&d), 
		-Mat3Det(&e), Mat3Det(&f), -Mat3Det(&g), Mat3Det(&h), 
		Mat3Det(&i), -Mat3Det(&j), Mat3Det(&k), -Mat3Det(&l), 
		-Mat3Det(&m), Mat3Det(&n), -Mat3Det(&o), Mat3Det(&p))

	var temp2 = Mat4Transpose(&temp)

	return Mat4DivScalar(&temp2, det)
}

func Mat4MulVec4(m *Mat4, v *Vec4) Vec4 {
	var temp Vec4

	temp.X = m.M00 * v.X + m.M01 * v.Y + m.M02 * v.Z + m.M03 * v.W
	temp.Y = m.M10 * v.X + m.M11 * v.Y + m.M12 * v.Z + m.M13 * v.W
	temp.Z = m.M20 * v.X + m.M21 * v.Y + m.M22 * v.Z + m.M23 * v.W
	temp.W = m.M30 * v.X + m.M31 * v.Y + m.M32 * v.Z + m.M33 * v.W

	return temp
}

func Mat4ToArray(m *Mat4) []float32 {
	var temp []float32
	temp = append(temp, m.M00)
	temp = append(temp, m.M01)
	temp = append(temp, m.M02)
	temp = append(temp, m.M03)

	temp = append(temp, m.M10)
	temp = append(temp, m.M11)
	temp = append(temp, m.M12)
	temp = append(temp, m.M13)

	temp = append(temp, m.M20)
	temp = append(temp, m.M21)
	temp = append(temp, m.M22)
	temp = append(temp, m.M23)

	temp = append(temp, m.M30)
	temp = append(temp, m.M31)
	temp = append(temp, m.M32)
	temp = append(temp, m.M33)

	return temp
}