package appactions

import (
	"testing"

	testenum "github.com/barbell-math/util/src/appActions/testEnum"
	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/test"
)

func TestAppActionParser(t *testing.T) {
	res := struct{ Action testenum.TestEnum }{}

	p, err := (&argparse.ArgBuilder{}).ToParser("", "")
	test.Nil(err, t)
	err = p.AddSubParsers(
		NewAppActionParser[testenum.TestEnum, *testenum.TestEnum](
			&res.Action,
			testenum.UnknownAppAction,
			[]argparse.ArgConditionality[testenum.TestEnum]{},
		),
	)
	test.Nil(err, t)

	res.Action = testenum.TestEnum(-1)
	err = p.Parse(argparse.ArgvIterFromSlice([]string{}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.UnknownAppAction, res.Action, t)

	res.Action = testenum.UnknownAppAction
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-a", "asdf"}).ToTokens())
	test.ContainsError(testenum.InvalidTestEnum, err, t)
	test.Eq(testenum.UnknownAppAction, res.Action, t)

	res.Action = testenum.UnknownAppAction
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-a", "UnknownAppAction"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.UnknownAppAction, res.Action, t)

	res.Action = testenum.UnknownAppAction
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-a", "AppActionOne"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.AppActionOne, res.Action, t)

	res.Action = testenum.UnknownAppAction
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-a", "AppActionTwo"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.AppActionTwo, res.Action, t)
}
