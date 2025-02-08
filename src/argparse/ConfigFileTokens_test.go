package argparse

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestValidConfigFile(t *testing.T) {
	tokens, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/ValidConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)

	tokens, err = ArgvIterFromSlice(
		[]string{"--config=./testData/ValidConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)
}

func TestGroupMissingOpenParenSyntaxErr(t *testing.T) {
	_, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/GroupMissingParenConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/GroupMissingParenConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)
}

func TestNameMissingValueSyntaxErr(t *testing.T) {
	_, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/NameMissingValueConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/NameMissingValueConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)
}

func TestToManyClosingBrackets(t *testing.T) {
	_, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/ToManyClosingBracketsConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/ToManyClosingBracketsConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)
}

func TestNotEnoughClosingBracketsSyntaxErr(t *testing.T) {
	_, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/NotEnoughClosingBracketsConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/NotEnoughClosingBracketsConfigFile.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.ContainsError(ParserConfigFileErr, err, t)
	test.ContainsError(ParserConfigFileSyntaxErr, err, t)
}

func TestValidConfigFileWithComments(t *testing.T) {
	tokens, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/ValidConfigFileWithComments.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/ValidConfigFileWithComments.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)
}

func TestValidConfigFileWithBlankLines(t *testing.T) {
	tokens, err := ArgvIterFromSlice(
		[]string{"--config", "./testData/ValidConfigFileBlankLines.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)

	_, err = ArgvIterFromSlice(
		[]string{"--config=./testData/ValidConfigFileBlankLines.txt"},
	).
		ToTokens().
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(tokens), 14, t)
	test.Eq(tokens[0], token{"Name0", longFlagToken}, t)
	test.Eq(tokens[1], token{"Value0", valueToken}, t)
	test.Eq(tokens[2], token{"Group1Group2Name1", longFlagToken}, t)
	test.Eq(tokens[3], token{"Value1", valueToken}, t)
	test.Eq(tokens[4], token{"Group1Group2Group3Name2", longFlagToken}, t)
	test.Eq(tokens[5], token{"Value2", valueToken}, t)
	test.Eq(tokens[6], token{"Group1Group2Name3", longFlagToken}, t)
	test.Eq(tokens[7], token{"Value3", valueToken}, t)
	test.Eq(tokens[8], token{"Group1Name4", longFlagToken}, t)
	test.Eq(tokens[9], token{"Value4", valueToken}, t)
	test.Eq(tokens[10], token{"Group1Name5", longFlagToken}, t)
	test.Eq(tokens[11], token{"Value5", valueToken}, t)
	test.Eq(tokens[12], token{"Name6", longFlagToken}, t)
	test.Eq(tokens[13], token{"Value6", valueToken}, t)
}
