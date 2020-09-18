# types

```go
import "bitbucket.org/veeafr/utils.go/types"
```


```go
// Check if its a slice.
func IsSlice(v interface{}) bool

// Check if its a map.
func IsMap(v interface{}) bool

// Check if an interface{} is a map and contains the provided key.
func HasKey(obj interface{}, key string) bool 

// Return an interface{} key value if its a map
func GetKey(obj interface{}, key string) interface{}

// Convert any types to string
func ToString(any interface{}) string
```
