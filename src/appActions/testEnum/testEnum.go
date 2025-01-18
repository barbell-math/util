package testenum

//go:generate ../../../bin/enum -type=TestEnum -package=testenum

type (
	//gen:enum unknownValue UnknownTestEnum
	//gen:enum default UnknownTestEnum
	TestEnum int
)

const (
	//gen:enum string unknownTestEnum
	UnknownTestEnum TestEnum = iota
	//gen:enum string AppActionOne
	AppActionOne
	//gen:enum string AppActionTwo
	AppActionTwo
)
