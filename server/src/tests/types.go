package tests

type Test struct {
	Name             string
	Body             map[string]string
	ExpectedResponse interface{}
	ExpectedStatus   int
}
