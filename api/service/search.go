package service

import (
	"context"
	"encoding/json"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/util/http"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type SearchClient struct {
	cliHTTP *http.Client
	cookie  string
	debug   bool
}

func (cli SearchClient) WithDebug(debug bool) {
	cli.debug = debug
}

func NewSearchClient(cookie string) *SearchClient {
	return &SearchClient{
		cliHTTP: http.NewDftClient(),
		cookie:  cookie,
	}
}

func (cli SearchClient) SearchKeyword(ctx context.Context, keyword, searchId string, offset int64) (res *model.SearchResp, err error) {
	resp, err := cli.cliHTTP.WithCtx(ctx).
		SetQueryParam("WebIdLastTime", "0").
		SetQueryParam("aid", "1988").
		SetQueryParam("app_name", "tiktok_web").
		SetQueryParam("keyword", keyword).
		SetQueryParam("search_id", searchId).
		SetQueryParam("offset", cast.ToString(offset)).
		SetHeader("authority", "www.tiktok.com").
		SetHeader("cookie", cli.cookie).SetDebug(cli.debug).
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
