package main

import (
	"fmt"

	reader "./fileMethods"
	mm "./matrixMethods"
	slae "./slae"
)

func main() {
	slae := slae.SLAE{}
	reader := reader.OSReader{}

	slae.ParseSize(reader.ReadFile("slaeSize1"))
	slae.ParseMatrix(reader.ReadFile("matrix1"))
	slae.ParseRightVec(reader.ReadFile("rightVec1"))
	fmt.Println(slae)
	fmt.Println(mm.MultiplyMatrix(slae.GetMatrix()))
	fmt.Println(slae)
	slae.Gauss()
	fmt.Println(slae)
}
