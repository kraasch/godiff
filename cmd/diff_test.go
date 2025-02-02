
package differ

import (

  // this is a test.
  "testing"

  // prints.
  "fmt"

  // other imports.
  // "strings"
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
        expectedValue: "-\n+A\n",
      },
      {
        testName:      "one-line+different-strings-01",
        inputA:        "A",
        inputB:        "",
        expectedValue: "-A\n+\n",
      },
      {
        testName:      "one-line+different-strings-02",
        inputA:        "A",
        inputB:        "B",
        expectedValue: "-A\n+B\n",
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
                       "-A\n" +
                       "-B\n" +
                       "+a\n" +
                       "+b\n",
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
                       "-B\n" +
                       "+b\n",
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
                       "-\n" +
                       "+b\n",
      },
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      exp := test.expectedValue
      got := suite.functionUnderTest(test.inputA, test.inputB)
      if exp != got {
        t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
      }
    }
  }
}

