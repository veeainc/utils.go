# retryables

```go
import "bitbucket.org/veeafr/utils.go/retryables"
```

Some utilities to make a function retryable.

## Example

```go
if ret, err := retry.Retry(3, 800 * time.Millisecond, func() (i interface{}, err error) {
    return checkHubOwnership()
}); err != nil {
    _logger.Info("failed to checkHubOwnership after 3 retry")
} else {
	_logger.Info("success")
}
```