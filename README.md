# Package [github.com/cloudengio/go-keychain](https://pkg.go.dev/github.com/cloudengio/go-keychain?tab=doc)

```go
import github.com/cloudengio/go-keychain
```

This file is an addition made in the cloudengio fork and has no upstream
counterpart. It reports the accessibility that kSecAttrAccessible records
for an item, which QueryItem does not: QueryResult has no field for it and
convertResult ignores the key.

It is deliberately kept to itself, rather than adding a field to QueryResult
and a case to convertResult, so that merging upstream changes touches
nothing of it: every declaration here is new, and the only upstream code it
depends on is called, not edited. The build tag matches macos.go, which is
where the accessibilities this maps back from are defined.

## Constants
### SynchronizableDefault, SynchronizableAny, SynchronizableYes, SynchronizableNo
```go
// SynchronizableDefault is the default setting
SynchronizableDefault Synchronizable = 0
// SynchronizableAny is for kSecAttrSynchronizableAny
SynchronizableAny = 1
// SynchronizableYes enables synchronization
SynchronizableYes = 2
// SynchronizableNo disables synchronization
SynchronizableNo = 3

```

### MatchLimitDefault, MatchLimitOne, MatchLimitAll
```go
// MatchLimitDefault is the default
MatchLimitDefault MatchLimit = 0
// MatchLimitOne limits to one result
MatchLimitOne = 1
// MatchLimitAll is no limit
MatchLimitAll = 2

```



## Variables
### ErrorUnimplemented, ErrorParam, ErrorAllocate, ErrorNotAvailable, ErrorAuthFailed, ErrorDuplicateItem, ErrorItemNotFound, ErrorInteractionNotAllowed, ErrorDecode, ErrorNoSuchKeychain, ErrorNoAccessForItem, ErrorReadOnly, ErrorInvalidKeychain, ErrorDuplicateKeyChain, ErrorWrongVersion, ErrorReadonlyAttribute, ErrorInvalidSearchRef, ErrorInvalidItemRef, ErrorDataNotAvailable, ErrorDataNotModifiable, ErrorInvalidOwnerEdit, ErrorUserCanceled
```go
// ErrorUnimplemented corresponds to errSecUnimplemented result code
ErrorUnimplemented = Error(C.errSecUnimplemented)
// ErrorParam corresponds to errSecParam result code
ErrorParam = Error(C.errSecParam)
// ErrorAllocate corresponds to errSecAllocate result code
ErrorAllocate = Error(C.errSecAllocate)
// ErrorNotAvailable corresponds to errSecNotAvailable result code
ErrorNotAvailable = Error(C.errSecNotAvailable)
// ErrorAuthFailed corresponds to errSecAuthFailed result code
ErrorAuthFailed = Error(C.errSecAuthFailed)
// ErrorDuplicateItem corresponds to errSecDuplicateItem result code
ErrorDuplicateItem = Error(C.errSecDuplicateItem)
// ErrorItemNotFound corresponds to errSecItemNotFound result code
ErrorItemNotFound = Error(C.errSecItemNotFound)
// ErrorInteractionNotAllowed corresponds to errSecInteractionNotAllowed result code
ErrorInteractionNotAllowed = Error(C.errSecInteractionNotAllowed)
// ErrorDecode corresponds to errSecDecode result code
ErrorDecode = Error(C.errSecDecode)
// ErrorNoSuchKeychain corresponds to errSecNoSuchKeychain result code
ErrorNoSuchKeychain = Error(C.errSecNoSuchKeychain)
// ErrorNoAccessForItem corresponds to errSecNoAccessForItem result code
ErrorNoAccessForItem = Error(C.errSecNoAccessForItem)
// ErrorReadOnly corresponds to errSecReadOnly result code
ErrorReadOnly = Error(C.errSecReadOnly)
// ErrorInvalidKeychain corresponds to errSecInvalidKeychain result code
ErrorInvalidKeychain = Error(C.errSecInvalidKeychain)
// ErrorDuplicateKeyChain corresponds to errSecDuplicateKeychain result code
ErrorDuplicateKeyChain = Error(C.errSecDuplicateKeychain)
// ErrorWrongVersion corresponds to errSecWrongSecVersion result code
ErrorWrongVersion = Error(C.errSecWrongSecVersion)
// ErrorReadonlyAttribute corresponds to errSecReadOnlyAttr result code
ErrorReadonlyAttribute = Error(C.errSecReadOnlyAttr)
// ErrorInvalidSearchRef corresponds to errSecInvalidSearchRef result code
ErrorInvalidSearchRef = Error(C.errSecInvalidSearchRef)
// ErrorInvalidItemRef corresponds to errSecInvalidItemRef result code
ErrorInvalidItemRef = Error(C.errSecInvalidItemRef)
// ErrorDataNotAvailable corresponds to errSecDataNotAvailable result code
ErrorDataNotAvailable = Error(C.errSecDataNotAvailable)
// ErrorDataNotModifiable corresponds to errSecDataNotModifiable result code
ErrorDataNotModifiable = Error(C.errSecDataNotModifiable)
// ErrorInvalidOwnerEdit corresponds to errSecInvalidOwnerEdit result code
ErrorInvalidOwnerEdit = Error(C.errSecInvalidOwnerEdit)
// ErrorUserCanceled corresponds to errSecUserCanceled result code
ErrorUserCanceled = Error(C.errSecUserCanceled)

```

