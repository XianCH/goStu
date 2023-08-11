package tree

import "fmt"

// 定义二叉树结点
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 插入结点
func insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }

    if val < root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }

    return root
}

// 搜索结点
func search(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }

    if val < root.Val {
        return search(root.Left, val)
    }

    return search(root.Right, val)
}

// 中序遍历打印二叉树
func inorderTraversal(root *TreeNode) {
    if root != nil {
        inorderTraversal(root.Left)
        fmt.Printf("%d ", root.Val)
        inorderTraversal(root.Right)
    }
}

func main() {
    var root *TreeNode
    elements := []int{5, 3, 7, 2, 4, 6, 8}

    for _, val := range elements {
        root = insert(root, val)
    }

    fmt.Println("中序遍历结果：")
    inorderTraversal(root)

    // 搜索结点测试
    searchVal := 4
    result := search(root, searchVal)
    if result != nil {
        fmt.Printf("\n找到了值为 %d 的结点\n", searchVal)
    } else {
        fmt.Printf("\n没有找到值为 %d 的结点\n", searchVal)
    }
}