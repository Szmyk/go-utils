package text

import "unicode"

func HasSpaces (s string) bool {
	for _, r := range []rune(s) {
		if unicode.IsSpace(r) {
			return true	
		}		
	}
		
	return false
}

func HasDiacritics (s string) bool {
	for _, r := range s {
		if r > 192 {
			return true	
		}
	}

	return false
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
	
    return false
}