### ServiceKey, ServerKey, ProtocolKey, AuthenticationTypeKey, PortKey, PathKey, LabelKey, AccountKey, AccessGroupKey, DataKey, DescriptionKey, CommentKey, CreationDateKey, ModificationDateKey, DataProtectionKeyChainKey
```go
// ServiceKey is for kSecAttrService
ServiceKey = attrKey(C.CFTypeRef(C.kSecAttrService))
// ServerKey is for kSecAttrServer
ServerKey = attrKey(C.CFTypeRef(C.kSecAttrServer))
// ProtocolKey is for kSecAttrProtocol
ProtocolKey = attrKey(C.CFTypeRef(C.kSecAttrProtocol))
// AuthenticationTypeKey is for kSecAttrAuthenticationType
AuthenticationTypeKey = attrKey(C.CFTypeRef(C.kSecAttrAuthenticationType))
// PortKey is for kSecAttrPort
PortKey = attrKey(C.CFTypeRef(C.kSecAttrPort))
// PathKey is for kSecAttrPath
PathKey = attrKey(C.CFTypeRef(C.kSecAttrPath))
// LabelKey is for kSecAttrLabel
LabelKey = attrKey(C.CFTypeRef(C.kSecAttrLabel))
// AccountKey is for kSecAttrAccount
AccountKey = attrKey(C.CFTypeRef(C.kSecAttrAccount))
// AccessGroupKey is for kSecAttrAccessGroup
AccessGroupKey = attrKey(C.CFTypeRef(C.kSecAttrAccessGroup))
// DataKey is for kSecValueData
DataKey = attrKey(C.CFTypeRef(C.kSecValueData))
// DescriptionKey is for kSecAttrDescription
DescriptionKey = attrKey(C.CFTypeRef(C.kSecAttrDescription))
// CommentKey is for kSecAttrComment
CommentKey = attrKey(C.CFTypeRef(C.kSecAttrComment))
// CreationDateKey is for kSecAttrCreationDate
CreationDateKey = attrKey(C.CFTypeRef(C.kSecAttrCreationDate))
// ModificationDateKey is for kSecAttrModificationDate
ModificationDateKey = attrKey(C.CFTypeRef(C.kSecAttrModificationDate))
DataProtectionKeyChainKey = attrKey(C.CFTypeRef(C.kSecUseDataProtectionKeychain))

```

### AccessibleKey
```go
AccessibleKey = attrKey(C.CFTypeRef(C.kSecAttrAccessible))

```
AccessibleKey is key for kSecAttrAccessible

### MatchLimitKey
```go
MatchLimitKey = attrKey(C.CFTypeRef(C.kSecMatchLimit))

```
MatchLimitKey is key type for MatchLimit

### ReturnAttributesKey
```go
ReturnAttributesKey = attrKey(C.CFTypeRef(C.kSecReturnAttributes))

```
ReturnAttributesKey is key type for kSecReturnAttributes

