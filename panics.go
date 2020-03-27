package wraperr

import (
	"errors"
	"fmt"
)

type Logger interface {
	Printf(format string, args ...interface{})
}

func LogPanicWithFuncParams(log Logger, params ...interface{}) {
	p := recover()
	if p == nil {
		return
	}

	err := newWithFuncParamsSkip(1, AsError(p), params...)

	log.Printf("LogPanicWithFuncParams: %s", err.Error())

	panic(p)
}

func RecoverAndLogPanicWithFuncParams(log Logger, params ...interface{}) {
	p := recover()
	if p == nil {
		return
	}

	err := newWithFuncParamsSkip(1, AsError(p), params...)

	log.Printf("RecoverAndLogPanicWithFuncParams: %s", err.Error())
}

func RecoverPanicAsErrorWithFuncParams(resultVar *error, params ...interface{}) {
	p := recover()
	if p == nil {
		return
	}

	err := newWithFuncParamsSkip(1, AsError(p), params...)

	if *resultVar != nil {
		*resultVar = fmt.Errorf("function returning error %s paniced with: %w", *resultVar, err)
	} else {
		*resultVar = err
	}
}

func AsError(val interface{}) error {
	switch x := val.(type) {
	case nil:
		return nil
	case error:
		return x
	case string:
		return errors.New(x)
	case fmt.Stringer:
		return errors.New(x.String())
	default:
		return fmt.Errorf("%+v", val)
	}
}
