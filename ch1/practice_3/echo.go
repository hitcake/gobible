package practice_3

import "strings"

func EchoV1(args []string) string {
	s, sep := "", ""
	for _, v := range args {
		s += sep + v
		sep = " "
	}
	return s
}

func EchoV2(args []string) string {
	return strings.Join(args, " ")
}
