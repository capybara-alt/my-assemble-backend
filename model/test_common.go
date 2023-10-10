package model

type Want[T any] struct {
	Value  T
	ErrMsg string
}
