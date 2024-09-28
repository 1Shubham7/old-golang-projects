package main

import "fmt"

type Stack struct{
    data []int
    top int
    max_size int
}

func main() {
    stack  := new(Stack)
    fmt.Println("Tell me the size of the stack")
    fmt.Scanf("%d", &stack.max_size)
    iniStack(stack)
    
    push(stack, 100)
    push(stack, 200)
    push(stack, 300)
    
    pop(stack)
    pop(stack)
    pop(stack)
}

func (s *Stack) iniStack(){
    s.top = -1
}

func (s *Stack) isEmpty() bool{
    if (s.top == -1){
        return true
    }
    return false
}

func (s *Stack) isFull() bool{
    if (s.top == s.max_size-1){
        return true
    }
    return false
}

func (s *Stack) pop (int, error) {
    if (isEmpty(s)){
		err := fmt.Errorf("the stack is underflow")
		return -1, err
	}
    x := s.data[s.top]
    s.top--
    return x, nil
}

func (s *Stack) push(value int) error{
    if (isFull(s)){
		err := fmt.Errorf("the stack is underflow")
		return err
	}
    s.top++
    s.data[s.top] = value
    return nil
}
