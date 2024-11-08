package argparse

import (
	"testing"

	"github.com/barbell-math/util/test"
)

func TestToTokensNoArguments(t *testing.T) {
	tokens, err:=ArgvIterFromSlice([]string{}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.Eq(0, len(tokens), t)
}

func TestToTokensShortFlag(t *testing.T) {
	tokens, err:=ArgvIterFromSlice([]string{"-t"}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{token{value: "t", _type: shortFlagToken}},
		t,
	)
}

func TestToTokensLongSpaceFlag(t *testing.T) {
	tokens, err:=ArgvIterFromSlice([]string{"--time"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{token{value: "time", _type: longFlagToken}},
		t,
	)
}

func TestToTokensLongEqualsFlag(t *testing.T) {
	tokens, err:=ArgvIterFromSlice([]string{"--time=123"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "time", _type: longFlagToken},
			{value: "123", _type: argumentToken},
		},
		t,
	)
}

func TestToTokensArgumentFlag(t *testing.T) {
	tokens, err:=ArgvIterFromSlice([]string{"argument"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "argument", _type: argumentToken},
		},
		t,
	)
}

func TestToTokensAllFlags(t *testing.T) {
	tokens, err:=ArgvIterFromSlice(
		[]string{"-t", "--time", "123", "--time=123"},
	).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "t", _type: shortFlagToken},
			{value: "time", _type: longFlagToken},
			{value: "123", _type: argumentToken},
			{value: "time", _type: longFlagToken},
			{value: "123", _type: argumentToken},
		},
		t,
	)
}

// func TestToPairsMissingFlag(t *testing.T) {
// 	tokenPairs, err:=ArgvIterFromSlice([]string{"123", "123"}).
// 		ToTokens().ToPairs().Collect()
// 	test.ContainsError(ExpectedFlag, err, t)
// 	test.Eq(0, len(tokenPairs), t)
// }
// 
// func TestToPairsMissingArgument(t *testing.T) {
// 	tokenPairs, err:=ArgvIterFromSlice([]string{"-t", "-t"}).
// 		ToTokens().ToPairs().Collect()
// 	test.ContainsError(ExpectedFlag, err, t)
// 	test.Eq(0, len(tokenPairs), t)
// 
// 	tokenPairs, err=ArgvIterFromSlice([]string{"--t", "--t"}).
// 		ToTokens().ToPairs().Collect()
// 	test.ContainsError(ExpectedFlag, err, t)
// 	test.Eq(0, len(tokenPairs), t)
// 
// 	tokenPairs, err=ArgvIterFromSlice([]string{"-t", "--t"}).
// 		ToTokens().ToPairs().Collect()
// 	test.ContainsError(ExpectedFlag, err, t)
// 	test.Eq(0, len(tokenPairs), t)
// 
// 	tokenPairs, err=ArgvIterFromSlice([]string{"--t", "-t"}).
// 		ToTokens().ToPairs().Collect()
// 	test.ContainsError(ExpectedFlag, err, t)
// 	test.Eq(0, len(tokenPairs), t)
// }
// 
// func TestToPairsPassing(t *testing.T) {
// 	exp:=[]containers.Vector[token, *token]{
// 		containers.VectorValInit[token, *token](
// 			token{value: "t", _type: shortFlagToken},
// 			token{value: "123", _type: argumentToken},
// 		),
// 	}
// 	err:=ArgvIterFromSlice([]string{"-t", "123"}).
// 		ToTokens().ToPairs().ForEach(
// 			func(
// 				index int, val staticContainers.Vector[token],
// 			) (iter.IteratorFeedback, error) {
// 				test.True(val.KeyedEq(&exp[index]), t)
// 				return iter.Continue, nil
// 			},
// 		)
// 	test.Nil(err, t)
// }
// 
// func TestToPairsMultiplePairs(t *testing.T) {
// 	exp:=[]containers.Vector[token, *token]{
// 		containers.VectorValInit[token, *token](
// 			token{value: "t", _type: shortFlagToken},
// 			token{value: "123", _type: argumentToken},
// 		),
// 		containers.VectorValInit[token, *token](
// 			token{value: "time", _type: longFlagToken},
// 			token{value: "456", _type: argumentToken},
// 		),
// 		containers.VectorValInit[token, *token](
// 			token{value: "time2", _type: longFlagToken},
// 			token{value: "789", _type: argumentToken},
// 		),
// 	}
// 	err:=ArgvIterFromSlice([]string{"-t", "123", "--time", "456", "--time2=789"}).
// 		ToTokens().ToPairs().ForEach(
// 			func(
// 				index int, val staticContainers.Vector[token],
// 			) (iter.IteratorFeedback, error) {
// 				test.True(val.KeyedEq(&exp[index]), t)
// 				return iter.Continue, nil
// 			},
// 		)
// 	test.Nil(err, t)
// }
