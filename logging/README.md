# logging

```go
import "bitbucket.org/veeafr/utils.go/logging"
```

Allow to retrieve configured and named logger.

The logger automatically take the file name as category.

This is based on logrus logger.

## Example

vbus.go
```go
logging.SetLogLevel(logrus.TraceLevel)
_vbusLogger = logging.GetNamedLogger()

_vbusLogger.WithField("msg", "hello world").Info("new message")
```

    [2020-04-23T15:23:22+02:00]  INFO vbus: new message msg=hello world