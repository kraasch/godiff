
package godiff

import (
  "strings"
)

const (
  red   = "\033[31m"
  green = "\033[32m"
  reset = "\033[0m"
)

func Diff(a string, b string) string {
  return coloredDiff(a, b, false)
}

func CDiff(a string, b string) string {
  return coloredDiff(a, b, true)
}

func coloredDiff(a string, b string, hasColor bool) (out string) {

  // temp variables.
  switched := false
  outA := ""
  outB := ""

  if hasColor {
    outA += red
    outB += green
  }

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
    if i >= len(bb) { // i cannot index bb anymore.
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
  if hasColor {
    out += reset
  }

  return
}

