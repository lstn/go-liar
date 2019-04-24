package liar

type Truth interface {
	Hash() string
	Lie() string
}