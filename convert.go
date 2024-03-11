package gohelpers

import (
	"errors"
	"fmt"
	"strconv"
)

// A NumError records a failed conversion.
type Response struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex)
	Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *Response) Error() string {
	return "convert." + e.Func + ": " + e.Err.Error()
}

func (e *Response) Unwrap() error { return e.Err }

func syntaxError(fn string, Err error) *Response {
	return &Response{fn, Err}
}

func Convert(value interface{}, needT string) (response any, err error) {
	switch value.(type) {
	case int:
		return intConvert(value.(int), needT)
	default:
		return nil, syntaxError("", errors.New("no support type"))
	}
}

func intConvert(value int, needT string) (resp any, err error) {
	const fnIntConvert = "IntConvert"

	if needT == "string" {
		s := strconv.Itoa(value)

		fmt.Println(s)

		return s, nil
	}
	return 0, &Response{Err: errors.New("no int convert")}
}
