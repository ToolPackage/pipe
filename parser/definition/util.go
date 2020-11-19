package definition

import "strings"

// Define function utils

func DefFuncs(funcDefList ...*FunctionDefinition) []*FunctionDefinition {
	return funcDefList
}

func DefLibFunc(name string, handler FunctionHandler, paramConstraint FuncParamConstraint) *FunctionDefinition {
	return &FunctionDefinition{Name: name, Builtin: false, Handler: handler, ParamConstraint: paramConstraint}
}

func DefBuiltinFunc(name string, handler FunctionHandler, paramConstraint FuncParamConstraint) *FunctionDefinition {
	return &FunctionDefinition{Name: name, Builtin: true, Handler: handler, ParamConstraint: paramConstraint}
}

func DefParams(paramDefList ...ParameterDefinition) FuncParamConstraint {
	return paramDefList
}

func DefParam(paramType ValueType, label string, optional bool, constValue ...interface{}) ParameterDefinition {
	label = strings.Trim(label, " \t\n\r")
	if len(label) == 0 {
		panic(InvalidFuncParamDefError)
	}
	return ParameterDefinition{Type: paramType, Name: label, Optional: optional, ConstValue: constValue}
}
