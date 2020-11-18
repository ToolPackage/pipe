package definition

import (
	"errors"
	"fmt"
)

var (
	InvalidFuncParamDefError = errors.New("invalid function parameter definition")

	NotEnoughParameterError         = errors.New("not enough parameter")
	InvalidTypeOfParameterError     = errors.New("invalid type of parameter")
	InvalidUsageOfVariableError     = errors.New("variable cannot receive dict value")
	InvalidConstValueError          = errors.New("invalid const value")
	InvalidNonLabeledParameterError = errors.New("non-labeled parameter follows labeled parameter")

	DuplicateFuncParamNameError = errors.New("duplicated script function parameter name")
)

func UpdateImmutableVariableError(variableName string) error {
	return fmt.Errorf("defined variable cannot be modified: %s", variableName)
}

func UndefinedVariableError(variableName string) error {
	return fmt.Errorf("variable must be defined as pipe node before accessing by function: %s", variableName)
}
