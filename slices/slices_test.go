package slices

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToMap(t *testing.T) {
	strSlice := []string{"4", "2", "3"}
	strMap := SliceToMap(strSlice)
	assert.Equal(t, len(strMap), 3)
	assert.Equal(t, strMap["2"], true)

	floatSlice := []float64{0.001, 1.234, 4.567}
	floatMap := SliceToMap(floatSlice)
	assert.Equal(t, len(floatMap), 3)
	assert.Equal(t, floatMap[1.234], true)
	assert.Equal(t, floatMap[1.334], false)
}

func TestSliceToMapFunc(t *testing.T) {
	strSlice := []string{"4", "2", "3"}
	strFunc := func(origin string) string {
		return origin
	}
	strMap := SliceToMapFunc(strSlice, strFunc)
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
	aMap := SliceToMapFunc(aSlice, aFunc)
	assert.Equal(t, len(aMap), 3)
	assert.Equal(t, aMap["sourcelliu:2"].id, 2)

	aIntFunc := func(origin *A) int {
		return origin.id
	}
	aIntMap := SliceToMapFunc(aSlice, aIntFunc)
	assert.Equal(t, len(aIntMap), 3)
	assert.Equal(t, aIntMap[2].name, "sourcelliu:2")

	var aNilSlice []*A
	var aNilFunc func(origin *A) int
	aNilMap := SliceToMapFunc(aNilSlice, aNilFunc)
	assert.Nil(t, aNilMap)
	aNilMap = SliceToMapFunc(aSlice, aNilFunc)
	assert.Nil(t, aNilMap)
}

type A struct {
	id   int
	name string
}
