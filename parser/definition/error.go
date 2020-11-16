package definition

import "errors"

var (
	InvalidFuncParamDefError = errors.New("invalid function parameter definition")

	NotEnoughParameterError         = errors.New("not enough parameter")
	InvalidTypeOfParameterError     = errors.New("invalid type of parameter")
	InvalidUsageOfVairableError     = errors.New("variable cannot receive dict value")
	InvalidConstValueError          = errors.New("invalid const value")
	InvalidNonLabeledParameterError = errors.New("non-labeled parameter follows labeled parameter")
	IllegalParameterError           = errors.New("illegal parameter")

	UpdateImmutableVariableError = errors.New("defined variable cannot be modified")
	UndefinedVariableError       = errors.New("variable must be defined as pipe node before accessing by function")
)
