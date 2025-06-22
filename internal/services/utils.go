package services

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var (
	ErrFileMissing = errors.New("file missing")
)

func FileExists(fPath string) bool {

	switch _, err := os.Stat(fPath); {
	case errors.Is(err, os.ErrNotExist):
		return false
	case err != nil:
		return false
	}

	return true
}

func EscapeChar(s string, charToEscape rune) string {
	var result strings.Builder
	for _, r := range s {
		if r == charToEscape {
			result.WriteString("\\")
		}
		result.WriteRune(r)
	}
	return result.String()
}

func EscapeMultipleChars(s string, charsToEscape []rune) string {
	var result strings.Builder
	escapeSet := make(map[rune]bool)
	for _, char := range charsToEscape {
		escapeSet[char] = true
	}

	for _, r := range s {
		if _, shouldEscape := escapeSet[r]; shouldEscape {
			result.WriteString("\\")
		}
		result.WriteRune(r)
	}
	return result.String()
}

func cleanString(s string) string {
	return strings.Trim(s, " ")
}

func quoteSubstring(story, frenzy, mark string) string {
	return strings.ReplaceAll(story, frenzy, mark+frenzy+mark)
}

func frenzyBold(story, frenzy string) string {
	quotted := quoteSubstring(story, frenzy, "*")
	// if frenzy was capitalized in story
	quotted = quoteSubstring(quotted, capitalizeFirst(frenzy), "*")

	return quotted 
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	firstRune := []rune(s)[0]
	cappedFirst := string(unicode.ToUpper(firstRune))
	restOfString := string([]rune(s)[1:])
	return cappedFirst + restOfString
}

func filesCountByMask(mask string) (int, error) {
	matches, err := filepath.Glob(mask)
	if err != nil {
		return 0, err
	}

	return len(matches), nil
}