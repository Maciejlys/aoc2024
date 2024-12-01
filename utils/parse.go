package utils

import "strings"

func Parse(fileContent string, separator string) [][]string {
	lines := strings.Split(fileContent, "\n")
	result := make([][]string, 0)

	for _, line := range lines {
		parts := strings.Split(line, separator)
		for i, part := range parts {
			parts[i] = strings.TrimSpace(part)
		}

		result = append(result, parts)
	}

	return result
}
