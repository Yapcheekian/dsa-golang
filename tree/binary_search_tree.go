package tree

type node struct {
	key, value  string
	left, right *node
}

type BST struct {
	root *node
}

func NewBST() *BST {
	return &BST{}
}

func (t *BST) Get(key string) string {
	node := t.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return node.value
		}
	}
	return ""
}

func (t *BST) Set(key, value string) {
	t.root = put(t.root, key, value)
}

func put(n *node, key, value string) *node {
	if n == nil {
		return &node{key: key, value: value}
	}
	if key < n.key {
		n.left = put(n.left, key, value)
	} else if key > n.key {
		n.right = put(n.right, key, value)
	} else {
		n.value = value
	}
	return n
}