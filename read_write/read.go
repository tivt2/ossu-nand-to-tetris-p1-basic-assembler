package read_write

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filePath string) []string {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := []string{}
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}
