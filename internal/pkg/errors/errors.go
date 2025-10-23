package errors

import "fmt"

type Error struct {
	Op  string
	Err error
	Msg string
}

func (e *Error) Error() string {
	switch {
	case e.Op != "" && e.Msg != "" && e.Err != nil:
		return fmt.Sprintf("%s: %s: %v", e.Op, e.Msg, e.Err)
	case e.Op != "" && e.Msg != "":
		return fmt.Sprintf("%s: %s", e.Op, e.Msg)
	case e.Op != "" && e.Err != nil:
		return fmt.Sprintf("%s: %v", e.Op, e.Err)
	case e.Msg != "" && e.Err != nil:
		return fmt.Sprintf("%s: %v", e.Msg, e.Err)
	case e.Msg != "":
		return e.Msg
	case e.Err != nil:
		return e.Err.Error()
	default:
		return "unknown error"
	}
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func New(op, msg string) error {
	return &Error{Op: op, Msg: msg}
}

func Wrap(op string, err error, msg string) error {
	if err == nil {
		return nil
	}
	return &Error{Op: op, Err: err, Msg: msg}
}
