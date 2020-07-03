package stringManipulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = [][]string{
	{
		"",
		"",
	},
	{
		"aaaaaa€€€€€cccdddaa effgjjjjjjjjjjjj",
		"a6€5c3d3a2 ef2gj9j3",
	},
	{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a5",
	},
	// TODO: escape sequences test(s)
	{
		"\n\n\n",
		"\n3",
	},
	// TODO: entities test(s)
	// {
	// 	"&#39;&#39;&#39;",
	// 	"&#39;&#39;&#39;",
	// },
	// TODO: numbers test(s)
	{
		"11111111111111111111",
		"191912",
	},
	{
		"22222222222222222222",
		"292922",
	},
	{
		"99999999999999999999",
		"999992",
	},
	// TODO: emojis test(s)
	// TODO: other special character tests
}

func TestCompress(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		input, expectedResult := tests[i][0], tests[i][1]
		t.Logf("Compress() should return %s when the input is %s", input, expectedResult)
		result := Compress(input)
		assert.Equal(t, expectedResult, result)
	}
}

func TestUnpack(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		input, expectedResult := tests[i][1], tests[i][0]
		t.Logf("Unpack() should return %s when the input is %s", input, expectedResult)
		result := Unpack(input)
		assert.Equal(t, expectedResult, result)
	}
}

var errorTests = [][]string{
	{
		"a99999999999999999999",
		"a999992",
	},
	{
		"aaaaaaaaa999999999999999999",
		"a99999",
	},
}

func TestUnpackFails(t *testing.T) {
	for i := 0; i < len(errorTests); i++ {
		input, expectedResult := tests[i][1], tests[i][0]
		t.Logf("Unpack() should return %s when the input is %s", input, expectedResult)
		result := Unpack(input)
		assert.Equal(t, expectedResult, result)
	}
}
