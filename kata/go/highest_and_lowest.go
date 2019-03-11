package kata

import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	sep := strings.Split(in, " ")
	vals := make([]int, len(sep))
	for i, v := range sep {
		vals[i], _ = strconv.Atoi(v)
	}
	sort.Ints(vals)
	maxMin := []string{strconv.Itoa(vals[len(sep)-1]), strconv.Itoa(vals[0])}
	return strings.Join(maxMin, " ")
}
