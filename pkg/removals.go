package pkg

import "regexp"

// RemoveComments will do just that to the given bytes
// Credit for this: https://siongui.github.io/2016/03/08/go-remove-c-cpp-style-comments/
// Note:
//   This will only remove comments from the start of a line, not a comment at the end of a line.
//   This is intentional as JSON strings can include the qualifiers used here,
//   even URL's contain the // comment qualifier, so it's just easier to go with this approach.
func RemoveComments(content string, removeEmptyLines bool) string {
	cmtGroupRegex := regexp.MustCompile(`(?m)^\s*/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	cmtRegex := regexp.MustCompile(`(?m)^\s*//.*`)
	result := cmtRegex.ReplaceAll([]byte(content), []byte(""))
	result = cmtGroupRegex.ReplaceAll(result, []byte(""))

	if removeEmptyLines {
		linesRegex := regexp.MustCompile(`(?m)^\s*\n`)
		result = linesRegex.ReplaceAll(result, []byte(""))
	}

	return string(result)
}
