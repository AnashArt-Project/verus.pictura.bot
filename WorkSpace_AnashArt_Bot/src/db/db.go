package db

import (
	"fmt"
	"io"
	"strings"
)

type OrderInfo struct {
	UserName string
	Email    string
	Print    string
	Size     string
	Addres   string
	Payment  string
	Status   string
}

// ---------------- Data Base ----------------

type Tree map[string]Tree

func (tree Tree) Add(path string) {
	frags := strings.Split(path, "/")
	tree.add(frags)
}

func (tree Tree) add(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		nextTree = Tree{}
		tree[frags[0]] = nextTree
	}

	nextTree.add(frags[1:])
}

// Идея в отм чтобы последний элемент массива удалить не трогая остальные
// Если прировнять к nil то удаляет все дерево
func (tree Tree) Remove(path string) {
	frags := strings.Split(path, "/")
	tree.remove(frags)
}

func (tree Tree) remove(frags []string) {
	if len(frags) == 0 {
		return
	}
	tempTree, ok := tree[frags[0]]
	if len(frags) == 1 {
		if ok {
			tempTree = Tree{}
			tree[frags[0]] = tempTree
		}
	}
	tempTree.remove(frags[1:])
}

// ВЫВОД В КОНСОЛЬ
func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

// ВЫВОД В ТЕЛЕГРАМ
func (tree Tree) TreePrint(root bool, padding string, msg string) string {
	if tree == nil {
		return "I DON'T WANT PIECE! I WANT PROBLEMS, ALWAYS!"
	}
	index := 0
	for k, v := range tree {
		msg += padding + getPadding(root, getBoxType(index, len(tree))) + k + "\n"
		msg = v.TreePrint(false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))), msg)
		index++
	}
	return msg
}

type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) BoxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) BoxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