### ReturnDataKey
```go
ReturnDataKey = attrKey(C.CFTypeRef(C.kSecReturnData))

```
ReturnDataKey is key type for kSecReturnData

### ReturnRefKey
```go
ReturnRefKey = attrKey(C.CFTypeRef(C.kSecReturnRef))

```
ReturnRefKey is key type for kSecReturnRef

### SecClassKey
```go
SecClassKey = attrKey(C.CFTypeRef(C.kSecClass))

```
SecClassKey is the key type for SecClass

### SynchronizableKey
```go
SynchronizableKey = attrKey(C.CFTypeRef(C.kSecAttrSynchronizable))

```
SynchronizableKey is the key type for Synchronizable



## Functions
### Func AddItem
```go
func AddItem(item Item) error
```
AddItem adds a Item to a Keychain

### Func ArrayToCFArray
```go
func ArrayToCFArray(a []C.CFTypeRef) C.CFArrayRef
```
ArrayToCFArray will return a CFArrayRef and if non-nil, must be released
with Release(ref).

### Func BytesToCFData
```go
func BytesToCFData(b []byte) (C.CFDataRef, error)
```
BytesToCFData will return a CFDataRef and if non-nil, must be released with
Release(ref).

### Func CFArrayToArray
```go
func CFArrayToArray(cfArray C.CFArrayRef) (a []C.CFTypeRef)
```
CFArrayToArray converts a CFArrayRef to an array of CFTypes.

### Func CFDataToBytes
```go
func CFDataToBytes(cfData C.CFDataRef) ([]byte, error)
```
CFDataToBytes converts CFData to bytes.

### Func CFDateToTime
```go
func CFDateToTime(d C.CFDateRef) time.Time
```
CFDateToTime will convert the given CFDateRef to a time.Time.

### Func CFDictionaryToMap
```go
func CFDictionaryToMap(cfDict C.CFDictionaryRef) (m map[C.CFTypeRef]C.CFTypeRef)
```
CFDictionaryToMap converts CFDictionaryRef to a map.

### Func CFNumberToInterface
```go
func CFNumberToInterface(cfNumber C.CFNumberRef) (interface{}, error)
```
CFNumberToInterface converts the CFNumberRef to the most appropriate numeric
type. This code is from github.com/kballard/go-osx-plist.

### Func CFStringToString
```go
func CFStringToString(s C.CFStringRef) string
```
CFStringToString converts a CFStringRef to a string.

### Func CFTypeDescription
```go
func CFTypeDescription(ref C.CFTypeRef) string
```
CFTypeDescription returns type string for CFTypeRef.

### Func Convert
```go
func Convert(ref C.CFTypeRef) (interface{}, error)
```
Convert converts a CFTypeRef to a go instance.

### Func ConvertCFDictionary
```go
func ConvertCFDictionary(d C.CFDictionaryRef) (map[interface{}]interface{}, error)
```
ConvertCFDictionary converts a CFDictionary to map (deep).

### Func ConvertMapToCFDictionary
```go
func ConvertMapToCFDictionary(attr map[string]interface{}) (C.CFDictionaryRef, error)
```
ConvertMapToCFDictionary converts a map to a CFDictionary and if non-nil,
must be released with Release(ref).

### Func DeleteGenericPasswordItem
```go
func DeleteGenericPasswordItem(service string, account string) error
```
DeleteGenericPasswordItem removes a generic password item.

### Func DeleteItem
```go
func DeleteItem(item Item) error
```
DeleteItem removes a Item

### Func GetAccountsForService
```go
func GetAccountsForService(service string) ([]string, error)
```
GetAccountsForService is deprecated

### Func GetGenericPassword
```go
func GetGenericPassword(service string, account string, label string, accessGroup string) ([]byte, error)
```
GetGenericPassword returns password data for service and account. This is a
convenience method. If item is not found returns nil, nil.

### Func GetGenericPasswordAccounts
```go
func GetGenericPasswordAccounts(service string) ([]string, error)
```
GetGenericPasswordAccounts returns generic password accounts for service.
This is a convenience method.

