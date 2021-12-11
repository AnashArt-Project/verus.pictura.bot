package db

type OrderInfo struct {
	UserName string
	Email    string
	Print    string
	Size     string
	Addres   string
	Payment  string
	Status   string
}

// Data base
// 		название_принта_1 string
// 			размер_1 string
// 				цвет_1 string
// 					количество в налиции int64
// 				цвет_2 string
// 					количество в налиции int64

type Products struct {
	PrintName string
	Size      SizeProduct
}

type SizeProduct struct {
	Size  string
	Color ColorProduct
}

type ColorProduct struct {
	Color string
	Value int64
}
