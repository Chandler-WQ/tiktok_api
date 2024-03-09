package service

import (
	"context"
	"regexp"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/util/http"
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

var re = regexp.MustCompile(`<script id="__UNIVERSAL_DATA_FOR_REHYDRATION__" type="application/json">([\w\W]*?)</script>`)

type UserExtractor struct {
	re   *regexp.Regexp
	body []byte
}

func NewUserExtractor(body []byte) UserExtractor {
	return UserExtractor{
		re:   re,
		body: body,
	}

}

func (e UserExtractor) ParseUserInfo() (*model.UserInfo, error) {
	strs := re.FindStringSubmatch(string(e.body))
	if len(strs) == 0 {
		return nil, errors.New("match __DEFAULT_SCOPE__ failed")
	}
	str := strs[len(strs)-1]
	node, err := sonic.GetFromString(str, "__DEFAULT_SCOPE__", "webapp.user-detail", "userInfo")
	if err != nil {
		return nil, errors.Wrapf(err, "Get __DEFAULT_SCOPE__->webapp.user-detail->userInfo failed: %s", str)
	}
	str, err = node.Raw()
	if err != nil {
		return nil, errors.Wrapf(err, "node.String")
	}
	userInfo := &model.UserInfo{}
	err = sonic.Unmarshal([]byte(str), userInfo)
	if err != nil {
		return nil, errors.Wrapf(err, "sonic.Unmarshal")
	}
	return userInfo, nil
}

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

func (cli UserClient) GetUserInfo(ctx context.Context, uniqueId string) (res *model.UserInfo, err error) {
	resp, err := cli.cliHTTP.WithCtx(ctx).
		SetQueryParam("is_from_webapp", "1").
		SetQueryParam("sender_device", "pc").
		SetHeader("cookie", cli.cookie).
		Get(host + "/@" + uniqueId)
	if err != nil {
		return nil, errors.Wrapf(err, "http get failed")
	}
	body := resp.Body()

	extractor := NewUserExtractor(body)
	userInfo, err := extractor.ParseUserInfo()
	if err != nil {
		return nil, errors.Wrapf(err, "extractor.ParseUserInfo")
	}
	return userInfo, nil
}
