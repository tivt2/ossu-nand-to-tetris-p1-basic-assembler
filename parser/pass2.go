package parser

func pass2(lines *[]string) string {
	out := ""
	for i, line := range *lines {
		if line[:1] == "@" {
			out += aInstruction(line[1:])
		} else {
			out += cInstruction(line)
		}

		if len(*lines) - i > 0 {
			out += "\n"
		}
	}
	return out
}