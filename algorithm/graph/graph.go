package main

import (
	"fmt"
	"sort"
)

// TColor 节点着色
type TColor uint8

// 用于遍历时着色
const (
	White TColor = 0
	Gray  TColor = 1
	Black TColor = 2
)

// Vertex 一张图
type Vertex struct {
	Data     interface{}
	Color    TColor
	Prev     *Vertex // 前驱
	Discover int
	Done     int
	Edges    []*Vertex
	Visited  []*Vertex
}

func makeKey(data interface{}) string {
	if data == nil {
		return ""
	}
	return fmt.Sprint(data)
}

// DirectedGraph 通过邻接链表表示一个有向图
type DirectedGraph struct {
	VertexMap map[string]*Vertex
	time      int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		VertexMap: make(map[string]*Vertex),
	}
}

func (g *DirectedGraph) AddEdge(from, to interface{}) {
	keyFrom := makeKey(from)
	vFrom, ok := g.VertexMap[keyFrom]
	if !ok {
		vFrom = &Vertex{
			Data: from,
		}
		g.VertexMap[keyFrom] = vFrom
	}
	if to != nil {
		keyTo := makeKey(to)
		vTo, ok := g.VertexMap[keyTo]
		if !ok {
			vTo = &Vertex{
				Data: to,
			}
			g.VertexMap[keyTo] = vTo
		}

		vFrom.Edges = append(vFrom.Edges, vTo)
	}
}

func (g *DirectedGraph) DFS() {
	g.time = 0
	for _, v := range g.VertexMap {
		if v.Color == White {
			g.DFSVisit(v)
		}

		fmt.Println(v.Data, "Discover=", v.Discover, "Done=", v.Done)
		for _, vv := range v.Visited {
			fmt.Println("#", v.Data, "Visited ", vv.Data)
		}
	}
}

func (g *DirectedGraph) DFSVisit(u *Vertex) {
	g.time++
	u.Discover = g.time
	u.Color = Gray

	// Record visited
	first := u
	for first.Prev != nil {
		first = first.Prev
	}
	first.Visited = append(first.Visited, u)

	// Visit all adjectives
	for _, v := range u.Edges {
		if v.Color == White {
			v.Prev = u
			g.DFSVisit(v)
		}
	}
	u.Color = Black
	g.time++
	u.Done = g.time
}

func (g *DirectedGraph) Transform() DirectedGraph {
	t := DirectedGraph{
		VertexMap: make(map[string]*Vertex),
	}

	// Copy values
	for k, v := range g.VertexMap {
		dst := &Vertex{
			Data:     v.Data,
			Color:    White,
			Discover: v.Discover,
			Done:     v.Done,
		}
		t.VertexMap[k] = dst
	}
	// transform edges
	for k, v := range g.VertexMap {
		for _, edgeDst := range v.Edges {
			src := t.VertexMap[makeKey(edgeDst.Data)]
			src.Edges = append(src.Edges, t.VertexMap[k])
		}
	}
	return t
}

// SCC 获取不同点构成的强连通分量
func (g *DirectedGraph) SCC() {
	g.DFS()

	t := g.Transform()

	var vertexSlice []*Vertex
	for _, v := range t.VertexMap {
		vertexSlice = append(vertexSlice, v)
	}
	// Resort
	sort.Slice(vertexSlice, func(i, j int) bool {
		return vertexSlice[i].Done > vertexSlice[j].Done
	})

	t.time = 0
	for _, v := range vertexSlice {
		if v.Color == White {
			t.DFSVisit(v)
			fmt.Println("SCC Group------------")
			for _, vv := range v.Visited {
				fmt.Print(vv.Data, " ")
			}
			fmt.Println()
		}
	}
}

func main() {
	g := NewDirectedGraph()

	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("b", "e")
	g.AddEdge("b", "f")
	g.AddEdge("c", "d")
	g.AddEdge("c", "g")
	g.AddEdge("d", "c")
	g.AddEdge("d", "h")
	g.AddEdge("e", "a")
	g.AddEdge("e", "f")
	g.AddEdge("f", "g")
	g.AddEdge("g", "f")
	g.AddEdge("g", "h")
	g.AddEdge("h", "h")
	g.AddEdge("i", nil) // 孤立节点
	g.AddEdge("j", "k") // 普通边
	g.AddEdge("m", "n")
	g.AddEdge("n", "o")
	g.AddEdge("o", "p")

	for k, v := range g.VertexMap {
		fmt.Println("Node:", k)
		for _, edge := range v.Edges {
			fmt.Print(edge.Data, " ")
		}
		fmt.Println()
	}

	g.SCC()
}
