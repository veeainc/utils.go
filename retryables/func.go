package retryables

import (
	"bitbucket.org/veeafr/utils.go/logging"
	"fmt"
	"github.com/pkg/errors"
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

func Retry(tries uint, delay time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	var ret interface{}
	var err, lastError error
	for ret, lastError = fn(); lastError != nil && tries > 1; tries = tries - 1 {
		if lastError != nil {
			if err == nil {
				err = errors.Wrap(lastError, fmt.Sprintf("failed on try: %d", tries))
			} else {
				err = errors.Wrap(err, fmt.Sprintf("failed on try: %d", tries))
			}
		}
		_retryLog.WithFields(logging.LF{
			"fn": getFunctionName(fn),
		}).Trace("fn failed, retrying")
		time.Sleep(delay)
		ret, lastError = fn()
	}
	if lastError != nil {
		return nil, errors.Wrap(err, "number of retry exceeded")
	}
	return ret, nil
}
