package bencode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInt(t *testing.T) {
	expectedInt := 235454353
	testCase1 := `i235454353e`
	_, resultInt1, _ := createInt(testCase1)

	testCase2 := `i235454353e6:string`
	_, resultInt2, _ := createInt(testCase2)

	assert.Equal(t, expectedInt, resultInt1)
	assert.Equal(t, expectedInt, resultInt2)
}

func TestCreateString(t *testing.T) {
	testCase1 := `41:http://bttracker.debian.org:6969/announce`
	expectedResult1 := "http://bttracker.debian.org:6969/announce"
	_, result1, _ := createString(testCase1)

	testCase2 := `6:stringi42e`
	expectedResult2 := "string"
	_, result2, _ := createString(testCase2)

	assert.Equal(t, expectedResult1, result1)
	assert.Equal(t, expectedResult2, result2)
}
