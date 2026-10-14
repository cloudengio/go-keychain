// Copyright 2026 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

//go:build darwin && !ios

// This file is an addition made in the cloudengio fork and has no upstream
// counterpart. It reports the accessibility that kSecAttrAccessible records
// for an item, which QueryItem does not: QueryResult has no field for it and
// convertResult ignores the key.
//
// It is deliberately kept to itself, rather than adding a field to
// QueryResult and a case to convertResult, so that merging upstream changes
// touches nothing of it: every declaration here is new, and the only
// upstream code it depends on is called, not edited. The build tag matches
// macos.go, which is where the accessibilities this maps back from are
// defined.
package keychain

/*
#cgo LDFLAGS: -framework CoreFoundation -framework Security
#include <CoreFoundation/CoreFoundation.h>
#include <Security/Security.h>
*/
import "C"

import "fmt"

// AccessibleResult is a query result together with the accessibility that
// kSecAttrAccessible records for the item, ie. when its contents can be
// decrypted.
//
// The attribute applies to items in the data protection keychain and to
// synchronizable items. An item that does not carry it, such as one in the
// macOS file based keychain, whose access is governed by an access control
// list instead, reports AccessibleDefault.
type AccessibleResult struct {
	QueryResult
	Accessible Accessible
}

// QueryItemAccessible is QueryItem, except that each result also reports the
// item's accessibility. As with QueryItem, the query must not set
// SetReturnRef, and SetReturnAttributes is needed for there to be any
// attributes to report.
func QueryItemAccessible(item Item) ([]AccessibleResult, error) {
	resultsRef, err := QueryItemRef(item)
	if err != nil {
		return nil, err
	}
	if resultsRef == 0 {
		return nil, nil
	}
	defer Release(resultsRef)

	switch C.CFGetTypeID(resultsRef) {
	case C.CFArrayGetTypeID():
		refs := CFArrayToArray(C.CFArrayRef(resultsRef))
		results := make([]AccessibleResult, 0, len(refs))
		for _, ref := range refs {
			if C.CFGetTypeID(ref) != C.CFDictionaryGetTypeID() {
				return nil, fmt.Errorf("invalid result type: %s", CFTypeDescription(ref))
			}
			result, err := convertAccessibleResult(C.CFDictionaryRef(ref))
			if err != nil {
				return nil, err
			}
			results = append(results, *result)
		}
		return results, nil
	case C.CFDictionaryGetTypeID():
		result, err := convertAccessibleResult(C.CFDictionaryRef(resultsRef))
		if err != nil {
			return nil, err
		}
		return []AccessibleResult{*result}, nil
	}
	return nil, fmt.Errorf("invalid result type: %s", CFTypeDescription(resultsRef))
}

// convertAccessibleResult converts an item's attributes exactly as QueryItem
// does, by calling convertResult, and then adds the accessibility that
// conversion leaves out.
func convertAccessibleResult(d C.CFDictionaryRef) (*AccessibleResult, error) {
	result, err := convertResult(d)
	if err != nil {
		return nil, err
	}
	return &AccessibleResult{QueryResult: *result, Accessible: dictAccessible(d)}, nil
}

// dictAccessible returns the accessibility recorded in an item's attributes,
// or AccessibleDefault if they carry none.
func dictAccessible(d C.CFDictionaryRef) Accessible {
	for k, v := range CFDictionaryToMap(d) {
		if CFStringToString(C.CFStringRef(k)) != AccessibleKey {
			continue
		}
		return accessibleFromString(CFStringToString(C.CFStringRef(v)))
	}
	return AccessibleDefault
}

// accessibleByValue is the reverse of accessibleTypeRef: it maps the value
// that kSecAttrAccessible is recorded as, eg. "ak" for
// kSecAttrAccessibleWhenUnlocked, to the Accessible that names it. It is
// derived from accessibleTypeRef so that the two cannot disagree, and so
// that an accessibility added there is understood here without change.
var accessibleByValue = func() map[string]Accessible {
	byValue := make(map[string]Accessible, len(accessibleTypeRef))
	for accessible, ref := range accessibleTypeRef {
		byValue[CFStringToString(C.CFStringRef(ref))] = accessible
	}
	return byValue
}()

// accessibleFromString returns the Accessible that value names, or
// AccessibleDefault if it names none of them.
func accessibleFromString(value string) Accessible {
	return accessibleByValue[value]
}
