package main

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library/database/db_common"
	"golang.org/x/xerrors"
	_ "gopkg.in/goracle.v2"
	"strings"
)

type CreateSRInput struct {
	ConnectionString string
	SiteId           string
	Description      string
	Requester        string
	AttachmentUrl    string
	Notes            string
	EquipmentId      string
	RequestStatus    string
	RequestType      string
	MaintenanceType  string
	Priority         string
	EnterUser        string
	NeedNotification bool
	RequestSource    string
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
	sqlString := "select atio_create_sr(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14) as APPTREE_ASSISTANT_SR_REQUEST from dual"
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, xerrors.Errorf("Unable to connect to database: %w", err)
	}
	defer db.Close()

	command := db_common.DatabaseCommand{
		ConnectionString: input.ConnectionString,
		Sql:              sqlString,
	}

	queryResult, err := db_common.PerformQueryWithArgs(db, command, createArgsFromInput(input)...)
	if err != nil {
		return nil, xerrors.Errorf("Error creating service request: %w", err)
	}

	output, ok := queryResult.(*db_common.RowOutput)
	if !ok {
		return nil, xerrors.Errorf("Response was not correctly parsed")
	}
	if len(output.Results) == 0 {
		return nil, xerrors.Errorf("Response contained no data")
	}

	requestId, ok := output.Results[0]["APPTREE_ASSISTANT_SR_REQUEST"].(string)
	if !ok {
		return nil, xerrors.Errorf("Response ID was not a string")
	}

	return &CreateSROutput{ServiceRequestId: requestId}, nil
}

func createArgsFromInput(input CreateSRInput) []interface{} {
	args := make([]interface{}, 14)
	args[0] = input.RequestStatus
	args[1] = input.RequestType
	args[2] = input.SiteId
	args[3] = input.MaintenanceType
	args[4] = input.Priority
	args[5] = input.Description
	args[6] = input.EnterUser
	args[7] = input.Requester
	args[8] = "N"
	if input.NeedNotification {
		args[8] = "Y"
	}
	args[9] = input.RequestSource
	args[10] = nil
	if input.EquipmentId != "" {
		args[10] = input.EquipmentId
	}
	args[11] = nil
	args[12] = nil
	if input.Notes != "" {
		args[12] = input.Notes
	}
	args[13] = nil
	if input.AttachmentUrl != "" {
		args[13] = input.AttachmentUrl
	}
	return args
}
