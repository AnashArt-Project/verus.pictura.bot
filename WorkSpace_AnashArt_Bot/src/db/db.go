package db

// ---------------- Data Base ----------------

import (
	"fmt"
	"io"
	"strings"
)

// ---------------- Order Blank Struct ----------------
type OrderInfo struct {
	UserName string
	Email    string
	Print    string
	Size     string
	Addres   string
	Payment  string
	Status   string
}

// ---------------- General db of Store ----------------

type Node struct {
	// Использовать ключевое слово по типу "Collection"
	key string 
	value string
	make([]*Node, 0)
}