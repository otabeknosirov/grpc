package model

type Order struct {
	Id          string
	Items       []string
	Description string
	Price       uint64
}
