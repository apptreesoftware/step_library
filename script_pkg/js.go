package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/robertkrimen/otto"
)

type ScriptInput struct {
	Script     string
	ScriptVars map[string]interface{}
}

type ScriptOutput struct {
	ReturnVal interface{}
}

type JavascriptRunner struct {
}

func (JavascriptRunner) Name() string {
	return "js"
}

func (JavascriptRunner) Version() string {
	return "1.0"
}

func (j JavascriptRunner) Execute(ctx step.Context) (interface{}, error) {
	input := ScriptInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return j.execute(input)
}

func (JavascriptRunner) execute(input ScriptInput) (ScriptOutput, error) {
	vm := otto.New()

	for key, value := range input.ScriptVars {
		err := vm.Set(key, value)
		if err != nil {
			return ScriptOutput{}, err
		}
	}

	jsScript := fmt.Sprintf(`(function() {
	%s
})()`, input.Script)

	val, err := vm.Run(jsScript)
	if err != nil {
		return ScriptOutput{}, err
	}
	var returnVal interface{}
	if val.IsDefined() {
		if val.IsBoolean() {
			returnVal, err = val.ToBoolean()
		} else if val.IsNumber() {
			returnVal, err = val.ToInteger()
			if err != nil {
				returnVal, err = val.ToFloat()
			}
		} else if val.IsString() {
			returnVal, err = val.ToString()
		} else if val.IsObject() {
			returnVal, _ = val.Export()
		}
	}
	return ScriptOutput{ReturnVal: returnVal}, err
}
