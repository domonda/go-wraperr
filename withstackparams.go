package wraperr

/*
Call argument parameters are available on the stack,
but in a platform dependent packed format and not directly accessible
via package runtime.
Could only be parsed from runtime.Stack() result text.

See https://www.ardanlabs.com/blog/2015/01/stack-traces-in-go.html
*/

func newWithFuncParams(err error, params ...interface{}) *withStackParams {
	switch w := err.(type) {
	case callStackParamsProvider:
		// OK, wrap the wrapped
	case callStackProvider:
		// Already wrapped with stack,
		// replace wrapper withStackParams
		return &withStackParams{
			withStack: withStack{
				err:   w.Unwrap(),
				stack: w.CallStack(),
			},
			params: params,
		}
	}

	return &withStackParams{
		withStack: withStack{
			err:   err,
			stack: callStack(2),
		},
		params: params,
	}
}

func WithFuncParams(resultVar *error, params ...interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, params...)
	}
}

func With0FuncParams(resultVar *error) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar)
	}
}

func With1FuncParams(resultVar *error, p0 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0)
	}
}

func With2FuncParams(resultVar *error, p0, p1 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1)
	}
}

func With3FuncParams(resultVar *error, p0, p1, p2 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2)
	}
}

func With4FuncParams(resultVar *error, p0, p1, p2, p3 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3)
	}
}

func With5FuncParams(resultVar *error, p0, p1, p2, p3, p4 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4)
	}
}

func With6FuncParams(resultVar *error, p0, p1, p2, p3, p4, p5 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4, p5)
	}
}

func With7FuncParams(resultVar *error, p0, p1, p2, p3, p4, p5, p6 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4, p5, p6)
	}
}

func With8FuncParams(resultVar *error, p0, p1, p2, p3, p4, p5, p6, p7 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func With9FuncParams(resultVar *error, p0, p1, p2, p3, p4, p5, p6, p7, p8 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func With10FuncParams(resultVar *error, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 interface{}) {
	if *resultVar != nil {
		*resultVar = newWithFuncParams(*resultVar, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

type callStackParamsProvider interface {
	CallStackParams() ([]uintptr, []interface{})
}

type withStackParams struct {
	withStack

	params []interface{}
}

func (w *withStackParams) Error() string {
	return formatError(w)
}

func (w *withStackParams) CallStackParams() ([]uintptr, []interface{}) {
	return w.stack, w.params
}
