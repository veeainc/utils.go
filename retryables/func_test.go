package retryables

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var count = 0

func work() (interface{}, error) {
	count++
	if count < 4 {
		return 0, errors.New("fail")
	} else {
		return 42, nil
	}
}

func TestRetryPromise(t *testing.T) {
	count = 0

	res, err := Retry(5, 10 * time.Millisecond, work)
	assert.NoError(t, err)
	assert.Equal(t, 42, res)
}
