package helpers

// ConstError is used to define an error as a constant
type ConstError string

var _ error = ConstError("")

// Error returns the error as a string
func (err ConstError) Error() string {
	return string(err)
}
