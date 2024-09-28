package cal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//let us now see paramitarized tests:

var testcases = []struct {
	name          string
	a             int
	b             int
	excepted      int
	exceptedError error
}{
	{"first test case", 10, 5, 2, nil},
	{"second test case", 4, 2, 2, nil},
	// {"third test case", 10, 0,0,  errors.New("division by zero")},
}

func TestDivide(t *testing.T) {
	for _, i := range testcases {
		t.Run(i.name, func(t *testing.T) {

			assert := assert.New(t)

			exp := i.excepted

			gotAns, gotError := Divide(i.a, i.b)
			
			
			assert.Equal(i.exceptedError, gotError)
			assert.Equal(gotAns, exp)
		})
	}
}
