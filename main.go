package main

import "fmt"
//import "github.com/aod6060/golang-game-space-escape/game"
//import "github.com/aod6060/golang-game-space-escape/game"

import "github.com/aod6060/golang-game-space-escape/engine/vmath"

func main() {
	// Vec2 
	var test vmath.Vec2 = vmath.Vec2Create(6.0, 2.0)
	fmt.Println("Vec2: ", test)

	var test2 vmath.Vec2 = vmath.Vec2Create(3.0, 5.0)
	fmt.Println("Vec2: ", test2)

	fmt.Println("Addition: ", vmath.Vec2Add(&test, &test2))

	fmt.Println("Subtraction: ", vmath.Vec2Sub(&test, &test2))

	fmt.Println("Multiplication: ", vmath.Vec2Mul(&test2, 2.0))

	fmt.Println("Division: ", vmath.Vec2Div(&test2, 2.0))

	fmt.Println("Length: ", vmath.Vec2Length(&test2))

	fmt.Println("Normalize: ", vmath.Vec2Normalize(&test2))

	fmt.Println("Dot: ", vmath.Vec2Dot(&test, &test2))

	fmt.Println("Angle: ", vmath.ToDegree(vmath.Vec2Angle(&test, &test2)))

	fmt.Println("Array: ", vmath.Vec2ToArray(&test2))


	// Vec3
	var test3 vmath.Vec3 = vmath.Vec3Create(6.0, 2.0, 5.0)
	fmt.Println("Vec3: ", test3)

	var test4 vmath.Vec3 = vmath.Vec3Create(3.0, 5.0, 7.0)
	fmt.Println("Vec3: ", test4)

	fmt.Println("Addition: ", vmath.Vec3Add(&test3, &test4))

	fmt.Println("Subtraction: ", vmath.Vec3Sub(&test3, &test4))

	fmt.Println("Multiplication: ", vmath.Vec3Mul(&test3, 2.0))

	fmt.Println("Division: ", vmath.Vec3Div(&test3, 2.0))

	fmt.Println("Length: ", vmath.Vec3Length(&test3))

	fmt.Println("Normalize: ", vmath.Vec3Normalize(&test3))

	fmt.Println("Dot: ", vmath.Vec3Dot(&test3, &test4))

	fmt.Println("Angle: ", vmath.ToDegree(vmath.Vec3Angle(&test3, &test4)))

	fmt.Println("Cross: ", vmath.Vec3Cross(&test3, &test4))

	fmt.Println("Array: ", vmath.Vec3ToArray(&test3))


	// Vec4
	var test5 vmath.Vec4 = vmath.Vec4Create(6.0, 2.0, 5.0, 3.0)
	fmt.Println("Vec4: ", test5)

	var test6 vmath.Vec4 = vmath.Vec4Create(3.0, 5.0, 7.0, 9.0)
	fmt.Println("Vec4: ", test6)

	fmt.Println("Addition: ", vmath.Vec4Add(&test5, &test6))

	fmt.Println("Subtraction: ", vmath.Vec4Sub(&test5, &test6))

	fmt.Println("Multiplication: ", vmath.Vec4Mul(&test5, 2.0))

	fmt.Println("Division: ", vmath.Vec4Div(&test5, 2.0))

	fmt.Println("Lengths: ", vmath.Vec4Length(&test5))

	fmt.Println("Normalize: ", vmath.Vec4Normalize(&test5))

	fmt.Println("Dot: ", vmath.Vec4Dot(&test5, &test6))

	fmt.Println("Angle: ", vmath.ToDegree(vmath.Vec4Angle(&test5, &test6)))

	fmt.Println("Array: ", vmath.Vec4ToArray(&test5))


	// Mat2
	var test7 vmath.Mat2 = vmath.Mat2Create(1.0, 2.0, 4.0, 3.0)
	fmt.Println("Mat2: ", test7)

	var test8 vmath.Mat2 = vmath.Mat2Create(3.0, 2.0, 8.0, 7.0)
	fmt.Println("Mat2: ", test8)

	fmt.Println("Identity: ", vmath.Mat2Identiy())

	fmt.Println("Addition: ", vmath.Mat2Add(&test7, &test8))

	fmt.Println("Subtraction: ", vmath.Mat2Sub(&test7, &test8))

	fmt.Println("Multiply Scalar: ", vmath.Mat2MulScalar(&test7, 2.0))

	fmt.Println("Divide Scalar: ", vmath.Mat2DivScalar(&test7, 2.0))

	// Multiply Matrix:  {19 16 36 29}
	fmt.Println("Multiply Matrix: ", vmath.Mat2MulMatrix(&test7, &test8))


	fmt.Println("Transpose: ", vmath.Mat2Transpose(&test7))

	fmt.Println("Trace: ", vmath.Mat2Trace(&test7))

	fmt.Println("Det: ", vmath.Mat2Det(&test7))

	fmt.Println("Inverse: ", vmath.Mat2Inverse(&test7))

	var test7_inv = vmath.Mat2Inverse(&test7)

	fmt.Println("Mul Test: ", vmath.Mat2MulMatrix(&test7, &test7_inv)) // this should be identity or just abount
}