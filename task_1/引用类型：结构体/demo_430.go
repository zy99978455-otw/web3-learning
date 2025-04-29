package main

/*
`扁平化`多级双向链表
你会得到一个双链表，其中包含的节点有一个下一个指针、一个前一个指针和一个额外的 子指针 。这个子指针可能指向一个单独的双向链表，也包含这些特殊的节点。这些子列表可以有一个或多个自己的子列表，以此类推，以生成如下面的示例所示的 多层数据结构 。
给定链表的头节点 head ，将链表 扁平化 ，以便所有节点都出现在单层双链表中。让 curr 是一个带有子列表的节点。子列表中的节点应该出现在扁平化列表中的 curr 之后 和 curr.next 之前 。
返回 扁平列表的 head 。列表中的节点必须将其 所有 子指针设置为 null 。
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(head *Node) *Node {
	if head == nil {
		return nil
	}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		curr := node
		var last *Node

		for curr != nil {
			next := curr.Next
			// 如果有子节点，递归扁平化
			if curr.Child != nil {
				childLast := dfs(curr.Child)

				// 连接curr和curr.Child
				curr.Next = curr.Child
				curr.Child.Prev = curr

				// 如果 next不为空，将childLast 与next连接
				if next != nil {
					childLast.Next = Next
					Next.Prev = childLast
				}

				// 子链表插入完成，置空child
				curr.Child = nil
				last = childLast
			} else {
				last = curr
			}
			curr = next
		}
		return last
	}

	dfs(head)
	return head
}
