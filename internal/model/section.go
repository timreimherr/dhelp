package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Section struct {
	Id   uint64
	Name string
}
