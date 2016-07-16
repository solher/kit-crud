package client

import "errors"

var (
	ErrNotFound = errors.New("the specified resource was not found or insufficient permissions")
)
