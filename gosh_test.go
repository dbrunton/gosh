package gosh

import(
 "testing"
)

var goshTests = []string{
 "gosh>",
 "gosh> ",
}

func TestGosh(t *testing.T) {
 for _, gt := range goshTests {
  prompt := Gosh()
  if prompt != gt { //fail! :)
   t.Errorf("Prompt is wrong: '%s'\n", gt)
  }
 }
}
