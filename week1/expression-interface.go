package week1

type expression interface {
	parse(expressionString string) error
	calculate() (float64, error)
}