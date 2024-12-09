package xrunes_test

import (
	"testing"

	runes "github.com/jolt9dev/go-xrunes"
	"github.com/stretchr/testify/assert"
)

func TestEqualFold(t *testing.T) {
	assert.Equal(t, runes.EqualFold([]rune("test"), []rune("TEST")), true)
	assert.Equal(t, runes.EqualFold([]rune("test"), []rune("Test")), true)
	assert.Equal(t, runes.EqualFold([]rune("test"), []rune("tEsT")), true)
	assert.Equal(t, runes.EqualFold([]rune("test"), []rune("teSt")), true)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("test ")), false)
	assert.Equal(t, runes.Equal([]rune("test"), []rune(" test")), false)
}

func TestEqual(t *testing.T) {
	assert.Equal(t, runes.Equal([]rune("test"), []rune("TEST")), false)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("Test")), false)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("tEsT")), false)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("teSt")), false)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.Equal([]rune("test"), []rune("test ")), false)
}

func TestIndex(t *testing.T) {
	assert.Equal(t, runes.Index([]rune("test"), []rune("test")), 0)
	assert.Equal(t, runes.Index([]rune("test"), []rune("test ")), -1)
	assert.Equal(t, runes.Index([]rune("test"), []rune("TEST")), -1)
	assert.Equal(t, runes.Index([]rune("test"), []rune("Test")), -1)
	assert.Equal(t, runes.Index([]rune("test"), []rune("tEsT")), -1)
	assert.Equal(t, runes.Index([]rune("test"), []rune("teSt")), -1)
}

func TestIndexFold(t *testing.T) {
	assert.Equal(t, 0, runes.IndexFold([]rune("test"), []rune("test")))
	assert.Equal(t, 0, runes.IndexFold([]rune("test"), []rune("TEST")))
	assert.Equal(t, 0, runes.IndexFold([]rune("test"), []rune("Test")))
	assert.Equal(t, 0, runes.IndexFold([]rune("test"), []rune("tEsT")))
	assert.Equal(t, 0, runes.IndexFold([]rune("test"), []rune("teSt")))
	assert.Equal(t, -1, runes.IndexFold([]rune("test"), []rune("test ")))
	assert.Equal(t, 1, runes.IndexFold([]rune(" test "), []rune("teSt")))
}

func TestContains(t *testing.T) {
	assert.Equal(t, runes.Contains([]rune("my TEtestst"), []rune("test")), true)
	assert.Equal(t, runes.Contains([]rune("test"), []rune("test ")), false)
	assert.Equal(t, runes.Contains([]rune("test"), []rune("TEST")), false)
	assert.Equal(t, runes.Contains([]rune("test"), []rune("Test")), false)
	assert.Equal(t, runes.Contains([]rune("test"), []rune("tEsT")), false)
	assert.Equal(t, runes.Contains([]rune("test"), []rune("teSt")), false)
}

func TestContainsFold(t *testing.T) {
	assert.Equal(t, runes.ContainsFold([]rune("my TEtestst	"), []rune("test")), true)
	assert.Equal(t, runes.ContainsFold([]rune("test"), []rune("test ")), false)
	assert.Equal(t, runes.ContainsFold([]rune("test"), []rune("TEST")), true)
	assert.Equal(t, runes.ContainsFold([]rune("test"), []rune("Test")), true)
	assert.Equal(t, runes.ContainsFold([]rune("test"), []rune("tEsT")), true)
	assert.Equal(t, runes.ContainsFold([]rune("test"), []rune("teSt")), true)
}

func TestHasPrefix(t *testing.T) {
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("test ")), false)
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("TEST")), false)
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("Test")), false)
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("tEsT")), false)
	assert.Equal(t, runes.HasPrefix([]rune("test"), []rune("teSt")), false)
}

func TestHasPrefixFold(t *testing.T) {
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("test ")), false)
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("TEST")), true)
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("Test")), true)
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("tEsT")), true)
	assert.Equal(t, runes.HasPrefixFold([]rune("test"), []rune("teSt")), true)
}

func TestHasSuffix(t *testing.T) {
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune(" test")), false)
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune("TEST")), false)
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune("Test")), false)
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune("tEsT")), false)
	assert.Equal(t, runes.HasSuffix([]rune("test"), []rune("teSt")), false)
}

func TestHasSuffixFold(t *testing.T) {
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune("test")), true)
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune(" test")), false)
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune("TEST")), true)
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune("Test")), true)
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune("tEsT")), true)
	assert.Equal(t, runes.HasSuffixFold([]rune("test"), []rune("teSt")), true)
}

func TestTrim(t *testing.T) {
	assert.Equal(t, runes.Trim([]rune("  test  "), []rune(" ")), []rune("test"))
	assert.Equal(t, runes.Trim([]rune("xxxtestxxx"), []rune("x")), []rune("test"))
	assert.Equal(t, runes.Trim([]rune("test"), []rune(" ")), []rune("test"))
	assert.Equal(t, runes.Trim([]rune(""), []rune(" ")), []rune(""))
}

func TestTrimLeft(t *testing.T) {
	assert.Equal(t, runes.TrimLeft([]rune("  test  "), []rune(" ")), []rune("test  "))
	assert.Equal(t, runes.TrimLeft([]rune("xxxtestxxx"), []rune("x")), []rune("testxxx"))
	assert.Equal(t, runes.TrimLeft([]rune("test"), []rune(" ")), []rune("test"))
	assert.Equal(t, runes.TrimLeft([]rune(""), []rune(" ")), []rune(""))
}

func TestTrimRight(t *testing.T) {
	assert.Equal(t, runes.TrimRight([]rune("  test  "), []rune(" ")), []rune("  test"))
	assert.Equal(t, runes.TrimRight([]rune("xxxtestxxx"), []rune("x")), []rune("xxxtest"))
	assert.Equal(t, runes.TrimRight([]rune("test"), []rune(" ")), []rune("test"))
	assert.Equal(t, runes.TrimRight([]rune(""), []rune(" ")), []rune(""))
}

func TestIsSpace(t *testing.T) {
	assert.Equal(t, runes.IsSpace([]rune("   ")), true)
	assert.Equal(t, runes.IsSpace([]rune("test")), false)
	assert.Equal(t, runes.IsSpace([]rune("")), true)
}

func TestIsEmptySpace(t *testing.T) {
	assert.Equal(t, runes.IsEmptySpace([]rune("   ")), true)
	assert.Equal(t, runes.IsEmptySpace([]rune("test")), false)
	assert.Equal(t, runes.IsEmptySpace([]rune("")), true)
}
