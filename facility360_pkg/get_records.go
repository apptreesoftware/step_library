package main

import (
	"bytes"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	json "github.com/json-iterator/go"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const loginEndpoint = "/MobileWebServices/api/Login"
const jsonContentType = "application/json"
const usernameKey = "Username"
const passwordKey = "Password"
const skipKey = "$skip"
const authTokenHeaderKey = "Authorization"
const contentTypeHeaderKey = "Content-Type"
const bearer = "Bearer"
const valueKey = "value"
const defaultTimeout = time.Minute * 1
const defaultWoBatchSize = 100

type Fetch struct {
	authItem AuthItem
	client   *http.Client
	username string
	password string
	url      string
}

func (Fetch) Name() string {
	return "get_records"
}

func (Fetch) Version() string {
	return "1.0"
}

func (fetch Fetch) Execute(in step.Context) (interface{}, error) {
	input := FetchInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.execute(input)
}

func (fetch Fetch) execute(input FetchInput) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authItem, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	uri, err := fetch.buildUrl(input)
	if err != nil {
		return nil, err
	}

	var records []jsoniter.RawMessage
	err = fetch.performPagedFetch(uri, input.Endpoint, func(messages []jsoniter.RawMessage) {
		records = append(records, messages...)
	})
	if err != nil {
		return nil, err
	}
	return FetchListOutputs{
		Count:   len(records),
		Records: records,
	}, nil
}

type PagedResultHandler func([]jsoniter.RawMessage)

func (fetch Fetch) performPagedFetch(uri *url.URL, endpoint string, handler PagedResultHandler) error {

	if err := fetch.RefreshIfNeeded(); err != nil {
		return err
	}
	req, err := fetch.buildRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return err
	}

	resp, err := fetch.getHttpClient().Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 404 {
		return fmt.Errorf("The service responded with a status code of %d. Verify that your Url and Endpoint inputs are correct.", resp.StatusCode)
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("The service responded with a status code of %d.", resp.StatusCode)
	}
	defer resp.Body.Close()

	var parsedResp PagedResponse
	if strings.HasSuffix(endpoint, "workorders") {
		parsedResp = NewWorkOrderPagedResponse(defaultWoBatchSize, *uri)
	} else {
		parsedResp = &NextLinkPagedResponse{}
	}
	dec := jsoniter.NewDecoder(resp.Body)
	err = dec.Decode(&parsedResp)
	if err != nil {
		return err
	}
	handler(parsedResp.GetCurrent())

	if parsedResp.HasMoreResults() {
		nextUri := parsedResp.NextPageUrl()
		return fetch.performPagedFetch(&nextUri, endpoint, handler)
	}
	return nil
}

func (fetch Fetch) buildRequest(method, url string, body io.Reader) (*http.Request, error) {
	headers := buildHeaders(fetch.authItem.AccessToken, jsonContentType)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	return req, nil
}

func (fetch Fetch) buildUrl(input FetchInput) (*url.URL, error) {
	uri, err := fetch.getUrl(input.Url, input.Endpoint)
	if err != nil {
		return nil, err
	}
	q := uri.Query()
	if input.Expand != "" {
		q.Add("$expand", input.Expand)
	}
	if input.Select != "" {
		q.Add("$select", input.Select)
	}
	if input.Filter != "" {
		q.Add("$filter", input.Filter)
	}
	if input.Skip > 0 {
		q.Add("$skip", fmt.Sprintf("%d", input.Skip))
	}
	if input.Top > 0 {
		q.Add("$top", fmt.Sprintf("%d", input.Top))
	}
	uri.RawQuery = q.Encode()
	uri = odataUrl(*uri)
	return uri, nil
}

func (fetch *Fetch) LogMeIn() error {
	auth, err := fetch.Login(fetch.username, fetch.password, fetch.url)
	if err != nil {
		return err
	}
	fetch.authItem = auth
	return nil
}

