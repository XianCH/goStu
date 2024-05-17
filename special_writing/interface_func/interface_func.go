package main

type sum interface {
	add(int, int) (int, error)
}

type addTest struct{}

func (at *addTest) add(a int, b int) (int, error) {
	return a + b, nil
}

type sumFunc func(int, int) (int, error)

func (sf sumFunc) add(a int, b int) (int, error) {
	return sf(a, b)
}
