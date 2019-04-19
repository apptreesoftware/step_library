package main

import (
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"reflect"
)

type ObjectCompare struct {
}

type ObjectCompareInput struct {
	Left            JsonMap
	Right           JsonMap
	FieldsToCompare []string
	FieldsToExclude []string
}

type ObjectCompareOutput struct {
	Equal              bool
	FieldsThatDiffered []string
}

func (ObjectCompare) Name() string {
	return "object_compare"
}

func (ObjectCompare) Version() string {
	return "1.0"
}

// This step takes two json objects `Left` and `Right` and a list of
// fields or `FieldsToCompare` you wish this step to compare
// and finally a list of fields you would NOT like to compare

// if `FieldsToCompare` is not provided this step will compare every single
// field in the left object and compare it to the right
//
func (diff ObjectCompare) Execute(in step.Context) (interface{}, error) {
	objectDiffIn := &ObjectCompareInput{}
	err := in.BindInputs(objectDiffIn)
	if err != nil {
		return nil, err
	}
	return diff.execute(objectDiffIn)
}

func (diff ObjectCompare) ExecuteJson(jsonString string) (interface{}, error) {
	objectDiffIn := &ObjectCompareInput{}
	err := json.Unmarshal([]byte(jsonString), objectDiffIn)
	if err != nil {
		return nil, err
	}
	return diff.execute(objectDiffIn)
}

func (diff ObjectCompare) execute(jsonObj *ObjectCompareInput) (interface{}, error) {
	left := jsonObj.Left
	right := jsonObj.Right
	exclude := diff.getLookupMap(jsonObj.FieldsToExclude)

	// did the user specify fields to compare?
	if fields := jsonObj.FieldsToCompare; fields != nil && len(fields) > 0 {
		return diff.diffFields(fields, left, right, exclude), nil
	}

	// we will always get the fields to check from the Left map
	mapFields := diff.getStringKeysFromMap(left)
	if len(mapFields) == 0 {
		mapFields = diff.getStringKeysFromMap(right)
	}
	return diff.diffFields(mapFields, left, right, exclude), nil
}

// This method takes a slice of fields names you would like to compare and two json objects to compare on
// It will then iterator each field and check to see if they differ in the given objects
// If it finds differences it will record the field that was different and indicate that the objects differed
// if also accepts an options LookupMap that will exclude any fields contained in the map
func (diff ObjectCompare) diffFields(fields []string, left JsonMap, right JsonMap, looks LookupMap) ObjectCompareOutput {
	// are the two objs different
	isDifferent := false
	// a collection of the fields that are different
	differentFields := make([]string, 0)
	// iterate  given fields and check if the there are differences
	for _, field := range fields {
		// if the user has selected to exclude this field
		if looks != nil && looks.contains(field) {
			continue
		}
		// the fields values differ
		if diff.fieldsDiffer(field, left, right) {
			isDifferent = true
			differentFields = append(differentFields, field)
		}
	}
	return ObjectCompareOutput{Equal: !isDifferent, FieldsThatDiffered: differentFields}

}

func (diff ObjectCompare) fieldsDiffer(field string, left JsonMap, right JsonMap) bool {
	leftData := left[field]
	rightData := right[field]

	if leftData == nil && rightData != nil {
		return true
	}
	if rightData == nil && leftData != nil {
		return true
	}
	if leftData == nil && rightData == nil {
		return false
	}

	// are these types comparable?
	// if not deepEqual
	// in most cases the non comparable data types are slices and maps
	if !reflect.TypeOf(leftData).Comparable() || !reflect.TypeOf(rightData).Comparable() {
		return !reflect.DeepEqual(leftData, rightData)
	}

	return leftData != rightData
}

func (diff ObjectCompare) getLookupMap(fields []string) LookupMap {
	result := make(map[string]bool, 0)
	if fields == nil || len(fields) < 1 {
		return result
	}

	for _, val := range fields {
		result[val] = true
	}
	return result
}

// I know this is causing me to make two passes over the map keys
// I am ok with that because of the simplicity on `diffing` the fields
func (diff ObjectCompare) getStringKeysFromMap(data JsonMap) []string {
	if data == nil {
		return make([]string, 0)
	}
	keys := reflect.ValueOf(data).MapKeys()
	if len(keys) < 1 {
		return make([]string, 0)
	}
	result := make([]string, len(keys))
	// iterate keys and put strings into result
	for idx, key := range keys {
		result[idx] = key.String()
	}
	return result
}
