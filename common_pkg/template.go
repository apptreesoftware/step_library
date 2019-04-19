package main

import (
	"bytes"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"reflect"
	"text/template"
)

type Template struct {
}

func (Template) Name() string {
	return "template"
}

func (Template) Version() string {
	return "1.0"
}

func (t Template) Execute(in step.Context) (interface{}, error) {
	input := TemplateInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return t.execute(input)
}

func (Template) execute(input TemplateInput) (TemplateOutput, error) {
	tmpl := template.New("step").Funcs(templateFuncs())
	compiledTemplate, err := tmpl.Parse(input.Template)
	if err != nil {
		return TemplateOutput{}, err
	}
	var buf bytes.Buffer
	err = compiledTemplate.Execute(&buf, input.Record)
	if err != nil {
		return TemplateOutput{}, err
	}
	return TemplateOutput{
		Output: buf.String(),
	}, nil
}

func templateFuncs() template.FuncMap {
	tf := template.FuncMap{
		"isInt": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				return true
			default:
				return false
			}
		},
		"isString": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.String:
				return true
			default:
				return false
			}
		},
		"isSlice": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Slice:
				return true
			default:
				return false
			}
		},
		"isArray": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Array:
				return true
			default:
				return false
			}
		},
		"isMap": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Map:
				return true
			default:
				return false
			}
		},
	}
	return tf
}

type TemplateInput struct {
	Template string
	Record   map[string]interface{}
}

type TemplateOutput struct {
	Output string
}