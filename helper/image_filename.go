package helper

import (
	"strconv"
	"strings"
	"time"
)

func GetFileName(filename string) string {
	name := strings.Split(filename, ".")
	name = name[:len(name)-1]
	return strconv.Itoa(int(time.Now().Unix())) + "-" + strings.Join(name, "-")
}
