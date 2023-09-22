package read_write

import (
	"fmt"
	"os"
)

func Write(fileName string, text string) {
	data := []byte(text)

	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err2 := f.Write(data)

	if err2 != nil {
		fmt.Println(err2)
	}
}
