package utils

func StartsWith(str string, look string) bool {
	l := len(look)

	compare := str[0:l]

	return compare == look
}
