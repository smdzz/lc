package main

type MyStack struct {
	Stack []int
}

func Constructor() MyStack {
	return MyStack{Stack: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyStack) Pop() int {
	var res int
	if len(this.Stack) > 0 {
		res = this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]
		return res
	}
	return 0
}

func (this *MyStack) Top() int {
	if len(this.Stack) > 0 {
		return this.Stack[len(this.Stack)-1]
	}
	return 0
}

func (this *MyStack) Empty() bool {
	if len(this.Stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
