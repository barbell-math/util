package argparse

import (
	"testing"

	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/io/argparse/argTypes"
)

func TestAddArgAllDefaultOptions(t *testing.T) {
	p:=NewParser("test","testing")
	res:=AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{})
	test.BasicTest((*string)(nil),res,
		"The parser threw an error when it should not have",t,
	)
	test.BasicTest(1,len(p.longNameArgs),
		"The parser did not add an arg to the long names map when it should have.",t,
	);
	test.BasicTest(0,len(p.shortNameArgs),
		"The parser added an arg to the short names map when it shouldn't have.",t,
	);
	test.BasicTest(false,p.longNameArgs["test"].Present,
		"The parser did not correctly set present to false by default.",t,
	)
	test.BasicTest(false,p.longNameArgs["test"].Required,
		"The parser did not correctly set required to false by default.",t,
	)
	test.BasicTest("",p.longNameArgs["test"].Description,
		"The parser did not correctly set the description to an empty string by default.",t,
	)
}

func TestAddArgNoDefaultOptions(t *testing.T){
	p:=NewParser("test","testing")
	res:=AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{
		ShortName: 't',
		Required: true,
		Description: "Test description",
		Default: "defaultString",
	})
	test.BasicTest((*string)(nil),res,
		"The parser threw an error when it should not have",t,
	)
	test.BasicTest(1,len(p.longNameArgs),
		"The parser did not add an arg to the long names map when it should have.",t,
	);
	test.BasicTest(1,len(p.shortNameArgs),
		"The parser added an arg to the short names map when it shouldn't have.",t,
	);
	test.BasicTest(false,p.longNameArgs["test"].Present,
		"The parser did not correctly set present to false by default.",t,
	)
	test.BasicTest(true,p.longNameArgs["test"].Required,
		"The parser did not correctly set required to false by default.",t,
	)
	test.BasicTest("Test description",p.longNameArgs["test"].Description,
		"The parser did not correctly set the description to an empty string by default.",t,
	)
	test.BasicTest(p.longNameArgs["test"],p.shortNameArgs['t'],
		"The parser did not correctly set the short name map.",t,
	)
}

func TestArgWithSameTranslatedType(t *testing.T) {
	p:=NewParser("test","testing")
	res:=AddArg[argTypes.IpV4Addr,argTypes.IpV4Addr](
		p,"test",
		&Options[argTypes.IpV4Addr,argTypes.IpV4Addr]{
			ShortName: 't',
			Required: true,
			Description: "Test description",
			Default: argTypes.LocalHost,
		},
	)
	test.BasicTest((*argTypes.IpV4Addr)(nil),res,
		"The parser threw an error when it should not have",t,
	)
	test.BasicTest(1,len(p.longNameArgs),
		"The parser did not add an arg to the long names map when it should have.",t,
	);
	test.BasicTest(1,len(p.shortNameArgs),
		"The parser added an arg to the short names map when it shouldn't have.",t,
	);
	test.BasicTest(false,p.longNameArgs["test"].Present,
		"The parser did not correctly set present to false by default.",t,
	)
	test.BasicTest(true,p.longNameArgs["test"].Required,
		"The parser did not correctly set required to false by default.",t,
	)
	test.BasicTest("Test description",p.longNameArgs["test"].Description,
		"The parser did not correctly set the description to an empty string by default.",t,
	)
	test.BasicTest(p.longNameArgs["test"],p.shortNameArgs['t'],
		"The parser did not correctly set the short name map.",t,
	)
}

