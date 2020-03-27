package wraperr

import (
	"fmt"
	"runtime"
	"strings"
)

func funcA(i int, s string) (err error) {
	defer WithFuncParams(&err, i, s)

	return funcB(s, "X\nX")
}

func funcB(s ...string) (err error) {
	defer WithFuncParams(&err, s)

	return funcC()
}

func funcC() (err error) {
	defer WithFuncParams(&err)

	return New("error in funcC")
}

func basePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file[:strings.Index(file, "github.com")]
}

func ExampleWithFuncParams() {
	err := funcA(666, "Hello World!")
	errStr := err.Error()
	errStr = strings.ReplaceAll(errStr, basePath(), "")
	fmt.Println(errStr)

	// Output:
	// error in funcC
	// github.com/domonda/go-wraperr.funcC()
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:24
	// github.com/domonda/go-wraperr.funcB([]string{"Hello World!", "X\nX"})
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:18
	// github.com/domonda/go-wraperr.funcA(666, "Hello World!")
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:12
}
