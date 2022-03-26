package db

import "fmt"

// ---------------- Data Base ----------------

// ---------------- Order Blank Struct ----------------
type OrderInfo struct {
	UserName string
	Email    string
	Print    string
	Size     string
	Phone    string
	City     string
	Street   string
	House    string
	Payment  string
	Status   string
}

func ToStringOrderInfo(order *OrderInfo) string {
	str := fmt.Sprintf("Email: %s\nPrint: %s\nSize: %s\nPhone: %s\nCity: %s\nStreet: %s\nHouse: %s\n", order.Email, order.Print, order.Size, order.Phone, order.City, order.Street, order.House)
	return str
}

func ToStringAllOrderInfo(order *OrderInfo) string {
	str := fmt.Sprintf("UserName: %s\nEmail: %s\nPrint: %s\nSize: %s\nPhone: %s\nCity: %s\nStreet: %s\nHouse: %s\nPayment: %s\nStatus: %s", order.UserName, order.Email, order.Print, order.Size, order.Phone, order.City, order.Street, order.House, order.Payment, order.Status)
	return str
}

// ---------------- General db of Store ----------------

type Node struct {
	key      string
	value    string
	children []*Node
}

func findById(root *Node, key string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.key == key {
			return nextUp
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

// func (node *Node) remove() {
// 	// Remove the node from it's parents children collection
// 	for idx, sibling := range n.parent.children {
// 		if sibling == node {
// 			node.parent.children = append(
// 				node.parent.children[:idx],
// 				node.parent.children[idx+1:]...,
// 			)
// 		}
// 	}

// 	// If the node has any children, set their parent to nil and set the node's children collection to nil
// 	if len(node.children) != 0 {
// 		for _, child := range node.children {
// 			child.parent = nil
// 		}
// 		node.children = nil
// 	}
// }

// https://stackoverflow.com/questions/16877427/how-to-implement-a-non-binary-tree

// https://ieftimov.com/posts/golang-datastructures-trees/#:~:text=A%20tree%20is%20a%20data,nodes%20that%20form%20a%20hierarchy.
