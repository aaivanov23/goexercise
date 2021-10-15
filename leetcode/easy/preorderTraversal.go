package main

/*import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
    tree := &TreeNode{
        Val: 1,
        Left: nil,
        Right: &TreeNode{
            Val:2,
            Left:&TreeNode{
                Val:3,
                Left: nil,
                Right:nil},
            Right: nil}}

    fmt.Println(preorderTraversal(tree))
}*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    result := []int{}

    if root == nil {
        return result
    }

    result = append(result, root.Val)
    result = append(result, preorderTraversal(root.Left)...)
    result = append(result, preorderTraversal(root.Right)...)

    return result
}
