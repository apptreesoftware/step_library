package main

import (
	"database/sql"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
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
	ServiceRequestId string
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
	sqlString := fmt.Sprintf("select atio_create_sr('REQUESTESD', 'S', '%s', 'CORRECTIVE', 3, '%s', 'APPTREEIO', '%s', 'N', 'ASSISTANT', '%s') as POTHOLE_REQUEST from dual",
		input.SiteId, input.Description, input.Requester, input.AttachmentUrl)

	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, xerrors.Errorf("Unable to connect to database: %w", err)
	}

	rows, err := db.Query(sqlString)
	if err != nil {
		return nil, xerrors.Errorf("Unable to run statement: %w", err)
	}

	defer rows.Close()
	var srId string
	for rows.Next() {
		err = rows.Scan(&srId)
		if err != nil {
			return &CreateSROutput{}, xerrors.Errorf("Unable to read results: %w", err)
		}
	}
	if err := rows.Err(); err != nil {
		return &CreateSROutput{}, xerrors.Errorf("Unable to read results: %w", err)
	}
	return &CreateSROutput{ServiceRequestId: srId}, nil
}
