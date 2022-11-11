package pdf

import (
	"bytes"
	"fmt"
	"io"
	"mdpaper/pdf/spec"
)

const (
	dashDbg       string = `├── `
	spacerDbg     string = `│   `
	dashLastDbg   string = `└── `
	spacerLastDbg string = `    `
)

type ChapterNode struct {
	Heading    *Heading
	Parent     *ChapterNode
	ChildNodes []*ChapterNode
}

func (n *ChapterNode) recChildCount() int {
	N := len(n.ChildNodes)
	for _, c := range n.ChildNodes {
		N += c.recChildCount()
	}
	return N
}

func (n *ChapterNode) formatDbg(root bool, prefix string, output io.Writer) {
	if root {
		output.Write([]byte("┌" + n.String() + "\n"))
	}
	for idx, node := range n.ChildNodes {
		lineBuffer := prefix
		childPrefix := prefix
		switch idx {
		case len(n.ChildNodes) - 1:
			lineBuffer += dashLastDbg
			childPrefix += spacerLastDbg

		default:
			lineBuffer += dashDbg
			childPrefix += spacerDbg
		}
		lineBuffer += node.String() + "\n"
		// Write node string representation to output.
		if _, err := output.Write([]byte(lineBuffer)); err != nil {
			panic(err)
		}
		node.formatDbg(false, childPrefix, output)
	}
}

func (n ChapterNode) String() string {
	return fmt.Sprintf("page %d: %s %s", n.Heading.Page, n.Heading.Numbering(), n.Heading.Text.String())
}

type ChapterTree []*ChapterNode

func (tree ChapterTree) Roots() []*ChapterNode {
	roots := make([]*ChapterNode, 0)
	for _, n := range tree {
		if n.Parent == nil {
			roots = append(roots, n)
		}
	}
	return roots
}

func (tree ChapterTree) String() string {
	var buffer bytes.Buffer
	for _, n := range tree.Roots() {
		n.formatDbg(true, "", &buffer)
	}
	return buffer.String()
}

func GenerateChapterTree(headings []*Heading) ChapterTree {
	t := make(ChapterTree, 0)
	for _, h := range headings {
		if h.Level == 1 {
			n := ChapterNode{Heading: h, Parent: nil, ChildNodes: make([]*ChapterNode, 0)}
			t = append(t, &n)
		} else {
			for i := len(t) - 1; i >= 0; i-- {
				if t[i].Heading.Level == h.Level-1 {
					n := ChapterNode{Heading: h, Parent: t[i], ChildNodes: make([]*ChapterNode, 0)}
					t = append(t, &n)
					break
				}
			}
		}
	}
	for _, n := range t {
		if n.Parent != nil {
			n.Parent.ChildNodes = append(n.Parent.ChildNodes, n)
		}
	}
	return t
}

func (tree ChapterTree) GenerateNumbering(root *ChapterNode) {
	for i, r := range root.ChildNodes {
		h := r.Heading
		h.Prefix = root.Heading.Prefix
		h.Prefix[h.Level-1] = i + 1
		tree.GenerateNumbering(r)
	}
}

func (tree ChapterTree) GenerateOutline(outlines *spec.DictionaryObject, pdf *spec.PDF) []*spec.DictionaryObject {
	items := make(map[*ChapterNode]*spec.DictionaryObject)
	roots := tree.Roots()
	for i, n := range roots {
		d := spec.NewDictObject()
		d.Set("Title", "("+n.Heading.Numbering()+" "+n.Heading.String()+")")
		d.Set("Parent", outlines.Reference())
		d.Set("Count", n.recChildCount())
		d.Set("Dest", n.Heading.Destination())
		items[n] = &d
		if i == 0 {
			outlines.Set("First", d.Reference())
		}
		if i == len(tree.Roots())-1 {
			outlines.Set("Last", d.Reference())
		}
	}
	for i, n := range roots {
		if i != 0 {
			items[n].Set("Prev", items[roots[i-1]].Reference())
		}
		if i != len(roots)-1 {
			items[n].Set("Next", items[roots[i+1]].Reference())
		}
	}
	for _, n := range tree {
		n.childrenOutline(&items)
	}
	ret := make([]*spec.DictionaryObject, 0)
	for _, n := range tree {
		if items[n] != nil {
			ret = append(ret, items[n])
		}
	}
	return ret
}

func (n *ChapterNode) childrenOutline(items *map[*ChapterNode]*spec.DictionaryObject) {
	c := n.ChildNodes

	for i, node := range c {
		d := spec.NewDictObject()
		d.Set("Title", "("+node.Heading.Numbering()+" "+node.Heading.String()+")")
		d.Set("Parent", (*items)[n].Reference())
		d.Set("Count", node.recChildCount())
		d.Set("Dest", node.Heading.Destination())
		(*items)[node] = &d
		if i == 0 {
			(*items)[n].Set("First", d.Reference())
		}
		if i == len(c)-1 {
			(*items)[n].Set("Last", d.Reference())
		}
	}

	for i, n := range c {
		if i != 0 {
			(*items)[n].Set("Prev", (*items)[c[i-1]].Reference())
		}
		if i != len(c)-1 {
			(*items)[n].Set("Next", (*items)[c[i+1]].Reference())
		}
	}

	for _, node := range c {
		node.childrenOutline(items)
	}
}