func (fetch Fetch) Login(username, password, baseUrl string) (AuthItem, error) {
	// validate and get login url
	loginUrl, err := fetch.getUrl(baseUrl, loginEndpoint)
	if err != nil {
		return AuthItem{}, err
	}
	// get login body from username and password
	loginBody, err := getReaderFromLoginBody(username, password)
	if err != nil {
		return AuthItem{}, err
	}
	client := fetch.getHttpClient()
	resp, err := client.Post(loginUrl.String(), jsonContentType, loginBody)
	if err != nil {
		return AuthItem{}, err
	}
	defer resp.Body.Close()

	loginResp, err := handleLoginResponse(resp.Body)
	if err != nil {
		return AuthItem{}, nil
	}
	if loginResp.Result == false {
		return AuthItem{}, fmt.Errorf("Login Failed: %s", loginResp.Message)
	}
	return loginResp.Auth, nil
}

func (fetch Fetch) RefreshIfNeeded() error {
	now := time.Now()
	// time subtract a minute
	now = now.Add(time.Duration(-1) * 1)
	// we're NOT expired
	expires, err := fetch.authItem.GetExpiration()
	if err != nil {
		return err
	}
	if !now.After(expires) {
		// do nothing
		return nil
	} else {
		// login again and set values on fetcher
		err := fetch.LogMeIn()
		if err != nil {
			return err
		}
		return nil
	}
}

func buildHeaders(auth, content string) http.Header {
	headers := http.Header{}
	// we need to append `Bearer` to the auth token
	headers.Add(authTokenHeaderKey, fmt.Sprintf("%s %s", bearer, auth))
	headers.Add(contentTypeHeaderKey, content)
	return headers
}

func handleFetchResponse(resp *http.Response) (JsonMap, error) {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := make(JsonMap, 0)
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func handleSingleResponse(data JsonMap) (FetchSingleOutput, error) {
	values := data[valueKey]
	res, ok := values.([]interface{})
	if !ok {
		return FetchSingleOutput{}, nil
	}
	if len(res) < 1 {
		return FetchSingleOutput{}, nil
	}

	curData := res[0].(map[string]interface{})
	return FetchSingleOutput{Record: curData, Found: true}, nil
}

func handleLoginResponse(reader io.Reader) (LoginResponse, error) {
	dec := jsoniter.NewDecoder(reader)
	loginResp := LoginResponse{}
	err := dec.Decode(&loginResp)

	return loginResp, err
}

func getReaderFromLoginBody(username, password string) (io.Reader, error) {
	loginContents, err := getLoginBody(username, password)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(loginContents), nil
}

func getLoginBody(username, password string) ([]byte, error) {
	body := make(JsonMap, 0)
	body[usernameKey] = username
	body[passwordKey] = password
	return getContents(body)
}

func getContents(data JsonMap) ([]byte, error) {
	return json.Marshal(data)
}

func (fetch Fetch) getUrl(base, endpoint string) (*url.URL, error) {
	urlVal, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	urlVal.Path = endpoint
	return urlVal, nil
}

func (fetch *Fetch) getHttpClient() *http.Client {
	if fetch.client == nil {
		fetch.client = &http.Client{
			Timeout: defaultTimeout,
		}
	}
	return fetch.client
}

func (fetch Fetch) handleFailedResponse(resp *http.Response) (Facility360UpsertOut, error) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Facility360UpsertOut{}, err
	}

	mapResp := map[string]interface{}{}
	err = json.Unmarshal(data, &mapResp)
	if err != nil {
		return Facility360UpsertOut{
			Success: false,
			Message: string(data),
			Record:  nil,
		}, nil
	}

	return Facility360UpsertOut{
		Success: false,
		Message: mapResp["Message"].(string),
		Record:  nil,
	}, nil
}

func (fetch *Fetch) LogMeInFacility360(facility Facility360Input) error {
	// set required inputs for the fetcher
	fetch.username = facility.Username
	fetch.password = facility.Password
	fetch.url = facility.Url
	return fetch.LogMeIn()
}

func (fetch Fetch) handleUpsertResponse(resp *http.Response) (Facility360UpsertOut, error) {
	data := make(JsonMap, 0)
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Facility360UpsertOut{}, err
	}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return Facility360UpsertOut{}, err
	}

	return Facility360UpsertOut{
		Success: true,
		Message: "",
		Record:  data,
	}, nil
}
