# Package [github.com/cloudengio/go-keychain/secretservice](https://pkg.go.dev/github.com/cloudengio/go-keychain/secretservice?tab=doc)

```go
import github.com/cloudengio/go-keychain/secretservice
```


## Constants
### DefaultCollection
```go
DefaultCollection dbus.ObjectPath = "/org/freedesktop/secrets/aliases/default"

```
DefaultCollection need not necessarily exist in the user's keyring.

### DefaultSessionOpenTimeout
```go
DefaultSessionOpenTimeout = 10 * time.Second

```
DefaultSessionOpenTimeout

### NilFlags
```go
NilFlags = 0

```
NilFlags

### NullPrompt
```go
NullPrompt = "/"

```
NullPrompt

### ReplaceBehaviorDoNotReplace
```go
ReplaceBehaviorDoNotReplace = 0

```
ReplaceBehaviorDoNotReplace

### ReplaceBehaviorReplace
```go
ReplaceBehaviorReplace = 1

```
ReplaceBehaviorReplace

### SecretServiceInterface
```go
SecretServiceInterface = "org.freedesktop.secrets"

```
SecretServiceInterface

### SecretServiceObjectPath
```go
SecretServiceObjectPath dbus.ObjectPath = "/org/freedesktop/secrets"

```
SecretServiceObjectPath



## Functions
### Func NewSecretProperties
```go
func NewSecretProperties(label string, attributes map[string]string) map[string]dbus.Variant
```
NewSecretProperties



## Types
### Type Attributes
```go
type Attributes map[string]string
```
Attributes


### Type AuthenticationMode
```go
type AuthenticationMode string
```
AuthenticationMode

### Constants
### AuthenticationDHAES
```go
AuthenticationDHAES AuthenticationMode = "dh-ietf1024-sha256-aes128-cbc-pkcs7"

```
AuthenticationDHAES

### AuthenticationInsecurePlain
```go
AuthenticationInsecurePlain AuthenticationMode = "plain"

```
AuthenticationInsecurePlain




### Type PromptCompletedResult
```go
type PromptCompletedResult struct {
	Dismissed bool
	Paths     dbus.Variant
}
```
PromptCompletedResult


### Type PromptDismissedError
```go
type PromptDismissedError struct {
	// contains filtered or unexported fields
}
```
PromptDismissedError

### Methods

```go
func (p PromptDismissedError) Error() string
```
Error


```go
func (p PromptDismissedError) Unwrap() error
```
Unwrap




### Type ReplaceBehavior
```go
type ReplaceBehavior int
```
ReplaceBehavior


### Type Secret
```go
type Secret struct {
	Session     dbus.ObjectPath
	Parameters  []byte
	Value       []byte
	ContentType string
}
```
Secret


### Type SecretService
```go
type SecretService struct {
	// contains filtered or unexported fields
}
```
SecretService

### Functions

```go
func NewService() (*SecretService, error)
```
NewService



### Methods

```go
func (s *SecretService) CloseSession(session *Session)
```
CloseSession


```go
func (s *SecretService) CreateItem(collection dbus.ObjectPath, properties map[string]dbus.Variant, secret Secret, replaceBehavior ReplaceBehavior) (item dbus.ObjectPath, err error)
```
CreateItem


```go
func (s *SecretService) DeleteItem(item dbus.ObjectPath) (err error)
```
DeleteItem


```go
func (s *SecretService) GetAttributes(item dbus.ObjectPath) (attributes Attributes, err error)
```
GetAttributes


```go
func (s *SecretService) GetSecret(item dbus.ObjectPath, session Session) (secretPlaintext []byte, err error)
```
GetSecret


```go
func (s *SecretService) LockItems(items []dbus.ObjectPath) (err error)
```
LockItems


```go
func (s *SecretService) Obj(path dbus.ObjectPath) dbus.BusObject
```
Obj


```go
func (s *SecretService) OpenSession(mode AuthenticationMode) (session *Session, err error)
```
OpenSession


```go
func (s *SecretService) PromptAndWait(prompt dbus.ObjectPath) (paths *dbus.Variant, err error)
```
PromptAndWait is NOT thread-safe.


```go
func (s *SecretService) SearchCollection(collection dbus.ObjectPath, attributes Attributes) (items []dbus.ObjectPath, err error)
```
SearchCollection


```go
func (s *SecretService) ServiceObj() dbus.BusObject
```
ServiceObj


```go
func (s *SecretService) SetSessionOpenTimeout(d time.Duration)
```
SetSessionOpenTimeout


```go
func (s *SecretService) Unlock(items []dbus.ObjectPath) (err error)
```
Unlock




### Type Session
```go
type Session struct {
	Mode    AuthenticationMode
	Path    dbus.ObjectPath
	Public  *big.Int
	Private *big.Int
	AESKey  []byte
}
```
Session

### Methods

```go
func (session *Session) NewSecret(secretBytes []byte) (Secret, error)
```
NewSecret







