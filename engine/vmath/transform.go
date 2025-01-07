package vmath

import "math"

// Translate
func TransformTranslate(v *Vec3) Mat4 {
	var temp Mat4 = Mat4Identity()
	temp.M30 = v.X
	temp.M31 = v.Y
	temp.M32 = v.Z
	return temp
}

// Rotation

// X Rotation
func TransformRotationX(radian float32) Mat4 {
	var temp Mat4 = Mat4Identity()
	temp.M11 = float32(math.Cos(float64(radian)))
	temp.M12 = float32(-math.Sin(float64(radian)))
	temp.M21 = float32(math.Sin(float64(radian)))
	temp.M22 = float32(math.Cos(float64(radian)))
	return temp
}

// Y Rotation
func TransformRotationY(radian float32) Mat4 {
	var temp Mat4 = Mat4Identity()
	temp.M00 = float32(math.Cos(float64(radian)))
	temp.M02 = float32(math.Sin(float64(radian)))
	temp.M20 = float32(-math.Sin(float64(radian)))
	temp.M22 = float32(math.Cos(float64(radian)))
	return temp
}

// Z Rotation
func TransformRotationZ(radian float32) Mat4 {
	var temp Mat4 = Mat4Identity()
	temp.M00 = float32(math.Cos(float64(radian)))
	temp.M01 = float32(-math.Sin(float64(radian)))
	temp.M10 = float32(math.Sin(float64(radian)))
	temp.M11 = float32(math.Cos(float64(radian)))
	return temp
}

// Combined Rotation
/*
	c = cos(radian), s = sin(radian), ||(x, y, z)|| = 1 if not I'll have to normalize the vector
*/

func TransformRotate(radian float32, v *Vec3) Mat4 {
	var r float64 = float64(radian)
	var c float32 = float32(math.Cos(r))
	var s float32 = float32(math.Sin(r))

	if Vec3Length(v) != 1 {
		var temp = Vec3Normalize(v)
		v.X = temp.X
		v.Y = temp.Y
	}

	var temp Mat4 = Mat4Identity()
	temp.M00 = (v.X * v.X) * (1 - c) + c
	temp.M01 = (v.X * v.Y) * (1 - c) - v.Z * s
	temp.M02 = (v.X * v.Z) * (1 - c) + v.Y * s
	temp.M10 = (v.X * v.Y) * (1 - c) + v.Z * s
	temp.M11 = (v.Y * v.Y) * (1 - c) + c
	temp.M12 = (v.Y * v.Z) * (1 - c) - v.X * s
	temp.M20 = (v.X * v.Z) * (1 - c) - v.Y * s
	temp.M21 = (v.Y * v.Z) * (1 - c) + v.X * s
	temp.M22 = (v.Z * v.Z) * (1 - c) + c

	return temp
}

// Scaling
func TransformScale(v *Vec3) Mat4 {
	var temp Mat4 = Mat4Identity()
	temp.M00 = v.X
	temp.M11 = v.Y
	temp.M22 = v.Z
	return temp
}


// Projection Transformations
func TransformOrtho(left float32, right float32, bottom float32, top float32, zmin float32, zmax float32) Mat4 {
	var temp Mat4 = Mat4Identity()
	var tx float32 = (right + left) / (right - left)
	var ty float32 = (top + bottom) / (top - bottom)
	var tz float32 = (zmax + zmin) / (zmax - zmin)

	temp.M00 = 2.0 / (right - left)
	temp.M11 = 2.0 / (top - bottom)
	temp.M22 = -2.0 / (zmax - zmin)
	temp.M30 = tx
	temp.M31 = ty
	temp.M32 = tz

	return temp
}

func TransformOrtho2D(left float32, right float32, bottom float32, top float32) Mat4 {
	return TransformOrtho(left, right, bottom, top, -1.0, 1.0)
}

func TransformFrustrum(left float32, right float32, bottom float32, top float32, zmin float32, zmax float32) Mat4 {
	var temp Mat4 = Mat4Identity()

	var A float32 = (right + left) / (right - left)
	var B float32 = (top + bottom) / (top - bottom)
	var C float32 = -((zmax + zmin) / (zmax - zmin))
	var D float32 = -((2.0 * zmax * zmin) / (zmax - zmin))

	temp.M00 = (2.0 * zmin) / (right - left)
	temp.M11 = (2.0 * zmin) / (top - bottom)
	temp.M20 = A
	temp.M21 = B
	temp.M22 = C
	temp.M23 = -1
	temp.M32 = D

	return temp
}

func TransformPerspective(fov float32, aspect float32, zmin float32, zmax float32) Mat4 {
	var ymax, xmax float32
	//var temp, temp1, temp2, temp3 float32
	ymax = zmin * float32(math.Tan(float64(fov)))
	xmax = ymax * aspect
	return TransformFrustrum(-xmax, xmax, -ymax, ymax, zmin, zmax)
}