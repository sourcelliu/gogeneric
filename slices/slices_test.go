package slices

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToMap(t *testing.T) {
	strSlice := []string{"4", "2", "3"}
	strFunc := func(origin string) string {
		return origin
	}
	strMap := SliceToMap(strSlice, strFunc)
	assert.Equal(t, len(strMap), 3)
	assert.Equal(t, strMap["2"], "2")

	aSlice := make([]*A, 0)
	for i := 1; i < 4; i++ {
		tempa := &A{
			id:   i,
			name: "sourcelliu:" + strconv.Itoa(i),
		}
		aSlice = append(aSlice, tempa)
	}
	aFunc := func(origin *A) string {
		return origin.name
	}
	aMap := SliceToMap(aSlice, aFunc)
	assert.Equal(t, len(aMap), 3)
	assert.Equal(t, aMap["sourcelliu:2"].id, 2)

	aIntFunc := func(origin *A) int {
		return origin.id
	}
	aIntMap := SliceToMap(aSlice, aIntFunc)
	assert.Equal(t, len(aIntMap), 3)
	assert.Equal(t, aIntMap[2].name, "sourcelliu:2")
}

type A struct {
	id   int
	name string
}
