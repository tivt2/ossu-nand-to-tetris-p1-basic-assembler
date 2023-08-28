package parser

import (
	"regexp"
)

func destInstruction(arr []byte, instruction string) {
	table := map[string][]byte {
		"M": {'0','0','1'},
		"D": {'0','1','0'},
		"MD": {'0','1','1'},
		"A": {'1','0','0'},
		"AM": {'1','0','1'},
		"AD": {'1','1','0'},
		"ADM": {'1','1','1'},
	}

	dest := table[instruction]
	copy(arr[10:13], dest)	
}

func jumpInstruction(arr []byte, instruction string) {
	table := map[string][]byte {
		"JGT": {'0','0','1'},
		"JEQ": {'0','1','0'},
		"JGE": {'0','1','1'},
		"JLT": {'1','0','0'},
		"JNE": {'1','0','1'},
		"JLE": {'1','1','0'},
		"JMP": {'1','1','1'},
	}

	jump := table[instruction]
	copy(arr[13:16], jump)
}

func compInstruction(arr []byte, instruction string) {
	table := map[string][]byte {
		"0": {'0','1','0','1','0','1','0'},
		"1": {'0','1','1','1','1','1','1'},
		"-1": {'0','1','1','1','0','1','0'},
		"D": {'0','0','0','1','1','0','0'},
		"A": {'0','1','1','0','0','0','0'},
		"M": {'1','1','1','0','0','0','0'},
		"!D": {'0','0','0','1','1','0','1'},
		"!A": {'0','1','1','0','0','0','1'},
		"!M": {'1','1','1','0','0','0','1'},
		"-D": {'0','0','0','1','1','1','1'},
		"-A": {'0','1','1','0','0','1','1'},
		"-M": {'1','1','1','0','0','1','1'},
		"D+1": {'0','0','1','1','1','1','1'},
		"A+1": {'0','1','1','0','1','1','1'},
		"M+1": {'1','1','1','0','1','1','1'},
		"D-1": {'0','0','0','1','1','1','0'},
		"A-1": {'0','1','1','0','0','1','0'},
		"M-1": {'1','1','1','0','0','1','0'},
		"D+A": {'0','0','0','0','0','1','0'},
		"D+M": {'1','0','0','0','0','1','0'},
		"D-A": {'0','0','1','0','0','1','1'},
		"D-M": {'1','0','1','0','0','1','1'},
		"A-D": {'0','0','0','0','1','1','1'},
		"M-D": {'1','0','0','0','1','1','1'},
		"D&A": {'0','0','0','0','0','0','0'},
		"D&M": {'1','0','0','0','0','0','0'},
		"D|A": {'0','0','1','0','1','0','1'},
		"D|M": {'1','0','1','0','1','0','1'},
	}

	comp := table[instruction]
	copy(arr[3:10], comp)
}

func cInstruction(text string) string {
	out := []byte{'1', '1', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	pattern := "=|;"
	regex := regexp.MustCompile(pattern)
	splited := regex.Split(text, -1)

	equal := "="
	semicolon := ";"
	regex2 := regexp.MustCompile(equal)
	regex3 := regexp.MustCompile(semicolon)
	hasEqual := regex2.Match([]byte(text))
	hasSemicolon := regex3.Match([]byte(text))
	
	if hasEqual {
		destInstruction(out, splited[0])
		compInstruction(out, splited[1])
		if hasSemicolon {
			jumpInstruction(out, splited[2])
		}
	} else {
		compInstruction(out, splited[0])
		if hasSemicolon {
			jumpInstruction(out, splited[1])
		}
	}

	return string(out)
}