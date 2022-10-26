package storage_engine


import "fmt"

type BTree struct {

}

func NewBTree() *BTree {
	fmt.Println("Creating new BTree")
	bt := BTree{}
	return &bt
}

