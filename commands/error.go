package commands

import "errors"

var (
	NotEnoughParameterError = errors.New("not enough parameter")
	IllegalParameterError   = errors.New("illegal parameter")
)
