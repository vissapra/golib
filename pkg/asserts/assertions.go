package assert

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func AssertEquals(t *testing.T, actual, expected interface{}, msg string) {
	if !reflect.DeepEqual(actual, expected) {
		printErrWithCaller(actual, expected, msg, t)
	}
}
func AssertNotEquals(t *testing.T, actual, expected interface{}, msg string) {
	if reflect.DeepEqual(actual, expected) {
		printErrWithCaller(actual, expected, msg, t)
	}
}

//Formats the error message to include the caller's assert location, instead of current Error location
func printErrWithCaller(actual interface{}, expected interface{}, msg string, t *testing.T) {
	message := fmt.Sprintf("actual:%v, expected:%v, msg: %s", actual, expected, msg)
	_, file, line, ok := runtime.Caller(2)
	if ok {
		parts := strings.Split(file, "/")
		fmt.Fprintf(os.Stdout, "%s:%d: %s", parts[len(parts)-1], line, message)
		t.Error()
	} else {
		t.Error(message)
	}
}
