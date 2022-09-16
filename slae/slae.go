package SLAE

import (
	"math"
	"strconv"
	"strings"
)

type SLAE struct {
	size     int
	matrix   [][]float64
	rightVec []float64
}

func (m *SLAE) GetMatrix() *[][]float64 {
	return &m.matrix
}

func (m *SLAE) ParseSize(buf string) {
	sliceString := strings.Fields(buf)

	m.size, _ = strconv.Atoi(sliceString[0])
}

func (m *SLAE) ParseMatrix(buf string) {
	sliceString := strings.Fields(buf)

	m.matrix = make([][]float64, m.size)

	for ind := range m.matrix {
		m.matrix[ind] = make([]float64, m.size)
	}

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			m.matrix[i][j], _ = strconv.ParseFloat(sliceString[i+j], 64)
		}
	}
}

func (m *SLAE) ParseRightVec(buf string) {
	sliceString := strings.Fields(buf)

	m.rightVec = make([]float64, m.size)

	for i := 0; i < m.size; i++ {
		m.rightVec[i], _ = strconv.ParseFloat(sliceString[i], 64)
	}
}

func (m *SLAE) elem(i int) {
	line := i
	for j := i + 1; j < m.size; j++ {
		if math.Abs(m.matrix[j][i]) > math.Abs(m.matrix[line][i]) {
			line = j
		}

		if line != i {
			var c float64
			for j := i; j < m.size; j++ {
				c = m.matrix[i][j]
				m.matrix[i][j] = m.matrix[line][j]
				m.matrix[line][j] = c
			}
			c = m.rightVec[i]
			m.rightVec[i] = m.rightVec[line]
			m.rightVec[line] = c
		}
	}
}

func (m *SLAE) Gauss() {
	for i := 0; i < m.size; i++ {
		m.elem(i)
		A_d := m.matrix[i][i]

		for p := i; p < m.size; p++ {
			m.matrix[i][p] /= A_d
		}
		m.rightVec[i] /= A_d

		for j := i + 1; j < m.size; j++ {
			A_j := m.matrix[j][i]
			if A_j != 0 {
				for k := i; k < m.size; k++ {
					m.matrix[j][k] -= m.matrix[i][k] * A_j
				}
				m.rightVec[j] -= m.rightVec[i] * A_j
			}
		}
	}
	for i := m.size - 1; i >= 0; i-- {
		for j := i + 1; j < m.size; j++ {
			m.rightVec[i] -= m.matrix[i][j] * m.rightVec[j]
		}
	}
}