### Func Int32ToCFNumber
```go
func Int32ToCFNumber(u int32) C.CFNumberRef
```
Int32ToCFNumber will return a CFNumberRef, must be released with
Release(ref).

### Func MapToCFDictionary
```go
func MapToCFDictionary(m map[C.CFTypeRef]C.CFTypeRef) (C.CFDictionaryRef, error)
```
MapToCFDictionary will return a CFDictionaryRef and if non-nil, must be
released with Release(ref).

### Func QueryItemRef
```go
func QueryItemRef(item Item) (C.CFTypeRef, error)
```
QueryItemRef returns query result as CFTypeRef. You must release it when you
are done.

### Func RandBytes
```go
func RandBytes(length int) ([]byte, error)
```
RandBytes returns random bytes of length

### Func RandomID
```go
func RandomID(prefix string) (string, error)
```
RandomID returns random ID (base32) string with prefix, using 256 bits as
recommended by tptacek: https://gist.github.com/tqbf/be58d2d39690c3b366ad

### Func Release
```go
func Release(ref C.CFTypeRef)
```
Release releases memory pointed to by a CFTypeRef.

### Func StringToCFString
```go
func StringToCFString(s string) (C.CFStringRef, error)
```
StringToCFString will return a CFStringRef and if non-nil, must be released
with Release(ref).

### Func TimeToCFDate
```go
func TimeToCFDate(t time.Time) C.CFDateRef
```
TimeToCFDate will convert the given time.Time to a CFDateRef, which must be
released with Release(ref).

### Func UpdateItem
```go
func UpdateItem(queryItem Item, updateItem Item) error
```
UpdateItem updates the queryItem with the parameters from updateItem



## Types
### Type Accessible
```go
type Accessible int
```
Accessible is the items accessibility

### Constants
### AccessibleDefault, AccessibleWhenUnlocked, AccessibleAfterFirstUnlock, AccessibleAlways, AccessibleWhenPasscodeSetThisDeviceOnly, AccessibleWhenUnlockedThisDeviceOnly, AccessibleAfterFirstUnlockThisDeviceOnly, AccessibleAccessibleAlwaysThisDeviceOnly
```go
// AccessibleDefault is the default
AccessibleDefault Accessible = iota
// AccessibleWhenUnlocked is when unlocked
AccessibleWhenUnlocked
// AccessibleAfterFirstUnlock is after first unlock
AccessibleAfterFirstUnlock
// AccessibleAlways is always
AccessibleAlways
// AccessibleWhenPasscodeSetThisDeviceOnly is when passcode is set
AccessibleWhenPasscodeSetThisDeviceOnly
// AccessibleWhenUnlockedThisDeviceOnly is when unlocked for this device only
AccessibleWhenUnlockedThisDeviceOnly
// AccessibleAfterFirstUnlockThisDeviceOnly is after first unlock for this device only
AccessibleAfterFirstUnlockThisDeviceOnly
// AccessibleAccessibleAlwaysThisDeviceOnly is always for this device only
AccessibleAccessibleAlwaysThisDeviceOnly

```




### Type AccessibleResult
```go
type AccessibleResult struct {
	QueryResult
	Accessible Accessible
}
```
AccessibleResult is a query result together with the accessibility that
kSecAttrAccessible records for the item, ie. when its contents can be
decrypted.

The attribute applies to items in the data protection keychain and to
synchronizable items. An item that does not carry it, such as one in the
macOS file based keychain, whose access is governed by an access control
list instead, reports AccessibleDefault.

### Functions

```go
func QueryItemAccessible(item Item) ([]AccessibleResult, error)
```
QueryItemAccessible is QueryItem, except that each result also reports
the item's accessibility. As with QueryItem, the query must not set
SetReturnRef, and SetReturnAttributes is needed for there to be any
attributes to report.




### Type Convertable
```go
type Convertable interface{ Convert() (C.CFTypeRef, error) }
```
Convertable knows how to convert an instance to a CFTypeRef.


### Type Error
```go
type Error int
```
Error defines keychain errors

### Methods

```go
func (k Error) Error() (msg string)
```




