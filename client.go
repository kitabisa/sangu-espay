package sangu_espay

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"moul.io/http2curl"
)

var ErrPendingTransaction = errors.New("Transaction is pending")

type Client struct {
	BaseUrl      string
	SignatureKey string
	LogLevel     int
	Timeout      time.Duration
	Logger       *log.Logger
	IsProduction bool
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	return Client{
		// LogLevel is the logging level used by the BRI library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		Timeout:      1 * time.Minute,
		Logger:       log.New(os.Stderr, "", log.LstdFlags),
		IsProduction: false,
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
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Request creation failed: ", err)
		}
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
	logLevel := c.LogLevel
	logger := c.Logger

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := c.getHTTPClient().Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return resBody, err
	}

	if logLevel > 2 {
		logger.Println("Espay HTTP status response: ", res.StatusCode)
		logger.Println("Espay body response: ", string(resBody))
	}

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)

	return resBody, err
}

func (c *Client) Call(method, path string, header map[string]string, body io.Reader) ([]byte, error) {
	req, err := c.NewRequest(method, path, header, body)

	if err != nil {
		return nil, err
	}

	return c.ExecuteRequest(req)
}

// ===================== END HTTP CLIENT ================================================
