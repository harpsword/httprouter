package httprouter

import (
	"net/http"
	"strings"
)

type trieTreeNodeType uint8

const (
	root trieTreeNodeType = iota
	normal
	param
	catchAll
)

type trieTreeNode struct {
	nType trieTreeNodeType
	path string
	iswild bool
	handle http.Handler
	children []*trieTreeNode
}

func (node *trieTreeNode) getHandler(paths []string) http.Handler {
	for _, nextNode := range node.children {
		if nextNode.path == paths[0] {
			if len(paths) == 1 {
				return nextNode.handle
			}else {
				return nextNode.getHandler(paths[1:])
			}
		}
	}
	panic("this route has not been registered!")
}

func (node *trieTreeNode) addRoute(paths []string, handler http.Handler){
	found := false
	var newNode *trieTreeNode
	for _, nextNode := range node.children {
		if nextNode.path == paths[0] {
			found = true
			newNode = nextNode
		}
	}
	if !found {
		newNode = &trieTreeNode{
			nType: normal,
			path: paths[0],
			iswild: false,
			handle: nil,
			children: []*trieTreeNode{},
		}
	}

	node.children = append(node.children, newNode)
	if len(paths) > 1 {
		newNode.addRoute(paths[1:], handler)
	}else {
		// last node is newNode
		newNode.handle = handler
	}
}

type trieTree struct {
	root map[string]*trieTreeNode
}

func (tree *trieTree) getHandelr(path string, method string) http.Handler {
	if path[0:1] == "/" {
		path = path[1:]
	}
	paths := strings.Split(path, "/")
	if node, ok := tree.root[method]; ok {
		return node.getHandler(paths)
	} else {
		panic("this route has not been registered!")
	}

}

func (tree *trieTree) addRoute(path string, method string, handler http.Handler) {
	if path[0:1] == "/" {
		path = path[1:]
	}
	paths := strings.Split(path, "/")
	node, ok := tree.root[method]

	if !ok {
		node = &trieTreeNode{
			nType: root,
			path: "/",
			iswild: false,
			handle: nil,
			children: []*trieTreeNode{},
		}
		tree.root[method] = node
	}
	node.addRoute(paths, handler)
}


