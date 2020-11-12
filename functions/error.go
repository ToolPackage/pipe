package functions

import "errors"

var (
	NotEnoughParameterError     = errors.New("not enough parameter")
	InvalidTypeOfParameterError = errors.New("invalid type of parameter error")
	IllegalParameterError       = errors.New("illegal parameter")
)
