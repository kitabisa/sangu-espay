package sangu_espay

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"moul.io/http2curl"
)

type Client struct {
	BaseUrl      string
	SignatureKey string
	Timeout      time.Duration
	Logger       Logger
	IsProduction bool
}

// NewClient : this function will always be called when the library is in use
func NewClient(isProd bool) Client {

	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
	}

	if isProd {
		logOption.Pretty = false
	} else {
		logOption.Pretty = true
	}

	logger := NewLogger(logOption)

	return Client{
		Timeout:      1 * time.Minute,
		Logger:       *logger,
		IsProduction: isProd,
	}
}

// ===================== HTTP CLIENT ================================================
var defHTTPBackoffInterval = 2 * time.Millisecond
var defHTTPMaxJitterInterval = 5 * time.Millisecond
var defHTTPRetryCount = 3

// getHTTPClient will get heimdall http client
func (c *Client) getHTTPClient() *httpclient.Client {
	backoff := heimdall.NewConstantBackoff(defHTTPBackoffInterval, defHTTPMaxJitterInterval)
	retrier := heimdall.NewRetrier(backoff)

	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(c.Timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(defHTTPRetryCount),
	)
}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	log := c.Logger

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		log.Error("Error during NewRequest %v", err)
		return nil, err
	}

	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}

	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request) ([]byte, error) {

	start := time.Now()
	command, _ := http2curl.GetCurlCommand(req)
	res, err := c.getHTTPClient().Do(req)
	if err != nil {
		c.Logger.Error("Request failed. Error : %v , Curl Request : %v", err, command)
		return nil, err
	}

	if !c.IsProduction {
		c.Logger.Info("Curl Request: %v ", command)
	}

	defer res.Body.Close()
	c.Logger.Info("Completed in %d", time.Since(start))

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error("Cannot read response body: %v ", err)
		return resBody, err
	}

	c.Logger.Info("Espay HTTP status response : %d", res.StatusCode)
	c.Logger.Info("Espay response body : %s", string(resBody))

	return resBody, err
}

func (c *Client) Call(method, path string, header map[string]string, body io.Reader) ([]byte, error) {
	req, err := c.NewRequest(method, path, header, body)

	if err != nil {
		c.Logger.Info("Failed during NewRequest %v", err)
		return nil, err
	}

	return c.ExecuteRequest(req)
}

// ===================== END HTTP CLIENT ================================================
