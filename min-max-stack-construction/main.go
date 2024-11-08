package main

import "fmt"


type MinMaxStack struct {
	ValueArr []int
	MinArr []int
	MaxArr []int
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	stack := &MinMaxStack{
		ValueArr: make([]int, 0),
		MinArr: make([]int, 0),
		MaxArr: make([]int, 0),
	}
	return stack
}

func (stack *MinMaxStack) Peek() int {
	// Write your code here.
	if len(stack.ValueArr) > 0 {
		return stack.ValueArr[len(stack.ValueArr) - 1]
	}
	return -1
}

func (stack *MinMaxStack) Pop() int {
	// Write your code here.
	if len(stack.ValueArr) > 0 {
		val := stack.ValueArr[len(stack.ValueArr) - 1]
		stack.ValueArr = stack.ValueArr[0:len(stack.ValueArr) - 1]
		
		if len(stack.MinArr) > 0 {
			stack.MinArr = stack.MinArr[:len(stack.MinArr) - 1]
		}

		if len(stack.MaxArr) > 0 {
			stack.MaxArr = stack.MaxArr[:len(stack.MaxArr) - 1]
		}
		return val
	}
	return -1
}

func (stack *MinMaxStack) Push(number int) {
	// Write your code here.
	stack.ValueArr = append(stack.ValueArr, number)
	// minStack
	if len(stack.MinArr) == 0 || number <= stack.MinArr[len(stack.MinArr) - 1] {
		stack.MinArr = append(stack.MinArr, number)
	} else {
		stack.MinArr = append(stack.MinArr, stack.MinArr[len(stack.MinArr) - 1])
	}
	//maxStack 
	if len(stack.MaxArr) == 0 || number >= stack.MaxArr[len(stack.MaxArr) - 1] {
		stack.MaxArr = append(stack.MaxArr, number)
	} else {
		stack.MaxArr = append(stack.MaxArr, stack.MaxArr[len(stack.MaxArr) - 1])
	}
}

func (stack *MinMaxStack) GetMin() int {
	// Write your code here.
	if len(stack.MinArr) > 0 {
		return stack.MinArr[len(stack.MinArr) - 1]
	}
	return -1
}

func (stack *MinMaxStack) GetMax() int {
	// Write your code here.
	if len(stack.MaxArr) > 0 {
		return stack.MaxArr[len(stack.MaxArr) - 1]
	}
	return -1
}

func main(){
	stack := NewMinMaxStack()

	fmt.Println("stack -> ", stack)
}