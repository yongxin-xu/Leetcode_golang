package minstack

import "testing"

type MinStack struct {
	_arr []int
	_min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), 0}
}

func (this *MinStack) Push(x int) {
	if len(this._arr) == 0 {
		this._arr = append(this._arr, x)
		this._min = x
	} else {
		if x <= this._min {
			this._arr = append(this._arr, this._min)
			this._min = x
			this._arr = append(this._arr, x)
		} else {
			this._arr = append(this._arr, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this._arr) == 0 {
		return
	}
	if this.Top() == this._min {
		this._arr = this._arr[:len(this._arr)-1]
		if len(this._arr) != 0 {
			this._min = this.Top()
			this._arr = this._arr[:len(this._arr)-1]
		}
	} else {
		this._arr = this._arr[:len(this._arr)-1]
	}
}

func (this *MinStack) Top() int {
	return this._arr[len(this._arr)-1]
}

func (this *MinStack) GetMin() int {
	return this._min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMinStack1(t *testing.T) {
	minStack := MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
}

func TestMinStack2(t *testing.T) {
	minStack := MinStack{}
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	t.Log(minStack.Top())
	minStack.Pop()
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
	minStack.Push(-2147483647)
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.GetMin())
}
