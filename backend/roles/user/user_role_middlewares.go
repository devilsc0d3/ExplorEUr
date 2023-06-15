package user

func CheckInsults(input string) bool {
	insults := []string{"fuck", "fuck off", "mother fucker", "bitch", "bastar", "bastard", "suck", "suck my dick", "fuck you", "fuck your mother"}
	for i := 0; i < len(insults); i++ {
		if insults[i] == input {
			return false
		}
	}
	return true
}

func CheckLength(input string) bool {
	if len(input) > 255 {
		return false
	}
	return true
}
