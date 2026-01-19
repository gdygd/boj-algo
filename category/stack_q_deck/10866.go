package main

import "fmt"

type Node struct {
	next  *Node
	prev  *Node
	Value int
}

type Dequeu struct {
	root  *Node
	tail  *Node
	count int
}

func NewDeque() *Dequeu {
	return &Dequeu{
		root:  nil,
		tail:  nil,
		count: 0,
	}
}

func (d *Dequeu) PushFront(x int) {
	node := &Node{
		Value: x,
	}

	if d.root == nil {
		// 첫 노드
		d.root = node
		d.tail = node
	}

	node.next = d.root
	d.root.prev = node
	d.root = node
	d.count++
}

func (d *Dequeu) PushBack(x int) {
	node := &Node{
		Value: x,
	}

	if d.tail == nil {
		d.root = node
		d.tail = node
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
	d.count++
}

func (d *Dequeu) Popfront() int {
	if d.count == 0 || d.root == nil {
		return -1
	}

	n := d.root
	d.root = d.root.next
	if d.root != nil {
		d.root.prev = nil
	} else {
		d.tail = nil
	}
	n.next = nil

	d.count--

	if d.count == 0 {
		d.root = nil
		d.tail = nil
	}
	return n.Value
}

func (d *Dequeu) Popback() int {
	if d.count == 0 || d.tail == nil {
		return -1
	}

	n := d.tail

	d.tail = d.tail.prev

	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.root = nil
	}

	n.prev = nil
	d.count--

	if d.count == 0 {
		d.root = nil
		d.tail = nil
	}

	return n.Value
}

func (d *Dequeu) front() int {
	if d.count == 0 || d.root == nil {
		return -1
	}

	return d.root.Value
}

func (d *Dequeu) back() int {
	if d.count == 0 || d.tail == nil {
		return -1
	}

	return d.tail.Value
}

func (d *Dequeu) Size() int {
	return d.count
}

func (d *Dequeu) Empty() int {
	if d.count == 0 {
		return 1
	}
	return 0
}

type Command struct {
	cmd string
	x   int
}

var cmdList []Command = []Command{
	{"front", 0},
	{"back", 0},
	{"pop_front", 0},
	{"pop_back", 0},
	{"push_front", 1},
	{"front", 0},
	{"pop_back", 0},
	{"push_back", 2},
	{"back", 0},
	{"pop_front", 0},
	{"push_front", 10},
	{"push_front", 333},
	{"front", 0},
	{"back", 0},
	{"pop_back", 0},
	{"pop_back", 0},
	{"push_back", 20},
	{"push_back", 1234},
	{"front", 0},
	{"back", 0},
	{"pop_back", 0},
	{"pop_back", 0},
}

func main() {
	d := NewDeque()

	var n int
	fmt.Scanln(&n)

	var rst []int = []int{}
	for i := 0; i < n; i++ {
		var cmd string
		var x int
		fmt.Scanf("%s %d", &cmd, &x)

		if cmd == "push_front" {
			d.PushFront(x)
		} else if cmd == "push_back" {
			d.PushBack(x)
		} else if cmd == "pop_front" {
			n := d.Popfront()
			rst = append(rst, n)
		} else if cmd == "pop_back" {
			n := d.Popback()
			rst = append(rst, n)
		} else if cmd == "size" {
			n := d.Size()
			rst = append(rst, n)
		} else if cmd == "empty" {
			n := d.Empty()
			rst = append(rst, n)
		} else if cmd == "front" {
			n := d.front()
			rst = append(rst, n)
		} else if cmd == "back" {
			n := d.back()
			rst = append(rst, n)
		}

	}

	for _, r := range rst {
		fmt.Println(r)
	}
}

// func main() {
// 	d := NewDeque()

// 	// d.PushFront(1)
// 	// fmt.Println(d.front())
// 	// fmt.Println(d.Popback())

// 	// d.PushBack(2)
// 	// fmt.Println(d.back())
// 	// fmt.Println(d.Popfront())

// 	d.PushFront(10)
// 	d.PushFront(333)
// 	//[333,10]
// 	// fmt.Println(d.front())
// 	// fmt.Println(d.back())
// 	fmt.Println(d.Popback())
// 	fmt.Println(d.Popback())

// 	// fmt.Println(d.Popfront())
// 	// fmt.Println(d.Popfront())
// 	// fmt.Println(d.Popback())

// 	// d.PushFront(3)
// 	// fmt.Println(d.Popback())
// 	// fmt.Println(d.Popfront())
// 	// // fmt.Println(d.Popfront())

// 	// d.PushBack(4)
// 	// fmt.Println(d.front())
// 	// fmt.Println(d.front())

// 	// d.PushBack(4)
// 	// d.PushBack(5)
// 	// d.PushBack(6)

// 	// d.PushFront(1)
// 	// d.PushFront(2)
// 	// d.PushFront(3)

// 	// fmt.Println(d.Popfront())
// 	// fmt.Println(d.Popback())
// 	// fmt.Println(d.front())

// 	// fmt.Println(d.back())
// 	// fmt.Println(d.Size())
// 	// fmt.Println(d.Empty())
// }
