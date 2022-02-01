package helper

import "strings"

func SlugGenerate(name string) string {
	return strings.Join(strings.Split(strings.ToLower(name), " "), "-")
}
