package translators

//go:generate ../../../bin/enum -type=testEnum -package=translators

type (
	//gen:enum unknownValue unknownTestEnum
	//gen:enum default unknownTestEnum
	testEnum int
)

const (
	//gen:enum string unknownTestEnum
	unknownTestEnum testEnum = iota
	//gen:enum string oneTestEnum
	oneTestEnum
	//gen:enum string twoTestEnum
	twoTestEnum
)
