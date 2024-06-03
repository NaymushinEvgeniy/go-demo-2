package exampleInterface

import "fmt"

type myInterface interface {
	MethodWithoutParams()
	MethodWithParams(float64)
	MethodReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParams() {
	fmt.Println(" This is method without params")
}

func (m MyType) MethodWithParams(expar float64) {
	fmt.Printf("This is method with params %f\n", expar)
}

func (m MyType) MethodReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInInterface() {
	fmt.Println("Method not in interface called")
}
