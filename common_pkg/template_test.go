package main

import (
	"github.com/andreyvit/diff"
	"github.com/json-iterator/go"
	"strings"
	"testing"
)

type Project struct {
	Username           string
	Password           string
	AutoComplete       bool
	Comment            string
	Locked             bool
	ProjectHierarchyID string
	ProjectName        string
	StateId            string
	StatusId           string
	StartDate          string
	Capital            bool
	CompanyReferenceId string
	WorkTags           []ProjectTag
	FundId             string
}

type ProjectTag struct {
	TagName  string
	TagValue string
}

func TestTemplate_Execute(t *testing.T) {

	record := Project{
		AutoComplete:       false,
		Comment:            "GATECH AppTree Project Integration Test",
		Locked:             false,
		ProjectHierarchyID: "Admin_&amp;_Finance_Exec_VP",
		ProjectName:        "Project AppTree Testing 12",
		StateId:            "PROJECT_STATE_PROJECT",
		StatusId:           "Active",
		StartDate:          "2018-11-02",
		Capital:            false,
		CompanyReferenceId: "CO503",
		FundId:             "FD10000",
		WorkTags: []ProjectTag{
			{
				TagName:  "Custom_Organization_Reference_ID",
				TagValue: "DE00000513",
			},
			{
				TagName:  "Cost_Center_Reference_ID",
				TagValue: "CC000024",
			},
			{
				TagName:  "Custom_Organization_Reference_ID",
				TagValue: "FN12120",
			},
			{
				TagName:  "Fund_ID",
				TagValue: "FD10000",
			},
			{
				TagName:  "Custom_Organization_Reference_ID",
				TagValue: "CL11200",
			},
		},
	}
	b, _ := jsoniter.Marshal(&record)
	jsonMap := map[string]interface{}{}
	err := jsoniter.Unmarshal(b, &jsonMap)
	if err != nil {
		t.Error(err)
		return
	}

	input := TemplateInput{
		Template: xmlTemplate,
		Record:   jsonMap,
	}
	out, err := Template{}.execute(input)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if a, e := strings.TrimSpace(xmlComplete), strings.TrimSpace(out.Output); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
	}
}

var xmlTemplate = `<?xml version="1.0" encoding="utf-8"?>
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
    <env:Header>
    </env:Header>
    <env:Body>
        <wd:Submit_Project_Request wd:Add_Only="true" wd:version="v30.2" 
            xmlns:wd="urn:com.workday/bsvc">
            <wd:Business_Process_Parameters>
                <wd:Auto_Complete>{{.AutoComplete}}</wd:Auto_Complete>
                <wd:Comment_Data>
                    <wd:Comment>{{.Comment}}</wd:Comment>
                </wd:Comment_Data>
            </wd:Business_Process_Parameters>
            <wd:Project_Data>
                <wd:Locked_in_Workday>{{.Locked}}</wd:Locked_in_Workday>
                <wd:Project_Hierarchy_Reference>
                    <wd:ID wd:type="Project_Hierarchy_ID">{{.ProjectHierarchyID}}</wd:ID>
                </wd:Project_Hierarchy_Reference>
                <wd:Project_Name>{{.ProjectName}}</wd:Project_Name>
                <wd:Project_State_Reference>
                    <wd:ID wd:type="Project_State_ID">{{.StateId}}</wd:ID>
                </wd:Project_State_Reference>
                <wd:Project_Status_Reference>
                    <wd:ID wd:type="Project_Status_ID">{{.StatusId}}</wd:ID>
                </wd:Project_Status_Reference>
                <wd:Start_Date>{{.StartDate}}</wd:Start_Date>
                <wd:Capital>{{.Capital}}</wd:Capital>
                <wd:Company_Reference>
                    <wd:ID wd:type="Company_Reference_ID">{{.CompanyReferenceId}}</wd:ID>
                </wd:Company_Reference>
                <wd:Worktags_Data>
					{{range .WorkTags}}
					<wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="{{.TagName}}">{{.TagValue}}</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
					{{end}}
                </wd:Worktags_Data>
                <wd:Balancing_Worktag_Reference>
                    <wd:ID wd:type="Fund_ID">{{.FundId}}</wd:ID>
                </wd:Balancing_Worktag_Reference>
            </wd:Project_Data>
        </wd:Submit_Project_Request>
    </env:Body>
</env:Envelope>
`

