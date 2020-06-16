package httprouter

import (
	"net/http"
	"testing"
)

type f struct {
	str string
}

var fakeHandlerValue string

func (ff f) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fakeHandlerValue = ff.str
}


func TestNode_addHandler(t *testing.T){
	var node = &trieTreeNode{
		nType: root,
		path: "/",
		iswild: false,
		handle: nil,
		children: []*trieTreeNode{},
	}
	paths := []string{"path1", "path2"}
	fExample := f{str: "hello world"}
	node.addRoute(paths, fExample)

	if len(node.children) != 1 {
		t.Errorf("the number of root 's children should be 1")
	}
	if len(node.children[0].children) != 1 {
		t.Errorf("the number of root's children's child should be 1")
	}
	if node.children[0].path != paths[0] {
		t.Errorf("error path for root's child")
	}
	if node.children[0].children[0].path != paths[1] {
		t.Errorf("error : path of root's child's child ")
	}
}

func TestNode_getHandler(t *testing.T){
	var node = &trieTreeNode{
		nType: root,
		path: "/",
		iswild: false,
		handle: nil,
		children: []*trieTreeNode{},
	}
	paths := []string{"path1", "path2"}
	fExample := f{str: "hello world"}
	node.addRoute(paths, fExample)

	tmpf := node.getHandler(paths)
	tmpf.ServeHTTP(nil, nil)
	if fakeHandlerValue != fExample.str {
		t.Errorf("wrong handler ")
	}
}

func TestTree_addRoute_getHandler(t *testing.T){
	var tree = &trieTree{
		root: map[string]*trieTreeNode{},
	}
	fExample := f{str: "hello world"}
	tree.addRoute("/path1/path2", "GET", fExample)

	tmpf := tree.getHandelr("/path1/path2", "GET")
	tmpf.ServeHTTP(nil, nil)
	if fakeHandlerValue != fExample.str {
		t.Errorf("wrong handler ")
	}
}

