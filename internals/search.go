package internals

import "strings"

func Search(data []string, keyword string) []string {
	var result []string
	for _, name := range data {
		if strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
			result = append(result, name)
		}
	}
	return result
}
