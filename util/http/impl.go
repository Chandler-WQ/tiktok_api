package http

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

var dftClient *Client

type Client struct {
	*resty.Client
}

type Request struct {
	*resty.Request
}

func init() {
	dftHTTPClient := resty.New().SetTimeout(120 * time.Second).SetRetryCount(3)

	mwErr := func(client *resty.Client, resp *resty.Response) error {
		if !resp.IsSuccess() {
			return fmt.Errorf("status_code:%v,body:%s", resp.StatusCode(), string(resp.Body()))
		}
		return nil
	}
	dftHTTPClient.OnAfterResponse(mwErr)
	dftClient = &Client{
		Client: dftHTTPClient,
	}
}

func NewDftClient() *Client {
	return dftClient
}

func (c *Client) WithCtx(ctx context.Context) *Request {
	return &Request{Request: c.R().SetContext(ctx)}
}

func (c *Request) SetContentTypeJSON() *Request {
	return &Request{Request: c.SetHeader("Content-Type", "application/json")}
}

func (c *Request) SetContentTypeText() *Request {
	return &Request{Request: c.SetHeader("Content-Type", "text/plain")}
}
