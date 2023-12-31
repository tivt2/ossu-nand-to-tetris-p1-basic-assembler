package parser

import (
	"fmt"
	"strconv"
	"unicode"
)

func getSymbolValue(symbol string) (string, bool) {
	table := map[string][]byte {
		"R0": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		"R1": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
		"R2": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0'},
		"R3": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1'},
		"R4": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0'},
		"R5": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '1'},
		"R6": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '0'},
		"R7": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1'},
		"R8": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0', '0'},
		"R9": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0', '1'},
		"R10": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '1', '0'},
		"R11": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '1', '1'},
		"R12": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '0', '0'},
		"R13": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '0', '1'},
		"R14": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0'},
		"R15": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '1'},
		"SCREEN": {'0', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		"KBD": {'0', '1', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		"SP": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		"LCL": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
		"ARG": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0'},
		"THIS": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1'},
		"THAT": {'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0'},
	}

	out, ok := table[symbol]
	return string(out), ok
}

func aInstruction(text string) string {
	if unicode.IsDigit(rune(text[0])) {
		num, err := strconv.Atoi(text)
		
		if err != nil {
			fmt.Println("Error while AInstruction: ", err)
		}

		out := decimalToBinary(num)
		return out
	} else {
		symbol, ok := getSymbolValue(text)
		if ok {
			return symbol
		}

		label, ok2 := labels[text]
		if ok2 {
			return label
		}

		variable, ok3 := variables[text]
		if ok3 {
			return variable
		}

		out := decimalToBinary(len(variables) + 16)
		variables[text] = out
		return out
	}
}