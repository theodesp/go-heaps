package treap

import (
    goheap "github.com/theodesp/go-heaps"
)

type Node struct {
    Priority    goheap.Integer
    Key         goheap.Item
}

func (a Node) Compare (b goheap.Item) {
    i1 := a
    i2 := b.(Node)
    switch {
        case i1.Key > i2.Key:
            return 1
        case i1.Key < i2.Key:
            return -1
        default:
            return 0
    }
}
