package helpers

import (
	"fmt"
	"strings"
)

// StrReplace replace placeholders by it value in a string.
func StrReplace(str string, old string, new interface{}) string {
	r := strings.NewReplacer(old, fmt.Sprintf("%v", new))
	str = r.Replace(str)
	return str
}
