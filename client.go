package sangu_espay

import (
	"encoding/json"
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
func (c *Client) ExecuteRequest(req *http.Request, v interface{}, vErr interface{}) error {
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
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("Espay HTTP status response: ", res.StatusCode)
		logger.Println("Espay body response: ", string(resBody))
	}

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)

	if res.StatusCode == 404 {
		return errors.New("invalid url")
	}

	if res.StatusCode == 204 {
		return errors.New("204: empty response")
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			if vErr != nil {
				err = json.Unmarshal(resBody, &vErr)
			}

			if res.StatusCode == http.StatusOK {
				return nil
			}
			return err
		}
	}

	return nil
}

func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}, vErr interface{}) error {
	req, err := c.NewRequest(method, path, header, body)

	if err != nil {
		return err
	}

	return c.ExecuteRequest(req, v, vErr)
}

// ===================== END HTTP CLIENT ================================================
