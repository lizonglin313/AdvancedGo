package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {

	t1 := tree{
		value: 1,
		left:  nil,
		right: nil,
	}

	t2 := t1
	fmt.Println(t1)
	fmt.Println(t2)

	values := []int{4, 5, 2, 11, 6, 7, 1, 6}

	re := sortUsingTree(values)
	fmt.Printf("%d %d\n", len(re), cap(re))
	fmt.Printf("%p %p\n", &re, &values)
	for _, i := range values {
		fmt.Printf("%d ", i)
	}
}

func sortUsingTree(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	re := values[:0]
	//fmt.Printf("%p %p\n", &re, &values)

	appendValues(re, root)
	fmt.Println(re)     // []
	fmt.Println(values) // [1 2 4 5 6 6 7 11]
	return re
}

// 左 根 右 的顺序 将树结点存入 slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 构造tree 将值按照二叉搜索树的顺序插入
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