var xmlComplete = `<?xml version="1.0" encoding="utf-8"?>
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
    <env:Header>
    </env:Header>
    <env:Body>
        <wd:Submit_Project_Request wd:Add_Only="true" wd:version="v30.2" 
            xmlns:wd="urn:com.workday/bsvc">
            <wd:Business_Process_Parameters>
                <wd:Auto_Complete>false</wd:Auto_Complete>
                <wd:Comment_Data>
                    <wd:Comment>GATECH AppTree Project Integration Test</wd:Comment>
                </wd:Comment_Data>
            </wd:Business_Process_Parameters>
            <wd:Project_Data>
                <wd:Locked_in_Workday>false</wd:Locked_in_Workday>
                <wd:Project_Hierarchy_Reference>
                    <wd:ID wd:type="Project_Hierarchy_ID">Admin_&amp;_Finance_Exec_VP</wd:ID>
                </wd:Project_Hierarchy_Reference>
                <wd:Project_Name>Project AppTree Testing 12</wd:Project_Name>
                <wd:Project_State_Reference>
                    <wd:ID wd:type="Project_State_ID">PROJECT_STATE_PROJECT</wd:ID>
                </wd:Project_State_Reference>
                <wd:Project_Status_Reference>
                    <wd:ID wd:type="Project_Status_ID">Active</wd:ID>
                </wd:Project_Status_Reference>
                <wd:Start_Date>2018-11-02</wd:Start_Date>
                <wd:Capital>false</wd:Capital>
                <wd:Company_Reference>
                    <wd:ID wd:type="Company_Reference_ID">CO503</wd:ID>
                </wd:Company_Reference>
                <wd:Worktags_Data>
                    <wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="Custom_Organization_Reference_ID">DE00000513</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
                    <wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="Cost_Center_Reference_ID">CC000024</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
                    <wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="Custom_Organization_Reference_ID">FN12120</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
                    <wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="Fund_ID">FD10000</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
                    <wd:Related_Worktags_by_Type_Data>
                        <wd:Default_Worktag_Data>
                            <wd:Default_Worktag_Reference>
                                <wd:ID wd:type="Custom_Organization_Reference_ID">CL11200</wd:ID>
                            </wd:Default_Worktag_Reference>
                        </wd:Default_Worktag_Data>
                    </wd:Related_Worktags_by_Type_Data>
                </wd:Worktags_Data>
                <wd:Balancing_Worktag_Reference>
                    <wd:ID wd:type="Fund_ID">FD10000</wd:ID>
                </wd:Balancing_Worktag_Reference>
            </wd:Project_Data>
        </wd:Submit_Project_Request>
    </env:Body>
</env:Envelope>
`

func TestSlackMessage(t *testing.T) {

	expectedOutput := `{
			"type" : "mrkdwn",
			"text" : "*VM (Vicky) Brasseur—   Open Source: What even is? How even to?*\n Portland JR DEVELOPER Meetup!\n https://www.meetup.com/Portland-JR-DEVELOPER-Meetup/events/256869518/\n 2019-04-17\n 17:30\n *Vacasa*\n 926 NW 13th Ave\n Portland, OR "
		}`

	template := `{
			"type" : "mrkdwn",
			"text" : "*{{.Name}}*\n {{.Group.Name}}\n {{.Link}}\n {{.Date}}\n {{.Time}}\n *{{.Venue.Name}}*\n {{.Venue.Address}}\n {{.Venue.City}}, {{.Venue.State}} {{.Venue.Zip}}"
		}`

	event := Json{
		"Name": "VM (Vicky) Brasseur—   Open Source: What even is? How even to?",
		"Group": Json{
			"Name": "Portland JR DEVELOPER Meetup!",
		},
		"Link": "https://www.meetup.com/Portland-JR-DEVELOPER-Meetup/events/256869518/",
		"Date": "2019-04-17",
		"Time": "17:30",
		"Venue": Json{
			"Name":    "Vacasa",
			"Address": "926 NW 13th Ave",
			"City":    "Portland",
			"State":   "OR",
			"Zip":     "",
		},
	}

	output, err := Template{}.execute(TemplateInput{
		Template: template,
		Record:   event,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if output.Output != expectedOutput {
		t.Logf("Outputs do not match: expected: \n%s; got: \n\n%s",
			expectedOutput, output.Output)
		t.Fail()
	}

}

type Json map[string]interface{}
