package helpers

// DoublyLinkedList maintains access to tail node
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func (dll *DoublyLinkedList) Push(data interface{}) *Node {
	node := &Node{data, nil, nil}
	if dll.head != nil {
		node.next = dll.head
		dll.head.prev = node
	}

	if dll.tail == nil {
		dll.tail = node
	}

	dll.head = node

	return node
}

func (dll *DoublyLinkedList) PushNode(node *Node) {
	if dll.head == node {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if dll.tail == node {
		dll.tail = node.prev
	}

	node.prev = nil
	node.next = dll.head
	if node.next != nil {
		node.next.prev = node
	}

	dll.head = node
}

func (dll *DoublyLinkedList) Tail() *Node {
	return dll.tail
}

func (dll *DoublyLinkedList) Delete(node *Node) {
	if dll.head == node {
		dll.head = node.next
	}

	if dll.tail == node {
		dll.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	node = nil
}

// MaxHeap
// [][0] row idx
// [][1] col idx
// [][2] effort needed to move to this cell
type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	col := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return col
}
