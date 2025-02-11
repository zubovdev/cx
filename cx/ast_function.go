package cxcore

// MakeFunction creates an empty function.
// Later, parameters and contents can be added.
//
func MakeFunction(name string, fileName string, fileLine int) *CXFunction {
	return &CXFunction{
		Name:     name,
		FileName: fileName,
		FileLine: fileLine,
	}
}

// MakeNativeFunction creates a native function such as i32.add()
// not used
func MakeNativeFunction(opCode int, inputs []*CXArgument, outputs []*CXArgument) *CXFunction {
	fn := &CXFunction{
		IsNative: true,
		OpCode:   opCode,
		IntCode:  -1,
		Version:  1,
	}

	offset := 0
	for _, inp := range inputs {
		inp.Offset = offset
		offset += GetSize(inp)
		fn.Inputs = append(fn.Inputs, inp)
	}
	for _, out := range outputs {
		fn.Outputs = append(fn.Outputs, out)
		out.Offset = offset
		offset += GetSize(out)
	}

	return fn
}

//not used
func MakeNativeFunctionV2(opCode int, inputs []*CXArgument, outputs []*CXArgument) *CXFunction {
	fn := &CXFunction{
		IsNative: true,
		OpCode:   opCode,
		IntCode:  -1,
		Version:  2,
	}

	offset := 0
	for _, inp := range inputs {
		inp.Offset = offset
		offset += GetSize(inp)
		fn.Inputs = append(fn.Inputs, inp)
	}
	for _, out := range outputs {
		fn.Outputs = append(fn.Outputs, out)
		out.Offset = offset
		offset += GetSize(out)
	}

	return fn
}

// ----------------------------------------------------------------
//                             `CXFunction` Getters

// GetExpressions is not used
func (fn *CXFunction) GetExpressions() []*CXExpression {
	return fn.Expressions
}

// GetExpressionByLabel
func (fn *CXFunction) GetExpressionByLabel(lbl string) *CXExpression {
	if fn.Expressions == nil {
		return nil
	}
	for _, expr := range fn.Expressions {
		if expr.Label == lbl {
			return expr
		}
	}
	return nil
}

// GetExpressionByLine ...
func (fn *CXFunction) GetExpressionByLine(line int) *CXExpression {
	if fn.Expressions != nil {
		if line <= len(fn.Expressions) {
			return fn.Expressions[line]
		}
		return nil

	}
	return nil

}

// GetCurrentExpression ...
func (fn *CXFunction) GetCurrentExpression() *CXExpression {
	if fn.CurrentExpression != nil {
		return fn.CurrentExpression
	} else if fn.Expressions != nil {
		return fn.Expressions[0]
	} else {
		return nil
	}
}

// ----------------------------------------------------------------
//                     `CXFunction` Member handling

// AddInput ...
func (fn *CXFunction) AddInput(param *CXArgument) *CXFunction {
	found := false
	for _, inp := range fn.Inputs {
		if inp.Name == param.Name {
			found = true
			break
		}
	}
	if !found {
		fn.Inputs = append(fn.Inputs, param)
	}

	return fn
}

// RemoveInput ...
func (fn *CXFunction) RemoveInput(inpName string) {
	if len(fn.Inputs) > 0 {
		lenInps := len(fn.Inputs)
		for i, inp := range fn.Inputs {
			if inp.Name == inpName {
				if i == lenInps {
					fn.Inputs = fn.Inputs[:len(fn.Inputs)-1]
				} else {
					fn.Inputs = append(fn.Inputs[:i], fn.Inputs[i+1:]...)
				}
				break
			}
		}
	}
}

// AddOutput ...
func (fn *CXFunction) AddOutput(param *CXArgument) *CXFunction {
	found := false
	for _, out := range fn.Outputs {
		if out.Name == param.Name {
			found = true
			break
		}
	}
	if !found {
		fn.Outputs = append(fn.Outputs, param)
	}

	param.Package = fn.Package

	return fn
}

// RemoveOutput ...
func (fn *CXFunction) RemoveOutput(outName string) {
	if len(fn.Outputs) > 0 {
		lenOuts := len(fn.Outputs)
		for i, out := range fn.Outputs {
			if out.Name == outName {
				if i == lenOuts {
					fn.Outputs = fn.Outputs[:len(fn.Outputs)-1]
				} else {
					fn.Outputs = append(fn.Outputs[:i], fn.Outputs[i+1:]...)
				}
				break
			}
		}
	}
}

// AddExpression ...
func (fn *CXFunction) AddExpression(expr *CXExpression) *CXFunction {
	// expr.Program = fn.Program
	expr.Package = fn.Package
	expr.Function = fn
	fn.Expressions = append(fn.Expressions, expr)
	fn.CurrentExpression = expr
	fn.Length++
	return fn
}

// RemoveExpression ...
func (fn *CXFunction) RemoveExpression(line int) {
	if len(fn.Expressions) > 0 {
		lenExprs := len(fn.Expressions)
		if line >= lenExprs-1 || line < 0 {
			fn.Expressions = fn.Expressions[:len(fn.Expressions)-1]
		} else {
			fn.Expressions = append(fn.Expressions[:line], fn.Expressions[line+1:]...)
		}
		// for i, expr := range fn.Expressions {
		// 	expr.Index = i
		// }
	}
}

// ----------------------------------------------------------------
//                             `CXFunction` Selectors

// MakeExpression ...
func MakeExpression(op *CXFunction, fileName string, fileLine int) *CXExpression {
	return &CXExpression{
		Operator: op,
		FileLine: fileLine,
		FileName: fileName}
}
