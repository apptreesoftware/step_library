package main

import (
	"github.com/pkg/errors"
	"testing"
)

func TestScript(t *testing.T) {
	input := ScriptInput{
		ScriptVars: map[string]interface{}{"multiplier": 32},
		Script:     "function callMethod(data) {\n    return data * 1.5;\n  \n}\nvar output = callMethod(multiplier);\n  return output;\n",
	}

	runner := JavascriptRunner{}
	output, err := runner.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	val, ok := output.ReturnVal.(int64)
	if !ok || val != 48 {
		t.Error(errors.New("output for multiplier script was incorrect"))
		t.FailNow()
	}

	name := "myName"
	input.ScriptVars["myMap"] = map[string]interface{}{"name": name}
	input.Script = "return myMap.name;\n"

	output, err = runner.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	stringVal, ok := output.ReturnVal.(string)
	if !ok || stringVal != name {
		t.Error(errors.New("output for reading map val was incorrect"))
		t.FailNow()
	}
}
