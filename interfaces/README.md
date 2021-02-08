# interfaces

```go
import "github.com/veeainc/utils.go/interfaces"
```

## Example

```go
func takePointer(obj interface{}) {
	interfaces.EnsurePointer(obj)
}

```