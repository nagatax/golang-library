package utils

type ExitStatus int

const (
	OK ExitStatus = iota
	ERR
)

func (obj *ExitStatus) Names() []string {
	return []string{
		"OK",
		"ERROR",
	}
}

func (obj *ExitStatus) String() string {
	return (obj.Names())[*obj]
}
