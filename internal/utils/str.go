package utils

import (
	"strings"
	"time"
)

func GenClientOrderId() string {
	return strings.Replace(time.Now().Format("20060102150405.000000"), ".", "", -1)
}
