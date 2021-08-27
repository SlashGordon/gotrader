package vector

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func TestDiff(t *testing.T) {
	diff, err := Diff([]float64{1.0}, []float64{1.0, 2.2})
	assert.Nil(t, diff)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrShape.string)
	diff, err = Diff([]float64{1.0, 2.0, 3.0}, []float64{1.0, 1.2, 2.0})
	assert.Nil(t, err)
	assert.Equal(t, diff, []float64{0.0, 0.8, 1.0})
}

func TestDiffCopy(t *testing.T) {
	diff, err := DiffCopy([]float64{1.0}, []float64{1.0, 2.2})
	assert.Nil(t, diff)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrShape.string)
	diff, err = DiffCopy([]float64{1.0, 2.0, 3.0}, []float64{1.0, 1.2, 2.0})
	assert.Nil(t, err)
	assert.Equal(t, diff, []float64{0.0, 0.8, 1.0})
}

func BenchmarkDiff(b *testing.B) {
	sliceA := randFloats(0, 299, 100)
	sliceB := randFloats(0, 299, 100)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Diff(sliceA, sliceB)
	}
}

func BenchmarkDiffCopy(b *testing.B) {
	sliceA := randFloats(0, 299, 100)
	sliceB := randFloats(0, 299, 100)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		DiffCopy(sliceA, sliceB)
	}
}
