package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"nand-p1/assembler/parser"
	"nand-p1/assembler/read_write"
)

func main() {
	start := time.Now()
	filePath := os.Args[1]

	lines := read_write.Read(filePath)

	parsedData := parser.Parse(lines)

	regex := regexp.MustCompile(`^(.*?)(\.asm)$`)
	extracted := regex.FindStringSubmatch(filePath)

	if len(extracted) >= 3 {
		fileName := fmt.Sprintf("%s.hack", extracted[1])
		read_write.Write(fileName, parsedData)
	} else {
		fmt.Println("File doesnt match")
	}
	fmt.Printf("Time taken to compile: %s\n", time.Since(start))
}
