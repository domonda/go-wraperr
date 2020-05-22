package wraperr

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

func funcA(ctx context.Context, i int, s string) (err error) {
	defer WithFuncParams(&err, ctx, i, s)

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
	err := funcA(context.Background(), 666, "Hello World!")
	errStr := err.Error()
	errStr = strings.ReplaceAll(errStr, basePath(), "")
	fmt.Println(errStr)

	// Output:
	// error in funcC
	// github.com/domonda/go-wraperr.funcC()
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:25
	// github.com/domonda/go-wraperr.funcB([]string{"Hello World!", "X\nX"})
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:19
	// github.com/domonda/go-wraperr.funcA(Context{Err:<nil>}, 666, "Hello World!")
	//     github.com/domonda/go-wraperr/withfuncparams_test.go:13
}
