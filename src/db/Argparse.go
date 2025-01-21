package db

import (
	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
)

//go:generate ../../bin/flags -type=RequiredArgs -package=db

type (
	// A bit flag that is used to set which arguments are required
	RequiredArgs int
	ArgparseVals struct {
		User       string
		EnvPswdVar string
		NetLoc     string
		Port       uint16
		DBName     string
	}
)

const (
	// A flag that will be set to one if the dbUser arg is required.
	UserRequired RequiredArgs = 1 << iota
	// A flag that will be set to one if the dbEnvPswdVar arg is required.
	EnvPswdVarRequired
	// A flag that will be set to one if the dbNetLoc arg is required.
	NetLocRequired
	// A flag that will be set to one if the dbPort arg is required.
	PortRequired
	// A flag that will be set to one if the dbName arg is required.
	DBNameRequired
)

// Returns a parser that can be added to the main application argparse parser
// as a sub parser. This parser adds the following arguments and populates a
// [ArgparseVals] struct.
//
//	--dbUser (-u)
//	--dbEnvPswdVar (-p)
//	--dbNetLoc
//	--dbPort
//	--dbName
func (a *ArgparseVals) GetParser(
	reqArgs RequiredArgs,
	defaults ArgparseVals,
) *argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddArg[string, translators.BuiltinString](
		&a.User,
		&b,
		"dbUser",
		argparse.NewOpts[string, translators.BuiltinString]().
			SetArgType(argparse.ValueArgType).
			SetShortName('u').
			SetDefaultVal(defaults.User).
			SetRequired(reqArgs.GetFlag(UserRequired)).
			SetDescription("The user to use when accessing the database."),
	)
	argparse.AddArg[string, translators.BuiltinString](
		&a.EnvPswdVar,
		&b,
		"dbEnvPswdVar",
		argparse.NewOpts[string, translators.BuiltinString]().
			SetArgType(argparse.ValueArgType).
			SetShortName('p').
			SetDefaultVal(defaults.EnvPswdVar).
			SetRequired(reqArgs.GetFlag(EnvPswdVarRequired)).
			SetDescription("The environment variable to use to look up the password to access the database."),
	)
	argparse.AddArg[string, translators.BuiltinString](
		&a.NetLoc,
		&b,
		"dbNetLoc",
		argparse.NewOpts[string, translators.BuiltinString]().
			SetArgType(argparse.ValueArgType).
			SetDefaultVal(defaults.NetLoc).
			SetRequired(reqArgs.GetFlag(NetLocRequired)).
			SetDescription("The network path to use when connecting to the database."),
	)
	argparse.AddArg[uint16, translators.BuiltinUint16](
		&a.Port,
		&b,
		"dbPort",
		argparse.NewOpts[uint16, translators.BuiltinUint16]().
			SetArgType(argparse.ValueArgType).
			SetDefaultVal(defaults.Port).
			SetRequired(reqArgs.GetFlag(PortRequired)).
			SetDescription("The port to use when connecting to the database."),
	)
	argparse.AddArg[string, translators.BuiltinString](
		&a.DBName,
		&b,
		"dbName",
		argparse.NewOpts[string, translators.BuiltinString]().
			SetArgType(argparse.ValueArgType).
			SetDefaultVal(defaults.DBName).
			SetRequired(reqArgs.GetFlag(DBNameRequired)).
			SetDescription("The name of the database to connect to."),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}
