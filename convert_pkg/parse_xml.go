package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	xj "github.com/basgys/goxml2json"
	"github.com/json-iterator/go"
	"strings"
)

type ParseXml struct{}

func (ParseXml) Name() string {
	return "parse_xml_to_object"
}

func (ParseXml) Version() string {
	return "1.0"
}

func (w ParseXml) Execute(in step.Context) (interface{}, error) {
	input := ParseXmlInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	// convert XML to some bytes
	xml := strings.NewReader(input.XML)
	bytes, err := xj.Convert(xml)
	if err != nil {
		return nil, err
	}
	// stuff the bytes into an interface
	var obj map[string]interface{}
	err = jsoniter.Unmarshal(bytes.Bytes(), &obj)
	if err != nil {
		return nil, err
	}

	return ParseXmlOutput{
		Record: obj,
	}, nil
}

type ParseXmlInput struct {
	XML string
}

type ParseXmlOutput struct {
	Record interface{}
}
