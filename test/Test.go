package test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

func FormatError(expected any, got any, base string, file string, line int, t *testing.T){
    t.Fatal(fmt.Sprintf(
        "Error | File %s: Line %d: %s\nExpected: '%v'\nGot: '%v'",
        file,line,base,expected,got,
    ));
}

func ContainsError(expected error, got error, t *testing.T){
    if !errors.Is(got,expected) {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            expected,
            got,
            "The expected error was not contained in the given error.",
            f,l,t,
        )
    }
}

func Panics(action func(), t *testing.T){
    defer func() {
        if r:=recover(); r==nil {
            _, f, l, _ := runtime.Caller(1)
            FormatError(
                "panic","",
                "The supplied funciton did not panic when it should have.",
                f,l,t,
            )
        }
    }()
    action()
}

func NoPanic(action func(), t *testing.T){
    defer func() {
        if r:=recover(); r!=nil {
            _, f, l, _ := runtime.Caller(1)
            FormatError(
                "panic","",
                "The supplied funciton paniced when it should have.",
                f,l,t,
            )
        }
    }()
    action()
}

func Eq(l any, r any, t *testing.T) {
    if r!=l {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            l,r,
            "The supplied values were not equal but were expected to be.",
            f,l,t,
        )
    }
}

func Neq(l any, r any, t *testing.T) {
    if r==l {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            l,r,
            "The supplied values were equal but were expected to not be.",
            f,l,t,
        )
    }
}

func True(v bool, t *testing.T) {
    if v!=true {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            true,v,
            "The supplied value was not true when it was expected to be.",
            f,l,t,
        )
    }
}

func False(v bool, t *testing.T) {
    if v!=false {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            false,v,
            "The supplied value was not false when it was expected to be.",
            f,l,t,
        )
    }
}

func Nil(v any, t *testing.T) {
    if v!=nil {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            nil,v,
            "The supplied value was not nil when it was expected to be.",
            f,l,t,
        )
    }
}

func NilPntr[T any](v *T, t *testing.T) {
    if v!=(*T)(nil) {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            nil,v,
            "The supplied value was not nil when it was expected to be.",
            f,l,t,
        )
    }
}

func NotNil(v any, t *testing.T) {
    if v==nil {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            "!nil",v,
            "The supplied value was nil when it was not expected to be.",
            f,l,t,
        )
    }
}

func NotNilPntr[T any](v *T, t *testing.T) {
    if v==(*T)(nil) {
        _, f, l, _ := runtime.Caller(1)
        FormatError(
            nil,v,
            "The supplied value was not nil when it was expected to be.",
            f,l,t,
        )
    }
}

func SlicesMatch[T any](actual []T, generated []T, t *testing.T){
    _, f, l, _ := runtime.Caller(1)
    if len(actual)!=len(generated) {
        FormatError(
            len(actual),
            len(generated),
            "Slices do not match in length.",
            f,l,t,
        )
    }
    min:=len(generated);
    if len(actual)<len(generated) {
        min=len(actual);
    }
    for i:=0; i<min; i++ {
        if any(actual[i])!=any(generated[i]) {
            FormatError(
                actual[i],
                generated[i],
                fmt.Sprintf("Values do not match | Index: %d",i),
                f,l,t,
            )
        }
    }
}
