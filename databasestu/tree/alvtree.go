package main

import (
    "fmt"
    "math"
)

// TreeNode 表示 AVL 树的节点。
type TreeNode struct {
    Key    int
    Left   *TreeNode
    Right  *TreeNode
    Height int
}

// AVLTree 表示 AVL 树的根节点。
type AVLTree struct {
    Root *TreeNode
}

// NewAVLTree 创建一个新的 AVL 树。
func NewAVLTree() *AVLTree {
    return &AVLTree{}
}

// Insert 向 AVL 树中插入一个节点。
func (t *AVLTree) Insert(key int) {
    t.Root = t.insertRec(t.Root, key)
}

func (t *AVLTree) insertRec(node *TreeNode, key int) *TreeNode {
    if node == nil {
        return &TreeNode{Key: key, Height: 1}
    }

    if key < node.Key {
        node.Left = t.insertRec(node.Left, key)
    } else if key > node.Key {
        node.Right = t.insertRec(node.Right, key)
    } else {
        return node // 重复的键不允许
    }

    node.Height = 1 + int(math.Max(float64(t.getHeight(node.Left)), float64(t.getHeight(node.Right))))

    // 获取平衡因子
    balance := t.getBalance(node)

    // 左子树高度大于右子树
    if balance > 1 && key < node.Left.Key {
        return t.rotateRight(node)
    }

    // 右子树高度大于左子树
    if balance < -1 && key > node.Right.Key {
        return t.rotateLeft(node)
    }

    // 左子树高度大于右子树，需要左右旋转
    if balance > 1 && key > node.Left.Key {
        node.Left = t.rotateLeft(node.Left)
        return t.rotateRight(node)
    }

    // 右子树高度大于左子树，需要右左旋转
    if balance < -1 && key < node.Right.Key {
        node.Right = t.rotateRight(node.Right)
        return t.rotateLeft(node)
    }

    return node
}

func (t *AVLTree) getHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return node.Height
}

func (t *AVLTree) getBalance(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return t.getHeight(node.Left) - t.getHeight(node.Right)
}

func (t *AVLTree) rotateLeft(z *TreeNode) *TreeNode {
    y := z.Right
    T2 := y.Left

    y.Left = z
    z.Right = T2

    z.Height = 1 + int(math.Max(float64(t.getHeight(z.Left)), float64(t.getHeight(z.Right))))
    y.Height = 1 + int(math.Max(float64(t.getHeight(y.Left)), float64(t.getHeight(y.Right))))

	    return y
}

func (t *AVLTree) rotateRight(y *TreeNode) *TreeNode {
    x := y.Left
    T2 := x.Right

    x.Right = y
    y.Left = T2

    y.Height = 1 + int(math.Max(float64(t.getHeight(y.Left)), float64(t.getHeight(y.Right))))
    x.Height = 1 + int(math.Max(float64(t.getHeight(x.Left)), float64(t.getHeight(x.Right))))

    return x
}

// InOrderTraversal 执行 AVL 树的中序遍历，返回升序排列的节点值切片。
func (t *AVLTree) InOrderTraversal() []int {
    var result []int
    t.inOrderTraversal(t.Root, &result)
    return result
}

func (t *AVLTree) inOrderTraversal(node *TreeNode, result *[]int) {
    if node != nil {
        t.inOrderTraversal(node.Left, result)
        *result = append(*result, node.Key)
        t.inOrderTraversal(node.Right, result)
    }
}

func main() {
    tree := NewAVLTree()
    keys := []int{9, 5, 10, 0, 6, 11, -1, 1, 2}

    for _, key := range keys {
        tree.Insert(key)
    }

