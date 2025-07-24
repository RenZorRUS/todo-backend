package tests

type TestCase[I any, O any] struct {
	Name     string
	Input    I
	Expected O
}
