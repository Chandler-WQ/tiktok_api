package service

import (
	"context"
	"encoding/json"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/util/http"
	"github.com/pkg/errors"
)

type UserClient struct {
	cliHTTP *http.Client
	cookie  string
}

func NewUserClient(cookie string) *UserClient {
	return &UserClient{
		cliHTTP: http.NewDftClient(),
		cookie:  cookie,
	}
}

func (cli UserClient) GetUserInfo(ctx context.Context, uniqueId string) (res *model.UserInfoResp, err error) {
	resp, err := cli.cliHTTP.WithCtx(ctx).
		SetQueryParam("WebIdLastTime", "0").
		SetQueryParam("aid", "1988").
		SetQueryParam("app_name", "tiktok_web").
		SetQueryParam("channel", "tiktok_web").
		SetQueryParam("uniqueId", uniqueId).
		SetHeader("authority", "www.tiktok.com").
		SetHeader("cookie", cli.cookie).SetDebug(true).
		Get(host + "/api/user/detail/?")
	if err != nil {
		return nil, errors.Wrapf(err, "http get failed")
	}
	res = &model.UserInfoResp{}
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
