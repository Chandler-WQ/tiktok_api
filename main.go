package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/api/service"
	"github.com/Chandler-WQ/tiktok_api/util/excel"
	"github.com/pkg/errors"
)

type Option struct {
	MsgToken string
	KeyWord  string
	Cookie   string
}

func (opt Option) Check() error {
	if opt.MsgToken == "" {
		return errors.Errorf("msg_token is empty")
	}
	if opt.KeyWord == "" {
		return errors.Errorf("KeyWord is empty")
	}
	if opt.Cookie == "" {
		return errors.Errorf("Cookie is empty")
	}
	return nil
}

func main() {
	opt := Option{}
	flag.StringVar(&opt.MsgToken, "msgtoken", "", "the msg token")
	flag.StringVar(&opt.KeyWord, "keyword", "", "the search KeyWord")
	flag.StringVar(&opt.Cookie, "cookie", "", "cookie")
	flag.Parse()

	ctx := context.Background()
	if err := opt.Check(); err != nil {
		flag.PrintDefaults()
		return
	}
	searchCli := service.NewSearchClient(opt.MsgToken, opt.Cookie)
	userCli := service.NewUserClient(opt.Cookie)
	resp, err := searchCli.SearchKeyword(ctx, opt.KeyWord)
	if err != nil {
		panic(err)
	}
	authorIDs := resp.CollectAuthorID()
	authorInfo := make([][]string, 0)
	authorInfo = append(authorInfo, model.UserInfoResp{}.ToSliceName())
	for id := range authorIDs {
		userInfo, err := userCli.GetUserInfo(ctx, id)
		if err != nil {
			fmt.Printf("get user info failed: %v,id:%s", err, id)
			continue
		}
		slice := userInfo.ToSlice()
		authorInfo = append(authorInfo, slice)
	}
	excelCli := excel.NewClient()
	err = excelCli.Create(ctx, "./"+opt.KeyWord+".xlsx", authorInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("爬取成功")
}
