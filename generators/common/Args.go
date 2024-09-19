package common

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type (
	argsOptions struct {
		printProgName    bool
		mustHaveShowInfo bool
		printArgs        bool
	}
)

var (
	ArgParseExitOnFail bool = true
)

func GetProgName(args []string) string {
	lastSplit := strings.LastIndex(args[0], "/")
	if lastSplit > 0 && lastSplit < len(args[0])-1 {
		return args[0][lastSplit+1:]
	}
	return args[0]
}

func CommentArgs(globalStruct any, comment CommentArgVals) error {
	flagArgs := []string{GetProgName(os.Args)}
	for arg, val := range comment {
		flagArgs = append(
			flagArgs,
			fmt.Sprintf("-%s=%s", arg, val),
		)
	}
	return parseArgs(globalStruct, flagArgs, &argsOptions{
		printProgName:    false,
		mustHaveShowInfo: false,
		printArgs:        true,
	})
}

func InlineArgs(globalStruct any, args []string) error {
	return parseArgs(globalStruct, args, &argsOptions{
		printProgName:    true,
		mustHaveShowInfo: true,
		printArgs:        true,
	})
}

func parseArgs(globalStruct any, args []string, opts *argsOptions) error {
	if opts.printProgName {
		fmt.Println(fmt.Sprintf("Running prog '%s'", args[0]))
	}

	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	refStructPntr := reflect.TypeOf(globalStruct)
	if refStructPntr.Kind() != reflect.Pointer ||
		(refStructPntr.Kind() == reflect.Pointer && refStructPntr.Elem().Kind() != reflect.Struct) {
		panic(fmt.Sprintf(
			"The global struct paramerter must be a pointer to a struct! Got: %s",
			reflect.TypeOf(globalStruct).Kind(),
		))
	}
	refStructType := refStructPntr.Elem()
	refStructVal := reflect.ValueOf(globalStruct).Elem()

	var showInfoRef *bool
	allArgs := map[string]struct{}{}
	requiredArgs := map[string]struct{}{}
	for i := 0; i < refStructType.NumField(); i++ {
		iterFName := refStructType.Field(i).Name
		iterFKind := refStructVal.Field(i).Kind()
		iterFTag := refStructType.Field(i).Tag
		if !refStructType.Field(i).IsExported() {
			continue
		}

		if !refStructVal.Field(i).CanAddr() {
			panic(fmt.Sprintf(
				"The %s field on the supplied struct was not addressable.",
				iterFName,
			))
		}
		iterFAddr := refStructVal.Field(i).Addr().Interface()

		if iterFName == "ShowInfo" {
			if iterFKind != reflect.Bool {
				panic("The ShowInfo field must be a boolean.")
			}
			showInfoRef = iterFAddr.(*bool)
		}

		helpTag, ok1 := iterFTag.Lookup("help")
		if !ok1 {
			panic(fmt.Sprintf(
				"The supplied struct was missing a help tag on field %s",
				iterFName,
			))
		}

		requiredTag, ok2 := iterFTag.Lookup("required")
		if !ok2 {
			panic(fmt.Sprintf(
				"The supplied struct was missing a required tag on field %s",
				iterFName,
			))
		}
		lowerCaseName := strings.ToLower(iterFName[0:1]) + iterFName[1:]
		requiredArg, err := strconv.ParseBool(requiredTag)
		if err == nil {
			if requiredArg {
				requiredArgs[lowerCaseName] = struct{}{}
			}
		} else {
			panic(fmt.Sprintf(
				"The required flag on field %s was not a valid boolean expression: %s",
				iterFName, err,
			))
		}

		defaultTag, ok2 := iterFTag.Lookup("default")
		if !requiredArg && !ok2 {
			panic(fmt.Sprintf(
				"The supplied struct was missing a default tag on field %s",
				iterFName,
			))
		}
		if requiredArg && ok2 {
			panic(fmt.Sprintf(
				"The supplied struct added a default value to a required argument on field %s",
				iterFName,
			))
		}

		allArgs[lowerCaseName] = struct{}{}

		switch iterFKind {
		case reflect.String:
			flagSet.StringVar(
				iterFAddr.(*string),
				lowerCaseName,
				defaultTag,
				helpTag,
			)
		case reflect.Bool:
			if requiredArg {
				defaultTag = "false"
			}
			defaultVal, err := strconv.ParseBool(defaultTag)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a bool: %s",
					iterFName, err,
				))
			}
			flagSet.BoolVar(
				iterFAddr.(*bool),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Float64:
			if requiredArg {
				defaultTag = "0"
			}
			defaultVal, err := strconv.ParseFloat(defaultTag, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a float64: %s",
					iterFName, err,
				))
			}
			flagSet.Float64Var(
				iterFAddr.(*float64),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Int:
			if requiredArg {
				defaultTag = "0"
			}
			defaultVal, err := strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a int64: %s",
					iterFName, err,
				))
			}
			flagSet.IntVar(
				iterFAddr.(*int),
				lowerCaseName,
				int(defaultVal),
				helpTag,
			)
		case reflect.Int64:
			if requiredArg {
				defaultTag = "0"
			}
			defaultVal, err := strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a int64: %s",
					iterFName, err,
				))
			}
			flagSet.Int64Var(
				iterFAddr.(*int64),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Uint:
			if requiredArg {
				defaultTag = "0"
			}
			defaultVal, err := strconv.ParseUint(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a uint64: %s",
					iterFName, err,
				))
			}
			flag.UintVar(
				iterFAddr.(*uint),
				lowerCaseName,
				uint(defaultVal),
				helpTag,
			)
		case reflect.Uint64:
			if requiredArg {
				defaultTag = "0"
			}
			defaultVal, err := strconv.ParseUint(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a uint64: %s",
					iterFName, err,
				))
			}
			flagSet.Uint64Var(
				iterFAddr.(*uint64),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		default:
			panic(fmt.Sprintf(
				"Field %s has an unsupported type: %s\nSupported Types are: %v",
				iterFName, iterFKind,
				[]string{"string", "int", "int64", "uint", "uint64", "float64"},
			))
		}
	}

	if showInfoRef == nil && opts.mustHaveShowInfo {
		panic("The boolean ShowInfo field must be present in the supplied struct.")
	}

	flagSet.Parse(args[1:])

	requiredCopy := map[string]struct{}{}
	for r, _ := range requiredArgs {
		requiredCopy[r] = struct{}{}
	}
	flagSet.Visit(func(f *flag.Flag) {
		if _, ok := requiredCopy[f.Name]; ok {
			delete(requiredCopy, f.Name)
		}
	})
	if len(requiredCopy) > 0 {
		args := []string{}
		for k, _ := range requiredCopy {
			args = append(args, k)
		}
		PrintRunningError("Not all required flags were passed.")
		PrintRunningError("The following flags must be added: %v", args)

		cntr := 0
		PrintRunningError("Received: ")
		flagSet.Visit(func(f *flag.Flag) {
			cntr++
			PrintRunningError("|- (%d) %s: %+v", cntr, f.Name, f.Value)
		})
		PrintRunningError("The accepted flags are as follows:")
		flagSet.PrintDefaults()
		if ArgParseExitOnFail {
			os.Exit(1)
		}
		return MissingRequiredArgs
	}

	if opts.printArgs {
		PrintRunningInfo("Received arguments:")
		for i := 0; i < refStructType.NumField(); i++ {
			if !refStructType.Field(i).IsExported() {
				continue
			}
			iterFName := refStructType.Field(i).Name
			iterFVal := refStructVal.Field(i).Interface()
			PrintRunningInfo(
				"|- (%d) | Name: %-20s | Value: %v",
				i+1, iterFName, iterFVal,
			)
		}
	}

	return nil
}
