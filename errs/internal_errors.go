package errs

import "github.com/ansel1/merry"

var (
	NotFound = merry.New("the specified resource was not found or insufficient permissions")
)
