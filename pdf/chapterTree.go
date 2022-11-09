package pdf

import (
	"github.com/jmichiels/tree"
)

type ChapterNode struct {
	Heading *Heading
	Parent  *ChapterNode
}

func (c ChapterNode) String() string {
	return c.Heading.String()
}

type ChapterTree []ChapterNode

func (c ChapterTree) RootNodes() []tree.Node {
	return c.ChildrenNodes(nil)
}

func (c ChapterTree) ChildrenNodes(parent tree.Node) (nodes []tree.Node) {
	if parent == nil {
		for _, n := range c {
			if n.Parent == nil {
				nodes = append(nodes, n)
			}
		}
		return
	} else {
		for _, n := range c[1:] {
			if n.Parent.Heading == parent.(ChapterNode).Heading {
				nodes = append(nodes, n)
			}
		}
	}
	return
}

func GenerateChapterTree(headings []*Heading) ChapterTree {
	t := make(ChapterTree, 0)
	for _, h := range headings {
		if h.Level == 1 {
			t = append(t, ChapterNode{Heading: h})
		} else {
			for i := len(t) - 1; i >= 0; i-- {
				if t[i].Heading.Level == h.Level-1 {
					t = append(t, ChapterNode{Heading: h, Parent: &t[i]})
					break
				}
			}
		}
	}
	return t
}

func (c ChapterTree) GenerateNumbering(parent ChapterNode) {
	for i, r := range c.ChildrenNodes(parent) {
		h := r.(ChapterNode).Heading
		h.Prefix = parent.Heading.Prefix
		h.Prefix[h.Level-1] = i + 1
		c.GenerateNumbering(r.(ChapterNode))
	}
}

func (c ChapterTree) String() string {
	return tree.String(c)
}
