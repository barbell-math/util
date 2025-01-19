package testenum

//go:generate ../../../bin/enum -type=TestEnum -package=testenum

type (
	//gen:enum unknownValue UnknownAppAction
	//gen:enum default UnknownAppAction
	TestEnum int
)

const (
	//gen:enum string UnknownAppAction
	UnknownAppAction TestEnum = iota
	//gen:enum string AppActionOne
	AppActionOne
	//gen:enum string AppActionTwo
	AppActionTwo
)
