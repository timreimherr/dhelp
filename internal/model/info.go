package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Info struct {
	Id      uint64
	Key     string
	Value   string
	Section *Section
}
