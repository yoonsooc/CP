package main

import "fmt"

type Node struct {
	leftmost  byte
	rightmost byte
	prefixLen int
	suffixLen int
	longest   int
}

type SegmentTree struct {
	n    int
	tree []Node
}

func (st *SegmentTree) build(node, start, end int, str string) *Node {
	if end-start == 1 {
		st.tree[node] = Node{str[start], str[start], 1, 1, 1}
		return &st.tree[node]
	}

	mid := (start + end) / 2

	leftNode := st.build(node*2, start, mid, str)
	rightNode := st.build(node*2+1, mid, end, str)

	// Set Current Value
	mergable := leftNode.rightmost == rightNode.leftmost
	if mergable {
		mergedStrLen := leftNode.suffixLen + rightNode.prefixLen
		curLongest := max(mergedStrLen, leftNode.longest, rightNode.longest)

		newPrefixLen := leftNode.prefixLen
		newSuffixLen := rightNode.suffixLen
		if newPrefixLen == (mid - start) {
			newPrefixLen += rightNode.prefixLen
		}
		if newSuffixLen == (end - mid) {
			newSuffixLen += leftNode.suffixLen
		}

		st.tree[node] = Node{
			leftNode.leftmost, rightNode.rightmost,
			newPrefixLen, newSuffixLen, curLongest,
		}
	} else {
		st.tree[node] = Node{
			leftNode.leftmost, rightNode.rightmost,
			leftNode.prefixLen, rightNode.suffixLen,
			max(leftNode.longest, rightNode.longest),
		}

	}
	return &st.tree[node]
}

func (st *SegmentTree) Update(index int, value byte) {
	st.update(1, 0, st.n, index, value)
}

func (st *SegmentTree) update(node, start, end, index int, value byte) {
	if end-start == 1 {
		st.tree[node].leftmost = value
		st.tree[node].rightmost = value
		return
	}

	mid := (start + end) / 2
	if index < mid {
		st.update(node*2, start, mid, index, value)
	} else {
		st.update(node*2+1, mid, end, index, value)
	}

	// Update Current Value
	leftNode := st.tree[node*2]
	rightNode := st.tree[node*2+1]
	mergable := leftNode.rightmost == rightNode.leftmost
	if mergable {
		mergedStrLen := leftNode.suffixLen + rightNode.prefixLen
		curLongest := max(mergedStrLen, leftNode.longest, rightNode.longest)

		newPrefixLen := leftNode.prefixLen
		newSuffixLen := rightNode.suffixLen
		if newPrefixLen == (mid - start) {
			newPrefixLen += rightNode.prefixLen
		}
		if newSuffixLen == (end - mid) {
			newSuffixLen += leftNode.suffixLen
		}

		st.tree[node] = Node{
			leftNode.leftmost, rightNode.rightmost,
			newPrefixLen, newSuffixLen, curLongest,
		}
	} else {
		st.tree[node] = Node{
			leftNode.leftmost, rightNode.rightmost,
			leftNode.prefixLen, rightNode.suffixLen,
			max(leftNode.longest, rightNode.longest),
		}

	}
}

func (st *SegmentTree) Query(left, right int) {
	st.query(1, 0, st.n, left, right)
}

func (st *SegmentTree) query(node, start, end, left, right int) int {
	if right <= start || end <= left {
		return 0
	}

	if left <= start && end <= right {
		return st.tree[node].longest
	}

	mid := (start + end) / 2
	return max(st.query(node*2, start, mid, left, right) + st.query(node*2+1, mid, end, left, right))
}

func newSegmentTree(str string) *SegmentTree {
	st := &SegmentTree{
		n:    len(str),
		tree: make([]Node, 4*len(str)),
	}

	if len(str) > 0 {
		st.build(1, 0, len(str), str)
	}

	return st
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	segTree := newSegmentTree(s)
	result := make([]int, len(queryIndices))
	// fmt.Println(segTree)
	for i, ci := range queryIndices {
		c := queryCharacters[i]
		segTree.update(1, 0, len(s), ci, c)
		result[i] = segTree.tree[1].longest
	}
	return result
}

func main() {
	fmt.Println(longestRepeating("babacc", "bcb", []int{1, 3, 3}))
	fmt.Println(longestRepeating("abyzz", "aa", []int{2, 1}))
}
