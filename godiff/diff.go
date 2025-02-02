
package godiff

import (
  //"fmt"
  "strings"
  //"strconv"
)

func Diff(a, b string) (out string) {

  // temp variables.
  switched := false
  outA := ""
  outB := ""

  // split input by line.
  aa := strings.Split(a, "\n")
  bb := strings.Split(b, "\n")

  // switch a and b so to loop over the longer one.
  if len(bb) > len(aa) { // if bb is longer, swap them. ==> aa is always the longer one.
    aa, bb = bb, aa
    switched = true
  }

  // loop over aa (the longer one).
  for i := range aa {
    if i > len(bb) {
      outA += "|" + aa[i] + "\n"
      continue
    }
    if aa[i] != bb[i] {
      outA += "|" + aa[i] + "\n"
      outB += ":" + bb[i] + "\n"
    }
  }

  // correct for having switched a and b.
  if switched {
    out = outB + outA
  } else {
    out = outA + outB
  }

  return
}

