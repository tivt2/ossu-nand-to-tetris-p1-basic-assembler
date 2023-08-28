package parser

import (
	"regexp"
	"strings"
)

func cleanLine(line *string) {
		regex := regexp.MustCompile(`//.*`)
		*line = regex.ReplaceAllString(*line, "")
		*line = strings.TrimSpace(*line)
}

func sliceLine(lines *[]string, idx int) {
	*lines = append((*lines)[:idx], (*lines)[idx+1:]...)
}

func pass1(lines *[]string) {
	removedLines := 0
	length := len(*lines)
	for i := 0; i < length; i++ {
		line := (*lines)[i - removedLines]
		cleanLine(&line)
		(*lines)[i-removedLines] = line
		
		if line == "" {
			sliceLine(lines, i - removedLines)
			removedLines++
			continue
		}

		lineLength := len(line)
		if line[:1] == "(" &&
		line[lineLength-1:] == ")" {
			label := line[1:lineLength-1]
			labels[label] = decimalToBinary(i - removedLines)
			sliceLine(lines, i - removedLines)
			removedLines++
		}
	}
}
