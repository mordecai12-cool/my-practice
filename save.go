if cmd == "(up," {
	for j := i - count; j < i; j++ {
		if j >= 0 {
			words[j] = strings.ToUpper(words[j])
		}
	}
}
if cmd == "(low," {
	for j := i - count; j < i; j++ {
		if j >= 0 {
			words[j] = strings.ToLower(words[j])
		}
	}
}
if cmd == "(cap," {
	for j := i - count; j < i; j++ {
		if j >= 0 {
			words[j] = strings.ToUpper(words[j][:1]) + strings.ToLower(words[j][1:])
		}
	}
}