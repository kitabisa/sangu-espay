package espay

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"moul.io/http2curl"
)

type EspayClient struct {
	BaseUrl      string
	SignatureKey string
	Timeout      time.Duration
	Logger       Logger
	IsProduction bool
}

// IEspayClient contains interface(s) that you have to implement. Mainly a representation of espay endpoints that kitabisa will send the request to
type IEspayClient interface {
	CreateVA(req CreateVaRequest) (res CreateVaResponse, err error)
}

// Call : base method to call Espay
func (c *EspayClient) Call(method, path string, header map[string]string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = c.BaseUrl + path
	req, err := c.NewRequest(method, path, header, body)

	if err != nil {
		c.Logger.Info("Failed during NewRequest %v", err)
		return nil, err
	}

	return c.ExecuteRequest(req)
}

// NewClient : this function will always be called when the library is in use
func NewClient(espayClient EspayClient) IEspayClient {
	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
	}

	if espayClient.IsProduction {
		logOption.Pretty = false
	} else {
		logOption.Pretty = true
	}

	espayClient.Logger = *NewLogger(logOption)
	espayClient.Timeout = 1 * time.Minute
	return &espayClient
}

// ===================== HTTP CLIENT ================================================
var defHTTPBackoffInterval = 2 * time.Millisecond
var defHTTPMaxJitterInterval = 5 * time.Millisecond
var defHTTPRetryCount = 3

// getHTTPClient will get heimdall http client
func (c *EspayClient) getHTTPClient() *httpclient.Client {
	backoff := heimdall.NewConstantBackoff(defHTTPBackoffInterval, defHTTPMaxJitterInterval)
	retrier := heimdall.NewRetrier(backoff)

	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(c.Timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(defHTTPRetryCount),
	)
}

// NewRequest : send new request
func (c *EspayClient) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
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
func (c *EspayClient) ExecuteRequest(req *http.Request) ([]byte, error) {

	command, _ := http2curl.GetCurlCommand(req)
	start := time.Now()
	res, err := c.getHTTPClient().Do(req)
	if err != nil {
		c.Logger.Error("Request failed. Error : %v , Curl Request : %v", err, command)
		return nil, err
	}

	if !c.IsProduction {
		c.Logger.Info("Curl Request: %v ", command)
	}

	defer res.Body.Close()
	c.Logger.Info("Completed in %v", time.Since(start))

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error("Cannot read response body: %v ", err)
		return resBody, err
	}

	c.Logger.Info("Espay HTTP status response : %d", res.StatusCode)
	c.Logger.Info("Espay response body : %s", string(resBody))

	return resBody, err
}

// ===================== END HTTP CLIENT ================================================
