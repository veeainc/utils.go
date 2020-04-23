# interfaces

```go
import "bitbucket.org/veeafr/utils.go/interfaces"
```

## Example

```go
func takePointer(obj interface{}) {
	interfaces.EnsurePointer(obj)
}

```