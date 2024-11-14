package common

import (
	"flag"
	"os"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestInlineArgsNonStruct(t *testing.T) {
	test.Panics(
		func() {
			tmp := 1
			InlineArgs(tmp, []string{"progName"})
		},
		t,
	)
	test.Panics(
		func() {
			tmp := struct{}{}
			InlineArgs(tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingShowInfoStructField(t *testing.T) {
	type bad struct {
		A int
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestShowInfoStructFieldBadType(t *testing.T) {
	type bad struct {
		ShowInfo int
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingHelpTag(t *testing.T) {
	type bad struct {
		A        int
		ShowInfo bool
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingDefaultTag(t *testing.T) {
	type bad struct {
		A        int  `help:""`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingRequiredTag(t *testing.T) {
	type bad struct {
		A        int  `help:"" default:""`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidRequiredTag(t *testing.T) {
	type bad struct {
		A        int  `help:"" default:"" required:"foobar"`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidBoolDefault(t *testing.T) {
	type bad struct {
		A        bool `help:"" default:"foobar" required:"false"`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidFloat64Default(t *testing.T) {
	type bad struct {
		A        float64 `help:"" default:"foobar" required:"false"`
		ShowInfo bool    `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidIntDefault(t *testing.T) {
	type bad struct {
		A        int  `help:"" default:"foobar" required:"false"`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidInt64Default(t *testing.T) {
	type bad struct {
		A        int64 `help:"" default:"foobar" required:"false"`
		ShowInfo bool  `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidUintDefault(t *testing.T) {
	type bad struct {
		A        uint `help:"" default:"foobar" required:"false"`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidUint64Default(t *testing.T) {
	type bad struct {
		A        uint64 `help:"" default:"foobar" required:"false"`
		ShowInfo bool   `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestUnsupportedType(t *testing.T) {
	type bad struct {
		A        complex64 `help:"" default:"foobar" required:"false"`
		ShowInfo bool      `required:"f" default:"f" help:"Show debug info."`
	}
	test.Panics(
		func() {
			tmp := bad{}
			InlineArgs(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingRequiredArgs(t *testing.T) {
	type bad struct {
		A        bool `help:"" required:"true"`
		B        bool `help:"" default:"false" required:"false"`
		ShowInfo bool `required:"f" default:"f" help:"Show debug info."`
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	ArgParseExitOnFail = false
	tmp := bad{}
	test.ContainsError(
		MissingRequiredArgs,
		InlineArgs(&tmp, []string{"progName", "-b=true"}),
		t,
	)
	ArgParseExitOnFail = true
}

func TestPassingInlineArgs(t *testing.T) {
	type good struct {
		A        bool   `help:"" required:"true"`
		B        string `help:"" default:"false" required:"false"`
		C        int    `help:"" default:"-1" required:"false"`
		ShowInfo bool   `required:"f" default:"f" help:"Show debug info."`
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	ArgParseExitOnFail = false
	tmp := good{}
	test.Nil(InlineArgs(&tmp, []string{"progName", "-b=true", "-a=false"}), t)
	ArgParseExitOnFail = true
	test.False(tmp.A, t)
	test.Eq("true", tmp.B, t)
	test.Eq(-1, tmp.C, t)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	ArgParseExitOnFail = false
	tmp = good{}
	test.Nil(InlineArgs(&tmp, []string{"progName", "-b=foobar", "-a=false", "-c=3"}), t)
	ArgParseExitOnFail = true
	test.False(tmp.A, t)
	test.Eq("foobar", tmp.B, t)
	test.Eq(3, tmp.C, t)
}
