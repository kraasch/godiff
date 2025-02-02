
package main

import (
  //"fmt"
  "strings"
  //"strconv"
)

func Diff(a, b string) (out string) {
  aa := strings.Split(a, "\n")
  bb := strings.Split(b, "\n")
  outA := ""
  outB := ""
  for i := range aa {
    if i > len(bb) || aa[i] != bb[i] {
      outA += "-" + aa[i] + "\n"
      outB += "+" + bb[i] + "\n"
    }
  }
  out = outA + outB
  return 
}

