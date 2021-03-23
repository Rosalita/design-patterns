package main

import "fmt"

type folder struct {
    children []node
    name      string
}

func (f *folder) print(indentation string) {
    fmt.Println(indentation + f.name)
    for _, child := range f.children {
        child.print(indentation + indentation)
    }
}

func (f *folder) clone() node {
    cloneFolder := &folder{name: f.name + "_clone"}
    var tempChildren []node
    for _, child := range f.children {
        copy := child.clone()
        tempChildren = append(tempChildren, copy)
    }
    cloneFolder.children = tempChildren
    return cloneFolder
}