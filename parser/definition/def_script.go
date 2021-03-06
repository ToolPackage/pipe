package definition

import (
	"github.com/ToolPackage/pipe/util"
	"io"
	"strconv"
)

// PipeScript
type PipeScript struct {
	Funcs []CompactFunction
}

func (p *PipeScript) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("PipeScript [")
	builder.IncIndent()

	for _, funcDef := range p.Funcs {
		builder.WriteMultiLine(funcDef.String(), true)
	}

	builder.DecIndent()
	builder.WriteWithIndent("]")

	return builder.String()
}

// CompactFunction
type CompactFunction struct {
	Name     string
	Md5      string
	FilePath string
	Params   FunctionParameters
	Pipes    *MultiPipe
}

func (c *CompactFunction) Exec(params Parameters, in io.Reader, out io.Writer) error {
	for idx := range c.Params {
		paramDef := c.Params[idx]
		param, ok := params.GetParameter(paramDef.Name, idx)
		if !ok {
			panic("Oops!")
		}
		v, _ := c.Pipes.Variables[paramDef.Name]
		v.Assign(param.Value.Get())
		// TODO: ???
		// use variable to accept dict value?
	}
	return nil
}

func (c *CompactFunction) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("CompactFunction {")

	builder.IncIndent()
	builder.WriteLine("Name: ", c.Name)

	builder.WriteLine("Params: [")
	builder.IncIndent()
	for _, paramDef := range c.Params {
		builder.WriteMultiLine(paramDef.String(), true)
	}
	builder.DecIndent()
	builder.WriteLine("]")

	builder.WriteWithIndent("Pipes: ")
	builder.WriteMultiLine(c.Pipes.String(), false)

	builder.DecIndent()
	builder.WriteWithIndent("}")

	return builder.String()
}

type ParameterDefinition struct {
	Name       string
	Type       ValueType
	Optional   bool
	ConstValue []interface{}
}

func (p *ParameterDefinition) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteWithIndent("FuncParamDef { Name: ", p.Name,
		", Optional: ", strconv.FormatBool(p.Optional),
		", Type: ", ValueType(p.Type).String(),
		" }")
	return builder.String()
}

// FunctionDefinition, used by registry

type FunctionDefinition struct {
	Name            string
	Builtin         bool
	Handler         FunctionHandler
	ParamConstraint FunctionParameters
}

type FunctionParameters []ParameterDefinition

func (fp FunctionParameters) Validate(params Parameters) error {
	// 1.validate labeled params
	idx := 0
	// skip all non-labeled parameters
	for ; idx < len(params) && !params[idx].Labeled(); idx++ {
	}
	// check all following parameters are labeled
	for ; idx < len(params); idx++ {
		if !params[idx].Labeled() {
			return InvalidNonLabeledParameterError
		}
	}

	paramSz := params.Len()
	for idx, paramDef := range fp {
		// 2.validate parameters size
		if idx == paramSz {
			for ; idx < len(fp); idx++ {
				if !paramDef.Optional {
					return NotEnoughParameterError
				}
			}
			return nil
		}
		param, ok := params.GetParameter(paramDef.Name, idx)
		if !ok {
			return nil
		}
		// 3.validate parameter type
		if param.Value.Type() == ReferenceValue {
			// if parameter is a variable, check if function param type is dict
			if paramDef.Type == DictValue {
				return InvalidUsageOfVariableError
			}
			// set reference value's type which will be used to convert variable text value to typed value
			refValue := param.Value.(*ReferenceParameterValue)
			refValue.RefType = paramDef.Type
		} else if param.Value.Type() != paramDef.Type {
			return InvalidTypeOfParameterError
		}
		// 4.validate parameter value
		if len(paramDef.ConstValue) > 0 && !util.SliceContains(paramDef.ConstValue, param.Value.Get()) {
			return InvalidConstValueError
		}
	}
	return nil
}
