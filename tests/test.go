package tests

type SearchArgs struct {
	Key   string
	Value string
}

type TestCase struct {
	Name           string
	Args           interface{}
	ExpectedResult interface{}
	ExpectedError  bool
}

type SearchTestCase struct {
	Name           string
	Args           SearchArgs
	ExpectedResult interface{}
	ExpectedError  bool
}
