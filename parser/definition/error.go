package definition

import "errors"

var (
	InvalidFuncParamDefError = errors.New("invalid function parameter definition")

	NotEnoughParameterError         = errors.New("not enough parameter")
	InvalidTypeOfParameterError     = errors.New("invalid type of parameter")
	InvalidConstValueError          = errors.New("invalid const value")
	InvalidNonLabeledParameterError = errors.New("non-labeled parameter follows labeled parameter")
	IllegalParameterError           = errors.New("illegal parameter")
)
