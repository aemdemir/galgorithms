package search

// BoyerMooreSearch makes pattern matching. Pattern matching is done backwards.
func BoyerMooreSearch(source, pattern string) int {
	if len(pattern) < 1 || len(pattern) > len(source) { return -1 }

	skipTable := buildSkipTable(pattern)
	
	i := len(pattern) - 1
	for i < len(source) {
		found := match(source, pattern, i)
		if found != -1 {
			return found
		}

		skip, ok := skipTable[string(source[i])]
		if ok {
			if skip == 0 { skip = 1 }
			i = i + skip
		} else {
			i = i + len(pattern)
		}
	}

	return -1
}

func buildSkipTable(pattern string) map[string]int {
	skipTable := make(map[string]int)
	for i, c := range pattern {
		skipTable[string(c)] = len(pattern) - 1 - i
	}
	return skipTable
}

func match(source, pattern string, index int) int {
	if index < 0 || index > len(source)-1 {
		return -1
	}

	if source[index] != pattern[len(pattern)-1] {
		return -1
	}

	if len(pattern) == 1 {
		return index
	}

	return match(source, pattern[:len(pattern)-1], index-1)
}
