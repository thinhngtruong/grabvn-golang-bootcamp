package week1

type expression interface {
	parse(expString string) error
	calculate() (float64, error)
}