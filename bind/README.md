# Package [github.com/cloudengio/go-keychain/bind](https://pkg.go.dev/github.com/cloudengio/go-keychain/bind?tab=doc)

```go
import github.com/cloudengio/go-keychain/bind
```


## Functions
### Func AddGenericPassword
```go
func AddGenericPassword(service string, account string, label string, password string, accessGroup string) error
```
AddGenericPassword adds generic password

### Func DeleteGenericPassword
```go
func DeleteGenericPassword(service string, account string, accessGroup string) error
```
DeleteGenericPassword deletes generic password

### Func GenericPasswordTest
```go
func GenericPasswordTest(t Test, service string, accessGroup string)
```
GenericPasswordTest runs test code for generic password keychain item.
This is here so we can export using gomobile bind and run this method on iOS
simulator and device. Access groups aren't supported in iOS simulator.



## Types
### Type Test
```go
type Test interface {
	Fail(s string)
}
```
Test is a bind interface for the test





