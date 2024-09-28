package stack

import (
	"testing"
)

func Test_IsEmpty(t *testing.T){
	
	stack := NewStack()

	expected := stack.empty

	got := stack.IsEmpty()

	if (expected != got){
		t.Errorf("Error occured")
	}
}