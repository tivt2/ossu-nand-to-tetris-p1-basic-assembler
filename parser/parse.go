package parser

var variables = map[string]string{}
var labels = map[string]string{}

func Parse(lines []string) string {
	pass1(&lines)
	out := pass2(&lines)
	return out
}
