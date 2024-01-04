package magazine

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	Address // встроенная структура реализация наследования из ООП
}

type Employee struct {
	Name string
	Salary float64
	HomeAddress Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}