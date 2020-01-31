package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

type CreateAttachmentInput struct {
	Facility360Input
	RequestId      int64
	AttachmentLink string
}

type AttachmentRequest struct {
	RequestId string
	Contents  string
	Name      string
}

type CreateAttachment struct {
	Fetch
}

func (CreateAttachment) Name() string {
	return "create_wo_attachment"
}

func (CreateAttachment) Description() string {
	return "Creates an attachment for a work order"
}

func (c CreateAttachment) Execute(ctx step.Context) (interface{}, error) {
	input := CreateAttachmentInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return c.execute(input)
}

func (create CreateAttachment) execute(input CreateAttachmentInput) (interface{}, error) {
	if input.AttachmentLink == "" {
		return nil, errors.New("attachment link is required")
	}
	fileBytes, err := getFileBytes(input.AttachmentLink)
	if err != nil {
		return nil, err
	}
	sEnc := base64.StdEncoding.EncodeToString(fileBytes)
	fileName, err := getFileName(input.AttachmentLink)
	if err != nil {
		return nil, err
	}

	attachment := AttachmentRequest{
		RequestId: strconv.Itoa(int(input.RequestId)),
		Contents:  sEnc,
		Name:      fileName,
	}

	err = create.LogMeInFacility360(input.Facility360Input)
	if err != nil {
		return nil, err
	}

	createUrl, err := create.getUrl(input.Url, input.Endpoint)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(attachment)
	if err != nil {
		return nil, err
	}

	req, err := create.buildRequest(createRequestMethod, createUrl.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := create.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		// if status code is not 200
		// read body as string and return
		return create.handleFailedResponse(resp)
	}
	defer resp.Body.Close()
	return create.handleUpsertResponse(resp)
}

func getFileBytes(attachmentLink string) ([]byte, error) {
	resp, err := http.Get(attachmentLink)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to download file from link: %s", err.Error()))
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading response bytes: %s", err.Error()))
	}

	return responseData, nil
}

func getFileName(attachmentLink string) (string, error) {
	u, err := url.Parse(attachmentLink)
	if err != nil {
		return "", err
	}

	_, fileName := path.Split(u.Path)
	if len(fileName) > 30 {
		splitName := strings.Split(fileName, ".")
		if len(splitName) < 2 {
			return "", errors.New("filename must include extension")
		}
		extension := splitName[len(splitName)-1]
		name := splitName[0]
		nameLength := 29 - len(extension)
		name = name[0:nameLength]
		fileName = fmt.Sprintf("%s.%s", name, extension)
	}
	return fileName, nil
}