func TestMultipleArgsDifferingTypes(t *testing.T) {
	p:=NewParser("test","testing")
	res:=AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{
		ShortName: 't',
		Required: true,
		Description: "Test description",
		Default: "default string",
	})
	res2:=AddArg[argTypes.IpV4Addr,argTypes.IpV4Addr](
		p,"Test",
		&Options[argTypes.IpV4Addr,argTypes.IpV4Addr]{
			ShortName: 'T',
			Required: true,
			Description: "Test description",
			Default: argTypes.LocalHost,
		},
	)
	test.BasicTest((*string)(nil),res,
		"The parser threw an error when it should not have",t,
	)
	test.BasicTest((*argTypes.IpV4Addr)(nil),res2,
		"The parser threw an error when it should not have",t,
	)
	test.BasicTest(2,len(p.longNameArgs),
		"The parser did not add an arg to the long names map when it should have.",t,
	);
	test.BasicTest(2,len(p.shortNameArgs),
		"The parser added an arg to the short names map when it shouldn't have.",t,
	);
	test.BasicTest(false,p.longNameArgs["test"].Present,
		"The parser did not correctly set present to false by default.",t,
	)
	test.BasicTest(false,p.longNameArgs["Test"].Present,
		"The parser did not correctly set present to false by default.",t,
	)
	test.BasicTest(true,p.longNameArgs["test"].Required,
		"The parser did not correctly set required to false by default.",t,
	)
	test.BasicTest(true,p.longNameArgs["Test"].Required,
		"The parser did not correctly set required to false by default.",t,
	)
	test.BasicTest("Test description",p.longNameArgs["test"].Description,
		"The parser did not correctly set the description to an empty string by default.",t,
	)
	test.BasicTest("Test description",p.longNameArgs["Test"].Description,
		"The parser did not correctly set the description to an empty string by default.",t,
	)
	test.BasicTest(p.longNameArgs["test"],p.shortNameArgs['t'],
		"The parser did not correctly set the short name map.",t,
	)
	test.BasicTest(p.longNameArgs["Test"],p.shortNameArgs['T'],
		"The parser did not correctly set the short name map.",t,
	)
}

func TestMultipleArgsSameLongName(t *testing.T) {
	p:=NewParser("test","testing")
	test.Panics(
		func() {
			AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{
				Required: true,
				Description: "Test description",
				Default: "default string",
			})
			AddArg[argTypes.IpV4Addr,argTypes.IpV4Addr](
				p,"test",
				&Options[argTypes.IpV4Addr,argTypes.IpV4Addr]{
					Required: true,
					Description: "Test description",
					Default: argTypes.LocalHost,
				},
			)
		}, "The parser did not panic with duplicate long names.",t,
	);
}

func TestMultipleArgsSameShortName(t *testing.T) {
	p:=NewParser("test","testing")
	test.Panics(
		func() {
			AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{
				ShortName: 't',
				Required: true,
				Description: "Test description",
				Default: "default string",
			})
			AddArg[argTypes.IpV4Addr,argTypes.IpV4Addr](
				p,"Test",
				&Options[argTypes.IpV4Addr,argTypes.IpV4Addr]{
					ShortName: 't',
					Required: true,
					Description: "Test description",
					Default: argTypes.LocalHost,
				},
			)
		}, "The parser did not panic with duplicate short names.",t,
	);
}

func TestMultipleArgsAllNoShortName(t *testing.T) {
	p:=NewParser("test","testing")
	test.NoPanic(
		func() {
			AddArg[argTypes.String,string](p,"test",&Options[argTypes.String,string]{
				Required: true,
				Description: "Test description",
				Default: "default string",
			})
			AddArg[argTypes.IpV4Addr,argTypes.IpV4Addr](
				p,"Test",
				&Options[argTypes.IpV4Addr,argTypes.IpV4Addr]{
					Required: true,
					Description: "Test description",
					Default: argTypes.LocalHost,
				},
			)
		}, "The parser paniced with multiple args and no short names.",t,
	);
}

func TestAddArgShortLongName(t *testing.T){
	p:=NewParser("test","testing")
	test.Panics(
		func() {
			AddArg[argTypes.String,string](p,"t",&Options[argTypes.String,string]{
				ShortName: 't',
				Required: true,
				Description: "Test description",
				Default: "defaultString",
			})
		}, "Parser did not panic with a single character long name", t,
	)
}

