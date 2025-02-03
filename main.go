
package main

import (
  "fmt"
  "github.com/kraasch/godiff/godiff"
)

func main() {
  diff := godiff.CDiff("a", "b")
  fmt.Print(diff)
}

