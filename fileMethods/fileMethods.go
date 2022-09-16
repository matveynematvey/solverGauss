package fileMethods

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) (buf string) {
	file, err := os.Open(fileName)

	check(err, "ReaderOS fail")
	defer func() { check(file.Close(), "File closing error") }()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		buf += scanner.Text() + " "
	}

	return
}

func WriteMatrixToFile(matrix *[][]int, fileName string) {
	file, err := os.Create(fileName)
	check(err, "Error creating file")

	for _, line := range *matrix {
		file.WriteString(fmt.Sprintln(line))
	}
}

func check(err error, text string) {
	if err != nil {
		fmt.Println(text)
	}
}
