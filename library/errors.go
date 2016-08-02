package library

import (
	"errors"
	"strings"

	"github.com/facebookgo/stackerr"
	stdopentracing "github.com/opentracing/opentracing-go"
)

var (
	ErrNotFound = errors.New("the specified resource was not found or insufficient permissions")
)

func Unexpected(err error) error {
	return &UnexpectedError{
		err: *stackerr.WrapSkip(err, 1).(*stackerr.Error),
	}
}

type UnexpectedError struct {
	err stackerr.Error
}

func (e *UnexpectedError) Error() string {
	return e.err.Underlying().Error()
}

func (e *UnexpectedError) Location() string {
	s := strings.SplitN(e.err.MultiStack().String(), "\n", 2)
	return s[0]
}

func (e *UnexpectedError) Stacktrace() string {
	return e.err.MultiStack().String()
}

func TraceIfUnexpectedError(span stdopentracing.Span, err error) {
	if err, ok := err.(*UnexpectedError); ok {
		span.SetTag("location", err.Location())
		span.SetTag("stacktrace", err.Stacktrace())
		span.SetTag("error", err.Error())
	}
}
