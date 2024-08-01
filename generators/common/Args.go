package common

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	exitOnFail bool = true
)

func Args(globalStruct any, args []string) error {
	fmt.Println(fmt.Sprintf("Running prog '%s", args[0]))

	refStructPntr := reflect.TypeOf(globalStruct)
	if refStructPntr.Kind() != reflect.Pointer && refStructPntr.Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf(
			"The global struct paramerter must be a pointer to a struct! Got: %s",
			reflect.TypeOf(globalStruct).Kind(),
		))
	}
	refStructType := refStructPntr.Elem()
	refStructVal := reflect.ValueOf(globalStruct).Elem()

	requiredArgs := []string{}
	for i := 0; i < refStructType.NumField(); i++ {
		iterFName := refStructType.Field(i).Name
		iterFKind := refStructVal.Field(i).Kind()
		iterFTag := refStructType.Field(i).Tag
		if !refStructType.Field(i).IsExported() {
			continue
		}

		if iterFName == "ShowInfo" {
			panic("The field name 'ShowInfo' is reserved.")
		}
		if !refStructVal.Field(i).CanAddr() {
			panic(fmt.Sprintf(
				"The %s field on the supplied struct was not addressable.",
				iterFName,
			))
		}
		iterFAddr := refStructVal.Field(i).Addr().Interface()

		helpTag, ok1 := iterFTag.Lookup("help")
		if !ok1 {
			panic(fmt.Sprintf(
				"The supplied struct was missing a help tag on field %s",
				iterFName,
			))
		}
		defaultTag, ok2 := iterFTag.Lookup("default")
		if !ok2 {
			panic(fmt.Sprintf(
				"The supplied struct was missing a default tag on field %s",
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
		if val, err := strconv.ParseBool(requiredTag); err == nil {
			if val {
				requiredArgs = append(requiredArgs, lowerCaseName)
			}
		} else {
			panic(fmt.Sprintf(
				"The required flag on field %s was not a valid boolean expression: %s",
				iterFName, err,
			))
		}

		switch iterFKind {
		case reflect.String:
			flag.StringVar(
				iterFAddr.(*string),
				lowerCaseName,
				defaultTag,
				helpTag,
			)
		case reflect.Bool:
			defaultVal, err := strconv.ParseBool(defaultTag)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a bool: %s",
					iterFName, err,
				))
			}
			flag.BoolVar(
				iterFAddr.(*bool),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Float64:
			defaultVal, err := strconv.ParseFloat(defaultTag, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a float64: %s",
					iterFName, err,
				))
			}
			flag.Float64Var(
				iterFAddr.(*float64),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Int:
			defaultVal, err := strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a int64: %s",
					iterFName, err,
				))
			}
			flag.IntVar(
				iterFAddr.(*int),
				lowerCaseName,
				int(defaultVal),
				helpTag,
			)
		case reflect.Int64:
			defaultVal, err := strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a int64: %s",
					iterFName, err,
				))
			}
			flag.Int64Var(
				iterFAddr.(*int64),
				lowerCaseName,
				defaultVal,
				helpTag,
			)
		case reflect.Uint:
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
			defaultVal, err := strconv.ParseUint(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(
					"Could not parse field %s as a uint64: %s",
					iterFName, err,
				))
			}
			flag.Uint64Var(
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
	var showInfo bool
	flag.BoolVar(&showInfo, "showInfo", true, "Print out the receved values or not.")

	if len(args)-1 < len(requiredArgs) {
		PrintRunningError("Not enough arguments.")
		PrintRunningError("Received: ", args[1:])
		PrintRunningError("The accepted flags are as follows:")
		flag.PrintDefaults()
		PrintRunningError("Re-run go generate after fixing the problem.")
		if exitOnFail {
			os.Exit(1)
		}
		return NotEnoughArgs
	}

	flag.CommandLine.Parse(args[1:])

	requiredCopy := append([]string{}, requiredArgs...)
	flag.CommandLine.Visit(func(f *flag.Flag) {
		for i, v := range requiredCopy {
			if f.Name == v {
				requiredCopy = append(requiredCopy[:i], requiredCopy[i+1:]...)
			}
		}
	})
	if len(requiredCopy) > 0 {
		PrintRunningError("Not all required flags were passed.")
		PrintRunningError("The following flags must be added: ", requiredCopy)

		cntr := 0
		fmt.Println("Received: ")
		flag.Visit(func(f *flag.Flag) {
			cntr++
			PrintRunningError(" |- (%d) %s: %+v\n", cntr, f.Name, f.Value)
		})
		PrintRunningError("The accepted flags are as follows:")
		flag.PrintDefaults()
		if exitOnFail {
			os.Exit(1)
		}
		return MissingRequiredArgs
	}

	if showInfo {
		PrintRunningInfo("Received arguments:")
		for i := 0; i < refStructType.NumField(); i++ {
			if !refStructType.Field(i).IsExported() {
				continue
			}
			iterFName := refStructType.Field(i).Name
			iterFVal := refStructVal.Field(i).Interface()
			PrintRunningInfo(
				" |- (%d) | Name: %-20s | Value: %v",
				i+1, iterFName, iterFVal,
			)
		}
	}

	return nil
}

func PrintRunningInfo(fmtStr string, args ...any) {
	fmtStr = " |- " + fmtStr + "\n"
	fmt.Printf(fmtStr, args...)
}

func PrintRunningError(fmtStr string, args ...any) {
	fmtStr = " |- ERROR: " + fmtStr + "\n"
	fmt.Printf(fmtStr, args...)
}
