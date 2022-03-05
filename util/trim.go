package util

// to get rid of \r or \n (removes last char)

func TrimString(str string) string {
	for str[len(str)-1] == '\n' || str[len(str)-1] == '\r' {
		str = str[0 : len(str)-1]
	}
	return str
}
