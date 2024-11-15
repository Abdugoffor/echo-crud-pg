package repository

type I64 = ID[int64] //@name I64

type ID[T any] struct {
	ID T `json:"id"`
} //@name ID

func NewID[T any](id T) ID[T] {
	return ID[T]{
		ID: id,
	}
}

type MessageString = Message[string] //@name MessageString

type Message[T any] struct {
	Message T `json:"message"`
} //@name Message

func NewMessage[T any](message T) Message[T] {
	return Message[T]{
		Message: message,
	}
}
