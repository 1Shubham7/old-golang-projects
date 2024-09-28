package cal

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T){
	excepted := 10/2
	got := Divide(10,2)

	if (got != excepted) {
		t.Errorf("got an error");
	} else {
		fmt.Print("MAJEEE")
	}
}