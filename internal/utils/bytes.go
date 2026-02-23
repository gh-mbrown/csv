package utils

func IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsOperator(ch byte) bool {
	return ch == '<' || ch == '>' || ch == '!' || ch == '='
}

func IsPeriod(ch byte) bool {
	return ch == '.'
}

func IsSeparator(ch byte) bool {
	return ch == ','
}
