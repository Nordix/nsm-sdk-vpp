// Code generated by "-output l2_bridge_domain_map.gen.go -type l2BridgeDomain<bridgeDomainKey,*bridgeDomain> -output l2_bridge_domain_map.gen.go -type l2BridgeDomain<bridgeDomainKey,*bridgeDomain>"; DO NOT EDIT.
package l2bridgedomain

import (
	"sync" // Used by sync.Map.
)

// Generate code that will fail if the constants change value.
func _() {
	// An "cannot convert l2BridgeDomain literal (type l2BridgeDomain) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)(l2BridgeDomain{})
}

var _nil_l2BridgeDomain_bridgeDomain_value = func() (val *bridgeDomain) { return }()

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *l2BridgeDomain) Load(key bridgeDomainKey) (*bridgeDomain, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if value == nil {
		return _nil_l2BridgeDomain_bridgeDomain_value, ok
	}
	return value.(*bridgeDomain), ok
}

// Store sets the value for a key.
func (m *l2BridgeDomain) Store(key bridgeDomainKey, value *bridgeDomain) {
	(*sync.Map)(m).Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *l2BridgeDomain) LoadOrStore(key bridgeDomainKey, value *bridgeDomain) (*bridgeDomain, bool) {
	actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
		return _nil_l2BridgeDomain_bridgeDomain_value, loaded
	}
	return actual.(*bridgeDomain), loaded
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *l2BridgeDomain) LoadAndDelete(key bridgeDomainKey) (value *bridgeDomain, loaded bool) {
	actual, loaded := (*sync.Map)(m).LoadAndDelete(key)
	if actual == nil {
		return _nil_l2BridgeDomain_bridgeDomain_value, loaded
	}
	return actual.(*bridgeDomain), loaded
}

// Delete deletes the value for a key.
func (m *l2BridgeDomain) Delete(key bridgeDomainKey) {
	(*sync.Map)(m).Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *l2BridgeDomain) Range(f func(key bridgeDomainKey, value *bridgeDomain) bool) {
	(*sync.Map)(m).Range(func(key, value interface{}) bool {
		return f(key.(bridgeDomainKey), value.(*bridgeDomain))
	})
}