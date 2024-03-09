package service

import (
	"context"
	"encoding/json"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/util/http"
	"github.com/pkg/errors"
)

type SearchClient struct {
	cliHTTP *http.Client
	cookie  string
}

func NewSearchClient(cookie string) *SearchClient {
	return &SearchClient{
		cliHTTP: http.NewDftClient(),
		cookie:  cookie,
	}
}

func (cli SearchClient) SearchKeyword(ctx context.Context, keyword string) (res *model.SearchResp, err error) {
	resp, err := cli.cliHTTP.WithCtx(ctx).
		SetQueryParam("WebIdLastTime", "0").
		SetQueryParam("aid", "1988").
		SetQueryParam("app_name", "tiktok_web").
		SetQueryParam("keyword", keyword).
		SetHeader("authority", "www.tiktok.com").
		SetHeader("cookie", cli.cookie).
		Get(host + "/api/search/general/full/?")
	if err != nil {
		return nil, errors.Wrapf(err, "http get failed")
	}
	res = &model.SearchResp{}
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
