package bencode_test

import (
	"testing"

	"github.com/nivaldogmelo/bittorrent-client/pkg/bencode"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	testCaseOnlyInt := `i235454353e`
	expectedResultOnlyInt := 235454353
	resultOnlyInt, _ := bencode.Parser(testCaseOnlyInt)

	testCaseOnlyString := `8:mystring`
	expectedResultOnlyString := "mystring"
	resultOnlyString, _ := bencode.Parser(testCaseOnlyString)

	assert.Equal(t, expectedResultOnlyInt, resultOnlyInt)
	assert.Equal(t, expectedResultOnlyString, resultOnlyString)
}
