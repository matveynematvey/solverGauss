package matrixMethods

func MultiplyMatrix(matrix *[][]float64) *[][]float64 {
	factor := 5.0

	for _, line := range *matrix {
		for ind := range line {
			line[ind] *= factor
		}
	}
	return matrix
}
