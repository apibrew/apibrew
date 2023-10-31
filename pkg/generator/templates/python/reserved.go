package python

func isPythonReservedKeyword(word string) bool {
	keywords := map[string]bool{
		"False":    true,
		"None":     true,
		"True":     true,
		"and":      true,
		"as":       true,
		"assert":   true,
		"async":    true,
		"await":    true,
		"break":    true,
		"class":    true,
		"continue": true,
		"def":      true,
		"del":      true,
		"elif":     true,
		"else":     true,
		"except":   true,
		"finally":  true,
		"for":      true,
		"from":     true,
		"global":   true,
		"if":       true,
		"import":   true,
		"in":       true,
		"is":       true,
		"lambda":   true,
		"nonlocal": true,
		"not":      true,
		"or":       true,
		"pass":     true,
		"raise":    true,
		"return":   true,
		"try":      true,
		"while":    true,
		"with":     true,
		"yield":    true,
	}

	_, exists := keywords[word]
	return exists
}
