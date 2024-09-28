package cal

import (
	"testing"
)

//let us now see paramitarized tests:

var testcases = []struct {
	name string
	a int
	b int
	excepted int
	exceptedError error
} {
	{"first test case", 10, 5, 2,  nil},
	{"second test case", 4,2, 2,  nil},
	// {"third test case", 10, 0,0,  errors.New("division by zero")},
}

func TestDivide(t *testing.T){
	for _,i := range testcases {
		t.Run(i.name, func(t *testing.T){
			exp := i.excepted
			gotAns, gotError := Divide(i.a,i.b);
			if (gotError != nil){
				if (gotError.Error() != i.exceptedError.Error()){
					t.Errorf("expected error occured")
				}
			}
			if (gotAns != exp){
				t.Errorf("Test is failing")
			}
		})
	}
}