package strings

func MyValid(s string) bool {
	if s == "" {
		return true
	}

	for i := 0; i < len(s); {
		b := s[i]

		switch {
		case b < 128:
			i++
		case b >= 192 && b <= 223:
			if i+1 < len(s) && s[i+1] <= 191 && s[i+1] >= 128 {
				i += 2
			} else {
				return false
			}
		case b >= 224 && b <= 239:
			if i+2 < len(s) && s[i+2] >= 128 && s[i+1] >= 128 &&
				s[i+1] <= 191 && s[i+2] <= 191 {
				i += 3
			} else {
				return false
			}
		case b >= 240 && b <= 247:
			if i+3 < len(s) && s[i+2] >= 128 && s[i+1] >= 128 &&
				s[i+1] <= 191 && s[i+2] <= 191 &&
				s[i+3] >= 128 && s[i+3] <= 191 {
				i += 4
			} else {
				return false
			}
		default:
			return false
		}
	}

	return true
}
