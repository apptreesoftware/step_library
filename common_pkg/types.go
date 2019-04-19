package main

// common types
type JsonMap map[string]interface{}
type LookupMap map[string]bool

func (mp LookupMap) contains(val string) bool {
	if _, ok := mp[val]; ok {
		return ok
	}
	return false
}
