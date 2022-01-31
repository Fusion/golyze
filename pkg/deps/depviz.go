package deps

import (
	"log"

	"github.com/KyleBanks/depth"
	"github.com/fusion/golyz/pkg/wrap"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func BuildDepsData(l *log.Logger) wrap.SankeyData {
	var t depth.Tree
	err := t.Resolve(".")
	if err != nil {
		l.Fatal(err)
	}
	var nodes []opts.SankeyNode
	var links []opts.SankeyLink
	visited := make(map[string]bool)

	rootNode := opts.SankeyNode{Name: t.Root.Name}
	nodes = append(nodes, rootNode)
	nodes, links = writePkg(l, *t.Root, rootNode, nodes, links, visited)

	return wrap.SankeyData{Nodes: nodes, Links: links}
}

func writePkg(l *log.Logger, p depth.Pkg, parentNode opts.SankeyNode, nodes []opts.SankeyNode, links []opts.SankeyLink, visited map[string]bool) ([]opts.SankeyNode, []opts.SankeyLink) {
	for _, pkg := range p.Deps {
		curNode := opts.SankeyNode{Name: pkg.Name}
		_, ok := visited[curNode.Name]
		if !ok {
			nodes = append(nodes, curNode)
			visited[curNode.Name] = true
		}
		links = append(links, opts.SankeyLink{Source: parentNode.Name, Target: curNode.Name, Value: 1})
		nodes, links = writePkg(l, pkg, curNode, nodes, links, visited)
	}
	return nodes, links
}