### Type Item
```go
type Item struct {
	// contains filtered or unexported fields
}
```
Item for adding, querying or deleting.

### Functions

```go
func NewGenericPassword(service string, account string, label string, data []byte, accessGroup string) Item
```
NewGenericPassword creates a generic password item with the default
keychain. This is a convenience method.


```go
func NewItem() Item
```
NewItem is a new empty keychain item



### Methods

```go
func (k *Item) SetAccessGroup(ag string)
```
SetAccessGroup sets the access group attribute


```go
func (k *Item) SetAccessible(accessible Accessible)
```
SetAccessible sets the accessible attribute


```go
func (k *Item) SetAccount(a string)
```
SetAccount sets the account attribute


```go
func (k *Item) SetAuthenticationType(s string)
```
SetAuthenticationType sets the authentication type attribute (for internet
password items)


```go
func (k *Item) SetComment(s string)
```
SetComment sets the comment attribute


```go
func (k *Item) SetData(b []byte)
```
SetData sets the data attribute


```go
func (k *Item) SetDataProtectionKeyChain(value bool)
```
SetDataProtectionKeyChain sets the data protection keychain attribute.


```go
func (k *Item) SetDescription(s string)
```
SetDescription sets the description attribute


```go
func (k *Item) SetInt32(key string, v int32)
```
SetInt32 sets an int32 attribute for a string key


```go
func (k *Item) SetLabel(l string)
```
SetLabel sets the label attribute


```go
func (k *Item) SetMatchLimit(matchLimit MatchLimit)
```
SetMatchLimit sets the match limit


```go
func (k *Item) SetPath(s string)
```
SetPath sets the path attribute (for internet password items)


```go
func (k *Item) SetPort(v int32)
```
SetPort sets the port attribute (for internet password items)


```go
func (k *Item) SetProtocol(s string)
```
SetProtocol sets the protocol attribute (for internet password items)
Example values are: "htps", "http", "smb "


```go
func (k *Item) SetReturnAttributes(b bool)
```
SetReturnAttributes sets the return value type on query


```go
func (k *Item) SetReturnData(b bool)
```
SetReturnData enables returning data on query


```go
func (k *Item) SetReturnRef(b bool)
```
SetReturnRef enables returning references on query


```go
func (k *Item) SetSecClass(sc SecClass)
```
SetSecClass sets the security class


```go
func (k *Item) SetServer(s string)
```
SetServer sets the server attribute (for internet password items)


```go
func (k *Item) SetService(s string)
```
SetService sets the service attribute (for generic application items)


```go
func (k *Item) SetString(key string, s string)
```
SetString sets a string attibute for a string key


```go
func (k *Item) SetSynchronizable(sync Synchronizable)
```
SetSynchronizable sets the synchronizable attribute




### Type MatchLimit
```go
type MatchLimit int
```
MatchLimit is whether to limit results on query


### Type QueryResult
```go
type QueryResult struct { // For generic application items
	Service            string // For internet password items
	Server             string
	Protocol           string
	AuthenticationType string
	Port               int32
	Path               string
	Account            string
	AccessGroup        string
	Label              string
	Description        string
	Comment            string
	Data               []byte
	CreationDate       time.Time
	ModificationDate   time.Time
}
```
QueryResult stores all possible results from queries. Not all fields are
applicable all the time. Results depend on query.

### Functions

```go
func QueryItem(item Item) ([]QueryResult, error)
```
QueryItem returns a list of query results.




### Type SecClass
```go
type SecClass int
```
SecClass is the items class code

### Variables
### SecClassGenericPassword, SecClassInternetPassword
```go
/*
	kSecClassGenericPassword item attributes:
	 kSecAttrAccess (OS X only)
	 kSecAttrAccessGroup (iOS; also OS X if kSecAttrSynchronizable specified)
	 kSecAttrAccessible (iOS; also OS X if kSecAttrSynchronizable specified)
	 kSecAttrAccount
	 kSecAttrService
*/ SecClassGenericPassword SecClass = 1
SecClassInternetPassword SecClass = 2

```
Keychain Item Classes




### Type Synchronizable
```go
type Synchronizable int
```
Synchronizable is the items synchronizable status





