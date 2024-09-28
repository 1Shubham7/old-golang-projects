package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StackSuite struct{
	suite.Suite
}

func (s *StackSuite) Test_Empty(){
	
	stack := NewStack()
	s.True(stack.empty)
}


func Test_IsEmpty(t *testing.T){
	
	stack := NewStack()

	expected := stack.empty

	got := stack.IsEmpty()

	assert := assert.New(t)
	
	assert.Equal(expected, got)
}