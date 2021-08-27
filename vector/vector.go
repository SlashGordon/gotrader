package vector

type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	ErrShape       = Error{"dimension mismatch"}
	ErrZeroLength  = Error{"zero length in matrix dimension"}
	ErrNegativeDim = Error{"negative dimension"}
	ErrNullPointer = Error{"pointer is nil"}
	ErrLogic       = Error{"Logical error"}
)

type Vector struct {
	vec []float64
}

type VectorPair struct {
	vecA []float64
	vecB []float64
}

func (a *Vector) Diff(b *Vector) ([]float64, error) {
	return Diff(a.vec, b.vec)
}

func (a *VectorPair) Diff() ([]float64, error) {
	return Diff(a.vecA, a.vecB)
}

func Diff(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, ErrShape
	}

	for idx, val := range b {
		a[idx] -= val
	}

	return a, nil
}

func DiffCopy(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, ErrShape
	}

	result := make([]float64, len(a))

	for idx, val := range b {
		result[idx] = a[idx] - val
	}

	return result, nil
}
