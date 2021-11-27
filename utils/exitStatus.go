// Package utils exitStatus
package utils

// ExitStatus is exit code
type ExitStatus int

const (
	// OK is exit code 0
	OK ExitStatus = iota
	// ERR is exit code 1
	ERR
)

// Names returns array which is included exit status
func (obj *ExitStatus) Names() []string {
	return []string{
		"OK",
		"ERROR",
	}
}

// String returns string value of exit code
func (obj *ExitStatus) String() string {
	return (obj.Names())[*obj]
}
