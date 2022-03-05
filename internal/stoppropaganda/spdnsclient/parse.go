package spdnsclient

// import "os"

// func open(name string) (*file, error) {
// 	fd, err := os.Open(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &file{fd, make([]byte, 0, 64*1024), false}, nil
// }

// stringsHasSuffixFold reports whether s ends in suffix,
// ASCII-case-insensitively.
func stringsHasSuffixFold(s, suffix string) bool {
	return len(s) >= len(suffix) && stringsEqualFold(s[len(s)-len(suffix):], suffix)
}

// stringsHasPrefix is strings.HasPrefix. It reports whether s begins with prefix.
func stringsHasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// stringsEqualFold is strings.EqualFold, ASCII only. It reports whether s and t
// are equal, ASCII-case-insensitively.
func stringsEqualFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if lowerASCII(s[i]) != lowerASCII(t[i]) {
			return false
		}
	}
	return true
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// Number of occurrences of b in s.
func count(s string, b byte) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			n++
		}
	}
	return n
}
