package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const timeout = time.Hour

type Config struct {
	// ApiKey for the api
	ApiKey string

	// Authentication cookie value
	Cca string
}

type Client struct {
	// HttpClient is client to use
	HttpClient *http.Client
	config     *Config
}

func NewHttpClient(config *Config) *Client {
	client := http.Client{Timeout: timeout}
	return &Client{
		HttpClient: &client,
		config:     config,
	}
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(hostname string, method string, path string, body io.Reader, data interface{}) error {
	apiRoot := hostname + "/api/v1/"
	url := apiRoot + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", c.config.ApiKey)
	req.AddCookie(&http.Cookie{Name: "cca", Value: c.config.Cca})

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if res.StatusCode == 403 {
			return fmt.Errorf("login token expired please login again")
		}
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return fmt.Errorf("lailed fetching %s, error: %s, statusCode: %s", url, errRes.Message, fmt.Sprint(res.StatusCode))
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Print("Errorrrrrr")
		return err
	}

	return nil
}
