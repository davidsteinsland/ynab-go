package ynab

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"io"
	"bytes"
	"fmt"
)

type Client struct {
	BaseURL   *url.URL
	client *http.Client
	accessToken string

	AccountsService *AccountsService
	BudgetService *BudgetsService
	CategoriesService *CategoriesService
	MonthsService *MonthsService
	PayeeLocationsService *PayeeLocationsService
	PayeesService *PayeesService
	ScheduledTransactionsService *ScheduledTransactionsService
	TransactionsService *TransactionsService
	UserService *UserService
}

type service struct {
	client *Client
}

func (s service) do(method string, url string, reqBody interface{}, respBody interface{}) error {
	err := s.client.do(method, url, reqBody, respBody)
	if err != nil {
		return err
	}
	return nil
}

type ErrorResponse struct {
	Response *http.Response
	ErrorDetails ApiError `json:"error"`
}

type ApiError struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Detail string `json:"detail"`
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorDetails.Name, r.ErrorDetails.Detail)
}

var DefaultBaseURL = "https://api.youneedabudget.com/v1/"

func NewDefaultClient(accessToken string) *Client {
	baseUrl, _ := url.Parse(DefaultBaseURL)
	return NewClient(baseUrl, nil, accessToken)
}

func NewClient(baseUrl *url.URL, httpClient *http.Client, accessToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		BaseURL: baseUrl,
		client: httpClient,
		accessToken: accessToken,
	}

	client.AccountsService = &AccountsService{client}
	client.BudgetService = &BudgetsService{client}
	client.CategoriesService = &CategoriesService{client}
	client.MonthsService = &MonthsService{client}
	client.PayeeLocationsService = &PayeeLocationsService{client}
	client.PayeesService = &PayeesService{client}
	client.ScheduledTransactionsService = &ScheduledTransactionsService{client}
	client.TransactionsService = &TransactionsService{client}
	client.UserService = &UserService{client}

	return client
}

func (yc Client) newRequest(method string, relUrl string, reqBody interface{}) (*http.Request, error) {
	rel := &url.URL{Path: relUrl}
	resolvedUrl := yc.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if reqBody != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(reqBody)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, resolvedUrl.String(), buf)

	if err != nil {
		return nil, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", "Bearer " + yc.accessToken)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (yc Client) checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (yc Client) do(method string, relUrl string, reqBody interface{}, v interface{}) error {
	req, err := yc.newRequest(method, relUrl, reqBody)

	if err != nil {
		return err
	}

	resp, err := yc.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := yc.checkResponse(resp); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}

func StringVal(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
