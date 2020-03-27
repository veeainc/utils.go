package utils

import (
	"bitbucket.org/boolangery/corelib.go/logging"
	"errors"
	"reflect"
	"runtime"
	"time"
)

var _retryLog = logging.GetNamedLogger()

type RetryableSimple = func() error
type Retryable = func() (interface{}, error)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Retry(tries int, delay time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	if tries == 0 {
		return nil, errors.New("number of retry exceeded")
	}

	if ret, err := fn(); err != nil {
		_retryLog.WithFields(logging.LF{
			"fn": getFunctionName(fn),
			"error": err,
		}).Trace("fn failed, retrying")
		return Retry(tries -1, delay, fn)
	} else {
		return ret, nil
	}
}
