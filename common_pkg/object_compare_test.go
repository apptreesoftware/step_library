package main

import "testing"

const testValue0 = `{"Left":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":["name","age","last","height","cousin","Teams"]}`
const testValue1 = `{"Left":{"name":"samo","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":["name","age","last","height","cousin","Teams"]}`
const testValue2 = `{"Left":{"name":"samo","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacyy","age":22},"Teams":["orozco","sco"]},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":["name","age","last","height","cousin","Teams"]}`
const testValue3 = `{"Left":{"name":"sam","age":24,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":[]}`
const testValue4 = `{"Left":{"name":"samo","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":["name","age","last","height","cousin","Teams"],"FieldsToExclude":["name"]}`
const testValue5 = `{"Left":{},"Right":{"name":"sam","age":25,"last":"orozco","height":6.1,"cousin":{"name":"kacy","age":22},"Teams":["orozco","sco"]},"FieldsToCompare":[]}`

func TestObjectDiff_ExecuteJson(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue0)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if !val.Equal {
		t.Fail()
	}
}

func TestObjectDiff_ExecuteJson1(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue1)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if val.Equal {
		t.Fail()
	}

	if len(val.FieldsThatDiffered) != 1 {
		t.Fail()
	}
}

func TestObjectDiff_ExecuteJson2(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue2)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if val.Equal {
		t.Fail()
	}

	if len(val.FieldsThatDiffered) != 2 {
		t.Fail()
	}
}

func TestObjectDiff_ExecuteJson3(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue3)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if val.Equal {
		t.Fail()
	}

	if len(val.FieldsThatDiffered) != 1 {
		t.Fail()
	}

	if val.FieldsThatDiffered[0] != "age" {
		t.Fail()
	}
}

func TestObjectDiff_ExecuteJson4(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue4)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if !val.Equal {
		t.Fail()
	}
}

func TestObjectDiff_ExecuteJson5(t *testing.T) {
	diff := ObjectCompare{}
	obj, err := diff.ExecuteJson(testValue5)
	if err != nil {
		panic(err)
	}
	val := obj.(ObjectCompareOutput)
	if val.Equal {
		t.Fail()
	}
}
