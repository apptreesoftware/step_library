package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Facility360Input struct {
	Username string
	Password string
	Url      string
	Endpoint string
}

type Facility360CreateIn struct {
	Facility360Input
	Record JsonMap
}

type Facility360UpdateIn struct {
	Facility360CreateIn
	Id int
}

type Facility360UpsertOut struct {
	Success bool
	Message string
	Record  JsonMap
}

type FetchInput struct {
	Username string
	Password string
	Url      string
	Endpoint string
	Filter   string
	Select   string
	Expand   string
	Skip     int
	Top      int
}

type FetchSingleOutput struct {
	Found  bool
	Record JsonMap
}

type FetchListOutputs struct {
	Count   int
	Records []jsoniter.RawMessage
}

type FetchAndQueueInput struct {
	FetchInput
	ChildPath string
	Workflow  string
}

type PagedResponse interface {
	GetCurrent() []jsoniter.RawMessage
	HasMoreResults() bool
	NextPageUrl() url.URL
}

type NextLinkPagedResponse struct {
	Value    []jsoniter.RawMessage `json:"value"`
	NextLink string                `json:"@odata.nextLink"`
}

func (r *NextLinkPagedResponse) GetCurrent() []jsoniter.RawMessage {
	return r.Value
}

func (r *NextLinkPagedResponse) HasMoreResults() bool {
	_, err := url.Parse(r.NextLink)
	if err != nil {
		return false
	}
	return r.NextLink != ""
}

func (r *NextLinkPagedResponse) NextPageUrl() url.URL {
	uri, _ := url.Parse(r.NextLink)
	uri.Scheme = "https"
	return *uri
}

type WorkOrderPagedResponse struct {
	Value      []jsoniter.RawMessage `json:"value"`
	BatchSize  int
	currentUrl url.URL
}

func NewWorkOrderPagedResponse(batchSize int, currentUrl url.URL) *WorkOrderPagedResponse {
	return &WorkOrderPagedResponse{BatchSize: batchSize, currentUrl: currentUrl}
}

func (r *WorkOrderPagedResponse) GetCurrent() []jsoniter.RawMessage {
	return r.Value
}

func (r *WorkOrderPagedResponse) HasMoreResults() bool {
	if len(r.Value) == r.BatchSize && r.BatchSize > 0 {
		return true
	}
	return false
}

func (r WorkOrderPagedResponse) NextPageUrl() url.URL {
	newUrl := r.currentUrl
	q := newUrl.Query()
	currentSkip := q.Get(skipKey)

	nextSkip := r.BatchSize
	if currentSkip != "" {
		skipInt, _ := strconv.ParseInt(currentSkip, 10, 64)
		nextSkip = int(skipInt) + r.BatchSize
	}
	q.Set(skipKey, fmt.Sprintf("%d", nextSkip))

	newUrl.RawQuery = q.Encode()
	finalUrl := odataUrl(newUrl)
	return *finalUrl
}

func odataUrl(url url.URL) *url.URL {
	queryStr := url.RawQuery
	queryStr = strings.Replace(queryStr, "%24skip", "$skip", -1)
	queryStr = strings.Replace(queryStr, "%24filter", "$filter", -1)
	queryStr = strings.Replace(queryStr, "%24top", "$top", -1)
	queryStr = strings.Replace(queryStr, "%24select", "$select", -1)
	queryStr = strings.Replace(queryStr, "%24expand", "$expand", -1)
	url.RawQuery = queryStr
	return &url
}

type LoginResponse struct {
	Result  bool     `json:"Result"`
	Message string   `json:"Message"`
	Auth    AuthItem `json:"Item"`
}

type AuthItem struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      string `json:".expires"`
}

func (i AuthItem) GetExpiration() (time.Time, error) {
	return time.Parse(time.RFC1123, strings.TrimSpace(i.Expires))
}

type JsonMap map[string]interface{}

func (mp *JsonMap) String() string {
	data, err := jsoniter.Marshal(mp)
	if err != nil {
		return ""
	}
	return string(data)
}
