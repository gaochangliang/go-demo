package model

type InitTableFunc interface {
	Init() (err error)
}
