package helper

import (
	"strconv"
	"strings"
	"time"
)

func GetFileName(filename string) string {
	name := strings.Split(filename, ".")
	name = name[:len(name)-1]

	img := strings.Join(name, "-")
	img = strings.Join(strings.Split(img, " "), "-")

	return strconv.Itoa(int(time.Now().Unix())) + "-" + img
}
