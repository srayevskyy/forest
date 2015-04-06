package rat

import "net/http"

// ApiTesting provides functions to call a REST api and validate its responses.
type ApiTesting struct {
	BaseURL string
	client  *http.Client
}

// NewClient returns a new ApiTesting that can be used to send Http requests.
func NewClient(baseURL string, httpClient *http.Client) *ApiTesting {
	return &ApiTesting{
		BaseURL: baseURL,
		client:  httpClient,
	}
}

// GET sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) GET(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("GET", a.BaseURL+config.pathAndQuery(), nil)
	if err != nil {
		t.Errorf("invalid Url:%s", a.BaseURL+config.pathAndQuery())
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// POST sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) POST(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("POST", a.BaseURL+config.pathAndQuery(), config.BodyReader)
	if err != nil {
		t.Errorf("invalid Url:%s", a.BaseURL+config.pathAndQuery())
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// PUT sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) PUT(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("PUT", a.BaseURL+config.pathAndQuery(), config.BodyReader)
	if err != nil {
		t.Errorf("invalid Url:%s", a.BaseURL+config.pathAndQuery())
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// DELETE sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) DELETE(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("DELETE", a.BaseURL+config.pathAndQuery(), nil)
	if err != nil {
		t.Errorf("invalid Url:%s", a.BaseURL+config.pathAndQuery())
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// PATCH sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) PATCH(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("PATCH", a.BaseURL+config.pathAndQuery(), config.BodyReader)
	if err != nil {
		t.Errorf("invalid Url:%s", a.BaseURL+config.pathAndQuery())
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}
