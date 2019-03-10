package kata

import (
  "strings"
)

func Accum(s string) string {
  str := ""
  sep := "-"
  for i, r := range s {
    str += strings.ToUpper(string(r)) + strings.Repeat(strings.ToLower(string(r)), i) + sep
  }
  trim := strings.TrimSuffix(str, "-")
  
  return trim
}
