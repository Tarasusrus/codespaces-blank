package prose

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T){
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want{
		t.Error(errorString(list,want,got))
	}
}

func TestTreeElements(t *testing.T){
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want{
		t.Error(errorString(list,want,got))
	}
}

func errorString(list []string, want string, got string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

func TestOneElement(t *testing.T){
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want{
		t.Error(errorString(list,want,got))
	}
}
