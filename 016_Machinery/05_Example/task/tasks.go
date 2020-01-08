package task

const (
	// ADD : add func
	ADD string = "add"

	// MULTIPLY : Multiply func
	MULTIPLY string = "multiply"
)

// Add func
func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

// Multiply func
func Multiply(args ...int64) (int64, error) {
	res := int64(1)
	for _, arg := range args {
		res *= arg
	}
	return res, nil
}
