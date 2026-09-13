package learngo

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

func (t *BSTNode) Insert(val int) {
	if t.Value == val {
		t.Value = val
		return
	}

	if val < t.Value {
		if t.Left == nil {
			t.Left = &BSTNode{Value: val}
			return
		}
		if t.Left != nil {
			t.Left.Insert(val)
			return
		}
	}

	if t.Right == nil {
		t.Right = &BSTNode{Value: val}
		return
	} else {
		t.Right.Insert(val)
	}

}

func (t *BSTNode) GetMin() *BSTNode {
	temp := t

	for temp.Left != nil {
		temp = temp.Left
	}

	return temp
}

func (t *BSTNode) GetMax() *BSTNode {
	temp := t

	for temp.Right != nil {
		temp = temp.Right
	}

	return temp
}

func (t *BSTNode) Delete(n BSTNode) *BSTNode {
	if n.Value < t.Value {
		if t.Left != nil {
			t.Left = t.Left.Delete(n)
		}
		return t
	}

	if n.Value > t.Value {
		if t.Right != nil {
			t.Right = t.Right.Delete(n)
		}
		return t
	}

	if t.Left == nil && t.Right == nil {
		return nil
	}

	minLarger := t.Right
	for minLarger.Left != nil {
		minLarger = minLarger.Left
	}
	t.Value = minLarger.Value
	t.Right = t.Right.Delete(*minLarger)

	return t
}
