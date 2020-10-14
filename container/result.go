package container

// Result is the interface
type Result interface {
	Status() Code
	Data() interface{}
	Error() error
}

// Code is the interface
type Code interface {
	Int() int
}
