/*
@Time : 2021/8/17 下午8:01
@Author : Mrxuexi
@File : trie
@Software: GoLand
*/
package Ez

import "strings"

type node struct {
	path  string 	/* 需要匹配的整体路由 */
	part     string 	/* 路由中的一部分，例如 :lang */
	children []*node 	/* 存储子节点们 */
	isBlurry   bool 	/* 如果模糊匹配则为true */
}

// 遍历children找到我们想要的目标节点（首个）
func (n *node) matchChild(part string) *node {
	//遍历子节点们，对比子节点的part和part是否相同，是或者遍历到的子节点支持模糊匹配则返回该子节点
	for _, child := range n.children {
		if child.part == part || child.isBlurry {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	//遍历选择满足条件的子节点，加入到nodes中，然后返回
	for _, child := range n.children {
		if child.part == part || child.isBlurry {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//插入方法，用一个递归实现，找匹配的路径直到找不到匹配当前part的节点，新建
func (n *node) insert(path string, parts []string, height int)  {
	//如果遍历到底部了，则将我们的path存入节点，开始返回。递归的归来条件。
	if len(parts) == height{
		n.path = path
		return
	}
	//获取这一节的part，并进行搜索
	part := parts[height]
	child := n.matchChild(part)

	//若没有搜索到匹配的子节点，则根据目前的part构造一个子节点
	if child == nil {
		 child = &node{
		 	part: part,
		 	isBlurry: part[0] == ':' || part[0] == '*',
		 }
		 n.children = append(n.children, child)
	}
	child.insert(path, parts, height+1)
}

//搜索方法
func (n *node) search(parts []string, height int) *node {
	//如果节点到头，或者存在*前缀的节点，开始返回
	if len(parts) == height || strings.HasPrefix(n.part,"*") {
		//如果此时遍历到的n没有存储路径，说明未到最底层，则返回空
		if n.path == "" {
			return nil
		}
		return n
	}
	//搜索找到满足part的子节点们放入children
	part := parts[height]
	children := n.matchChildren(part)
	//接着遍历子节点们，递归调用获得下一级的子节点们，要走到头的同时，找到了对应的节点，才返回最终我们找到的result
	//这里为什么要遍历子节点们进行深入搜索，因为它还存在满足isBlurry:true的节点，我们也需要在其中深入搜索。
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			//返回满足要求的节点
			return result
		}
	}
	return nil
}