package kata

func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	}
	if word == false {
		return "No"
	}

	return ""
}
