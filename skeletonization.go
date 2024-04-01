package GoSkeletonPy

import "strings"

// GeneratePythonSkeleton takes a slice of strings representing lines of Python code,
// and returns a new slice where function bodies and variable assignments are replaced with '...'.
// It also replaces tabs with spaces and removes comments.
func GeneratePythonSkeleton(lines []string) []string {
	ellipsis := false

	var output []string

	for _, line := range lines {

		// replace tabs with spaces
		line = strings.ReplaceAll(line, "\t", "    ")

		// remove everything after the # symbol
		if strings.Contains(line, "#") {
			line = strings.Split(line, "#")[0]
		}

		// check if line has def or class as prefix
		if !(strings.HasPrefix(strings.TrimSpace(line), "def") || (strings.Contains(strings.ToLower(line), "from") && strings.Contains(strings.ToLower(line), "import")) || strings.HasPrefix(strings.TrimSpace(line), "class")) {

			// detect if line creates a variable
			if strings.Contains(line, "=") && !strings.Contains(line, "==") && !strings.Contains(line, "!=") && !strings.Contains(line, "<=") && !strings.Contains(line, ">=") && !strings.Contains(line, "<=") && !strings.Contains(line, ">=") {
				varname := strings.Split(line, "=")[0]
				if strings.Contains(varname, "(") || strings.Contains(varname, "[") || strings.Contains(varname, "{") {
					continue
				}

				ellipsis = true
				output = append(output, strings.Split(line, "=")[0]+" = ...")
				continue
			}

			if !ellipsis {
				ellipsis = true

				indent := 0
				for _, char := range line {
					if char == ' ' {
						indent++
					} else {
						break
					}
				}

				// append the line to the output
				if indent == 0 {
					continue
				}
				output = append(output, strings.Repeat(" ", indent)+"return")

			} else {
				continue
			}
		} else {
			ellipsis = false
			output = append(output, line)
		}
	}
	var formatedOutput []string
	for _, line := range output {
		formatedOutput = append(formatedOutput, strings.ReplaceAll(line, "    ", "\t"))
	}
	return formatedOutput
}
