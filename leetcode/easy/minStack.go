package main

import "fmt"

func main() {
    stack := Constructor()
    fmt.Printf("%T", stack)
    fmt.Println(stack)
    fmt.Println(size)
	stack.Push(-2)
    fmt.Println(stack)
    fmt.Println(size)
	stack.Push(0)
    fmt.Println(stack)
    fmt.Println(size)
	stack.Push(-3)
    fmt.Println(stack)
    fmt.Println(size)
    min := stack.GetMin()
    fmt.Printf("min=%d\n",min)
	stack.Pop()
    top := stack.Top()
    fmt.Printf("top=%d\n", top)
    min = stack.GetMin()
    fmt.Printf("min=%d\n", min)
}

type MinStack struct {
    Value int
    Min int
    Next *MinStack
}

var size = 0

func Constructor() *MinStack {
    return new(MinStack)
}


func (this *MinStack) Push(val int)  {
    if size == 0 {
        this.Value = val
        this.Min = val
        fmt.Println(this)
        size++
        return
    }

    var min int = this.Min

    if val < min {
        min = val
    }

    temp := MinStack{
        Value: val,
        Min: min,
        Next: this,
    }
    fmt.Printf("temp next = %v\n", temp.Next)
    //temp.Next = this
    *this = temp
    fmt.Printf("this next = %v\n", this.Next)
    size++
}


func (this *MinStack) Pop()  {
    fmt.Println("pop")
    fmt.Println(this)
    fmt.Println(this.Next)
}

func (this *MinStack) Top() int {
    return this.Value
}


func (this *MinStack) GetMin() int {
    return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
