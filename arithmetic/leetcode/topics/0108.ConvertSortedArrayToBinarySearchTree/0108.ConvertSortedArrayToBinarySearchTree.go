package leetcode

import "github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	return &TreeNode{Val: nums[size/2], Left: sortedArrayToBST(nums[:size/2]), Right: sortedArrayToBST(nums[size/2+1:])}
}
