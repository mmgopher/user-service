package helpers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// HTTPProvider provides functionality to do http requests
type HTTPProvider interface {
	// DoRequest sends http request and returns status code and response body
	DoRequest(method string, url string, urlParams map[string]string, headerParams map[string]string, requestBody []byte) (
		int, []byte, error)
}

// HTTPService implements HTTPProvider
type HTTPService struct {
	httpClient *http.Client
}

// NewHTTPService creates new HTTPProvider instance
func NewHTTPService(httpClient *http.Client) *HTTPService {
	return &HTTPService{
		httpClient: httpClient,
	}
}

// DoRequest sends http request and returns status code and response body
func (s HTTPService) DoRequest(method string,
	url string,
	urlParams map[string]string,
	headerParams map[string]string,
	requestBody []byte) (int, []byte, error) {

	var body io.Reader
	if requestBody != nil {
		body = bytes.NewReader(requestBody)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error creating new http request")
	}

	q := req.URL.Query()
	for k, v := range urlParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range headerParams {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)

	if resp != nil {
		//nolint
		defer resp.Body.Close()
	}

	if err != nil {
		return 0, nil, errors.Wrapf(err, "error doing %s request to endpoint: %s", method, url)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error reading response body")
	}

	return resp.StatusCode, respBody, nil
}
