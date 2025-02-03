
package godiff

import (

  // this is a test.
  "testing"
  "github.com/google/go-cmp/cmp"

  // prints.
  "fmt"

  // other imports.
  // "strings"
  // "reflect"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  inputA            string
  inputB            string
  expectedValue     string
}

type TestSuite struct {
  functionUnderTest func(string, string) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function TEMPLATE().
  */
  {
    nil,
    []TestList{
    },
  },
  /*
  * Test for the function Diff().
  */
  {
    functionUnderTest: Diff,
    tests: []TestList{
      {
        testName:      "one-line+same-strings-00",
        inputA:        "",
        inputB:        "",
        expectedValue: "",
      },
      {
        testName:      "one-line+same-strings-01",
        inputA:        "A",
        inputB:        "A",
        expectedValue: "",
      },
      {
        testName:      "one-line+same-strings-02",
        inputA:        "longer line: \t WITH some Words! in it",
        inputB:        "longer line: \t WITH some Words! in it",
        expectedValue: "",
      },
      {
        testName:      "one-line+different-strings-00",
        inputA:        "",
        inputB:        "A",
        expectedValue:
                       "|\n" +
                       ":A\n",
      },
      {
        testName:      "one-line+different-strings-01",
        inputA:        "A",
        inputB:        "",
        expectedValue:
                       "|A\n" +
                       ":\n",
      },
      {
        testName:      "one-line+different-strings-02",
        inputA:        "A",
        inputB:        "B",
        expectedValue:
                       "|A\n" +
                       ":B\n",
      },
      {
        testName:      "multi-line+same-strings-00",
        inputA:
                       "A\n" +
                       "B\n",
        inputB:
                       "A\n" +
                       "B\n",
        expectedValue:
                       "",
      },
      {
        testName:      "multi-line+same-strings-01",
        inputA:
                       "\n" +
                       "\n",
        inputB:
                       "\n" +
                       "\n",
        expectedValue:
                       "",
      },
      {
        testName:      "multi-line+same-strings-02",
        inputA:
                       "A 1 \n" +
                       "B b\n",
        inputB:
                       "A 1 \n" +
                       "B b\n",
        expectedValue:
                       "",
      },
      {
        testName:      "multi-line+different-strings-00",
        inputA:
                       "A\n" +
                       "B\n",
        inputB:
                       "a\n" +
                       "b\n",
        expectedValue:
                       "|A\n" +
                       "|B\n" +
                       ":a\n" +
                       ":b\n",
      },
      {
        testName:      "multi-line+different-strings-01",
        inputA:
                       "A\n" +
                       "B\n",
        inputB:
                       "A\n" +
                       "b\n",
        expectedValue:
                       "|B\n" +
                       ":b\n",
      },
      {
        testName:      "multi-line+different-strings-02",
        inputA:
                       "A\n" +
                       "\n",
        inputB:
                       "A\n" +
                       "b\n",
        expectedValue:
                       "|\n" +
                       ":b\n",
      },
      {
        testName:      "more-complicated-00",
        inputA:
                       "Mo Tu We Th Fr Sa Su" + NL +
                       "       1  2  3  4  5" + NL,
        inputB:
                       "Mo Tu We Th Fr Sa Su" + NL +
                       "                1  2" + NL,
        expectedValue:
                       "|       1  2  3  4  5" + NL +
                       ":                1  2" + NL,
      },
      {
        testName:      "more-empty-lines-in-a-than-b-00",
        inputA:
                       "A" + NL +
                       ""  + NL +
                       ""  + NL,
        inputB:
                       "B" + NL +
                       ""  + NL,
        expectedValue:
                       "|A" + NL +
                       "|"  + NL +
                       ":B" + NL,
      },
      {
        testName:      "more-empty-lines-in-b-than-a-00",
        inputA:
                       "A" + NL +
                       ""  + NL,
        inputB:
                       "B" + NL +
                       ""  + NL +
                       ""  + NL,
        expectedValue:
                       ":A" + NL +
                       "|B" + NL +
                       "|"  + NL,
      },
      {
        testName:      "example-00",
        inputA:
                       "Mo Tu We Th Fr Sa Su" + NL +
                       "       1  2  3  4  5" + NL +
                       " 6  7  8  9 10 11 12" + NL +
                       "13 14 15 16 17 18 19" + NL +
                       "20 21 22 23 24 25 26" + NL +
                       "27 28"                + NL +
                       ""                     + NL,
        inputB:
                       "Mo Tu We Th Fr Sa Su" + NL +
                       "                1  2" + NL +
                       " 3  4  5  6  7  8  9" + NL +
                       "10 11 12 13 14 15 16" + NL +
                       "17 18 19 20 21 22 23" + NL +
                       "24 25 26 27 28"       + NL +
                       ""                     + NL +
                       ""                     + NL,
        expectedValue:
                       ":       1  2  3  4  5" + NL +
                       ": 6  7  8  9 10 11 12" + NL +
                       ":13 14 15 16 17 18 19" + NL +
                       ":20 21 22 23 24 25 26" + NL +
                       ":27 28"                + NL +
                       "|                1  2" + NL +
                       "| 3  4  5  6  7  8  9" + NL +
                       "|10 11 12 13 14 15 16" + NL +
                       "|17 18 19 20 21 22 23" + NL +
                       "|24 25 26 27 28"       + NL +
                       "|"                     + NL,
      },
      //{
      //  testName:      "example-01",
      //  inputA:
      //                 "Mo Tu We Th Fr Sa Su" + NL +
      //                 "       1  2  3  4  5" + NL +
      //                 " 6  7  8  9 10 11 12" + NL +
      //                 "13 14 15 16 17 18 19" + NL +
      //                 "20 21 22 23 24 25 26" + NL +
      //                 "27 28"                + NL +
      //                 ""                     + NL,
      //  inputB:
      //                 "Mo Tu We Th Fr Sa Su" + NL +
      //                 "       1  2  3  4  5" + NL +
      //                 " 6  7  8  9 10 11 12" + NL +
      //                 "13 14 25 16 17 18 19" + NL +
      //                 "20 21 22 23 24 25 26" + NL +
      //                 "27 28"                + NL +
      //                 ""                     + NL,
      //  expectedValue:
      //                 "|13 14 15 17 17 18 19" + NL +
      //                 ":13 14 25 16 17 18 19" + NL,
      //},
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      t.Run(name, func(t *testing.T) {
        exp := test.expectedValue
        got := suite.functionUnderTest(test.inputA, test.inputB)

        // if exp != got {
        //   t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
        // }

        // if !reflect.DeepEqual(exp, got) {
        //   t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
        // }

        diff := cmp.Diff(exp, got)
        if diff != "" {
          t.Fatalf("\n%s\n", diff)
        }

      })
    }
  }
}

