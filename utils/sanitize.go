package utils

import (
	"strings"
)

func SanitizeString(s string) string {
	var result strings.Builder
	for _, r := range s {
		if isAllowedCharacter(r) {
			result.WriteRune(r)
		}
		// Skip any character that's not allowed
	}
	return result.String()
}

func isAllowedCharacter(r rune) bool {
	// ASCII printable characters (32-126) including space
	if r >= 32 && r <= 126 {
		return true
	}
	// Persian/Farsi characters
	if (r >= 0x0600 && r <= 0x06FF) || // Arabic block (includes Persian)
		(r >= 0xFB50 && r <= 0xFDFF) || // Arabic Presentation Forms-A
		(r >= 0xFE70 && r <= 0xFEFF) { // Arabic Presentation Forms-B
		return true
	}

	//U+1F300 - U+1F5FF misc
	//U+1F680 - U+1F6FF transport and map symbols
	//U+1F100 - U+1F1FF regional
	// Emojis
	if (r >= 0x1F100 && r <= 0x1F1FF) || // regional
		(r >= 0x1F680 && r <= 0x1F6FF) || // transport
		(r >= 0x1F700 && r <= 0x1F77F) || // Alchemical Symbols
		(r >= 0x1F780 && r <= 0x1F7FF) || // Geometric Shapes Extended
		(r >= 0x1F800 && r <= 0x1F8FF) || // Supplemental Arrows-C
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(r >= 0x1FA00 && r <= 0x1FA6F) || // Chess Symbols
		(r >= 0x1FA70 && r <= 0x1FAFF) || // Symbols and Pictographs Extended-A
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) { // Dingbats
		return true
	}

	return false
}
