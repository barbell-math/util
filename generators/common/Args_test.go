package common

import (
	"flag"
	"os"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestArgsNonStruct(t *testing.T) {
	test.Panics(
		func() {
			tmp := 1
			Args(tmp, []string{"progName"})
		},
		t,
	)
	test.Panics(
		func() {
			tmp := struct{}{}
			Args(tmp, []string{"progName"})
		},
		t,
	)
}

func TestArgsStruct(t *testing.T) {
	test.NoPanic(
		func() {
			tmp := struct{}{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestDuplicatedShowInfoField(t *testing.T) {
	type bad struct {
		ShowInfo int
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingHelpTag(t *testing.T) {
	type bad struct {
		A int
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingDefaultTag(t *testing.T) {
	type bad struct {
		A int `help:""`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestMissingRequiredTag(t *testing.T) {
	type bad struct {
		A int `help:"" default:""`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidRequiredTag(t *testing.T) {
	type bad struct {
		A int `help:"" default:"" required:"foobar"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidBoolDefault(t *testing.T) {
	type bad struct {
		A bool `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidFloat64Default(t *testing.T) {
	type bad struct {
		A float64 `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidIntDefault(t *testing.T) {
	type bad struct {
		A int `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidInt64Default(t *testing.T) {
	type bad struct {
		A int64 `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidUintDefault(t *testing.T) {
	type bad struct {
		A uint `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestInvalidUint64Default(t *testing.T) {
	type bad struct {
		A uint64 `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestUnsupportedType(t *testing.T) {
	type bad struct {
		A complex64 `help:"" default:"foobar" required:"false"`
	}
	test.Panics(
		func() {
			tmp := bad{}
			Args(&tmp, []string{"progName"})
		},
		t,
	)
}

func TestNotEnoughArgs(t *testing.T) {
	type bad struct {
		A bool `help:"" default:"false" required:"true"`
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	exitOnFail = false
	tmp := bad{}
	test.ContainsError(
		NotEnoughArgs,
		Args(&tmp, []string{"progName"}),
		t,
	)
	exitOnFail = true
}

func TestMissingRequiredArgs(t *testing.T) {
	type bad struct {
		A bool `help:"" default:"false" required:"true"`
		B bool `help:"" default:"false" required:"false"`
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	exitOnFail = false
	tmp := bad{}
	test.ContainsError(
		MissingRequiredArgs,
		Args(&tmp, []string{"progName", "-b=true"}),
		t,
	)
	exitOnFail = true
}

func TestPassingArgs(t *testing.T) {
	type good struct {
		A bool   `help:"" default:"false" required:"true"`
		B string `help:"" default:"false" required:"false"`
		C int    `help:"" default:"-1" required:"false"`
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	exitOnFail = false
	tmp := good{}
	test.Nil(Args(&tmp, []string{"progName", "-b=true", "-a=false"}), t)
	exitOnFail = true
	test.False(tmp.A, t)
	test.Eq("true", tmp.B, t)
	test.Eq(-1, tmp.C, t)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	exitOnFail = false
	tmp = good{}
	test.Nil(Args(&tmp, []string{"progName", "-b=foobar", "-a=false", "-c=3"}), t)
	exitOnFail = true
	test.False(tmp.A, t)
	test.Eq("foobar", tmp.B, t)
	test.Eq(3, tmp.C, t)
}
