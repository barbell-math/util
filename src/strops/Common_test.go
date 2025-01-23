package strops

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestWorpWrapLines(t *testing.T) {
	line := "this is a line"
	wrapped := WordWrapLines(line, 7)
	test.Eq(len(wrapped), 2, t)
	test.Eq(wrapped[0], "this is", t)
	test.Eq(wrapped[1], "a line", t)
}

func TestWorpWrapLinesShorterThanMaxLen(t *testing.T) {
	line := "this is a line"
	wrapped := WordWrapLines(line, 20)
	test.Eq(len(wrapped), 1, t)
	test.Eq(wrapped[0], "this is a line", t)
}

func TestWorpWrapLinesNoText(t *testing.T) {
	line := ""
	wrapped := WordWrapLines(line, 20)
	test.Eq(len(wrapped), 1, t)
	test.Eq(wrapped[0], "", t)
}
