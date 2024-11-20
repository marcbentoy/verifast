package component

type Component interface {
	Inputs() []Pin
	Outputs() []Pin
}
