package prototype

type Cloneable[T any] interface {
	Clone() T
}
