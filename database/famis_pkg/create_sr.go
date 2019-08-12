package main

import (
	"database/sql"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library/database/db_common"
	"golang.org/x/xerrors"
	_ "gopkg.in/goracle.v2"
	"strings"
)

type CreateSRInput struct {
	SiteId           string
	Description      string
	Requester        string
	ConnectionString string
	AttachmentUrl    string
}

func (input CreateSRInput) Validate() []string {
	errors := make([]string, 0)
	if input.SiteId == "" {
		errors = append(errors, "site ID is required.")
	}
	if input.Description == "" {
		errors = append(errors, "description is required.")
	}
	if input.Requester == "" {
		errors = append(errors, "requester is required.")
	}
	if input.ConnectionString == "" {
		errors = append(errors, "connection string is required.")
	}
	return errors
}

type CreateSROutput struct {
	ServiceResponse map[string]interface{}
}

type CreateRequest struct {
}

func (CreateRequest) Name() string {
	return "create_request"
}

func (CreateRequest) Version() string {
	return "1.0"
}

func (c CreateRequest) Execute(in step.Context) (interface{}, error) {
	var input CreateSRInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, xerrors.Errorf("Unable to read inputs: %w", err)
	}
	if errors := input.Validate(); len(errors) > 0 {
		return nil, xerrors.Errorf("Invalid inputs: %s", strings.Join(errors, ", "))
	}

	return c.execute(input)
}

func (CreateRequest) execute(input CreateSRInput) (*CreateSROutput, error) {
	println("Creating sql string")
	sqlString := fmt.Sprintf("select atio_create_sr('REQUESTESD', 'S', '%s', 'CORRECTIVE', 3, '%s', 'APPTREEIO', '%s', 'N', 'ASSISTANT', '%s') as POTHOLE_REQUEST from dual",
		input.SiteId, input.Description, input.Requester, input.AttachmentUrl)

	println("Connecting to oracle")
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, xerrors.Errorf("Unable to connect to database: %w", err)
	}

	println("Running command")
	command := db_common.DatabaseCommand{
		ConnectionString: input.ConnectionString,
		Sql:              sqlString,
	}
	queryResult, err := db_common.PerformQuery(db, command)
	if err != nil {
		return &CreateSROutput{}, xerrors.Errorf("Error creating service request: %w", err)
	}

	println("parsing output")
	output, ok := queryResult.(*db_common.RowOutput)
	if !ok {
		return &CreateSROutput{}, xerrors.Errorf("Response was not correctly parsed")
	}
	if len(output.Results) == 0 {
		return &CreateSROutput{}, xerrors.Errorf("Response contained no data")
	}

	println("returning response")
	return &CreateSROutput{ServiceResponse: output.Results[0]}, nil
}
