package kata

import (
	"fmt"
	"strings"
)

func PartList(arr []string) string {
	s := arr[:]
	var res strings.Builder

	for i := range arr {
		max := len(arr) - 1

		if i < max {
			p1 := s[:i+1]
			p2 := s[i+1:]

			fmt.Fprintf(&res, "(%s, %s)", strings.Join(p1, " "), strings.Join(p2, " "))
		}
	}
	return res.String()
}
