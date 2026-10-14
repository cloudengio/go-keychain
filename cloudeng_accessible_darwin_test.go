// Copyright 2026 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

//go:build darwin && !ios

// This file is an addition made in the cloudengio fork and has no upstream
// counterpart; it is kept separate from macos_test.go so that merging
// upstream changes does not touch it.
package keychain

import "testing"

// TestAccessibleFromString verifies that the values kSecAttrAccessible is
// recorded as are mapped back to the Accessible that names them, ie. that a
// query reports the accessibility an item was stored with. The mapping is
// derived from accessibleTypeRef, which is what SetAccessible stores, so the
// test checks that it covers every accessibility that can be set, that no
// two of them share a representation, and that a value naming none of them
// is reported as the default rather than as one of them.
func TestAccessibleFromString(t *testing.T) {
	if got, want := len(accessibleByValue), len(accessibleTypeRef); got != want {
		t.Fatalf("got %v values, want %v: two Accessible values share a representation", got, want)
	}
	seen := make(map[Accessible]string, len(accessibleByValue))
	for value, accessible := range accessibleByValue {
		if len(value) == 0 {
			t.Errorf("%v: empty value", accessible)
			continue
		}
		if dup, ok := seen[accessible]; ok {
			t.Errorf("%v: recorded as both %q and %q", accessible, dup, value)
			continue
		}
		seen[accessible] = value
		if got, want := accessibleFromString(value), accessible; got != want {
			t.Errorf("%q: got %v, want %v", value, got, want)
		}
	}
	for accessible := range accessibleTypeRef {
		if _, ok := seen[accessible]; !ok {
			t.Errorf("%v: cannot be mapped back from a query result", accessible)
		}
	}
	for _, value := range []string{"", "not-an-accessibility"} {
		if got, want := accessibleFromString(value), AccessibleDefault; got != want {
			t.Errorf("%q: got %v, want %v", value, got, want)
		}
	}
}
