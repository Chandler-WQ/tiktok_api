package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/api/service"
	"github.com/Chandler-WQ/tiktok_api/util/excel"
	"github.com/Chandler-WQ/tiktok_api/util/structinfo"
	"github.com/pkg/errors"
)

type Option struct {
	KeyWord    string
	Cookie     string
	Cache      bool
	SaveCookie bool
}

func (opt Option) Check() error {
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
	flag.StringVar(&opt.KeyWord, "keyword", "", "keyword of search")
	flag.StringVar(&opt.Cookie, "cookie", "", "the cookie of request info")
	flag.BoolVar(&opt.Cache, "cache", true, "whether use cache cookie")
	flag.BoolVar(&opt.SaveCookie, "save_cookie", false, "whether just save cookie")
	flag.Parse()
	if opt.Cookie != "" && opt.SaveCookie {
		_ = SetCookies(opt.Cookie)
		return
	}

	var useCache bool
	if opt.Cache && opt.Cookie == "" {
		cookie, err := GetCookies()
		if err != nil {
			fmt.Println("please with cookie")
			return
		}
		useCache = true
		opt.Cookie = cookie
	}
	if opt.Cookie != "" && !useCache {
		_ = SetCookies(opt.Cookie)
	}
	if err := opt.Check(); err != nil {
		flag.PrintDefaults()
		fmt.Println("please with some arguments")
		return
	}
	ctx := context.Background()
	err := DoFind(ctx, &opt)
	if err != nil {
		panic(err)
	}

	return
}

func SetCookies(cookie string) error {
	dir := os.TempDir()
	path := dir + "/tt_cookie"
	err := ioutil.WriteFile(path, []byte(cookie), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("cookie save in %s success\n", path)
	return nil
}

func GetCookies() (string, error) {
	dir := os.TempDir()
	cookie, err := ioutil.ReadFile(dir + "/tt_cookie")
	if err != nil {
		return "", err
	}
	return string(cookie), nil
}

func DoFind(ctx context.Context, opt *Option) error {
	searchCli := service.NewSearchClient(opt.Cookie)
	userCli := service.NewUserClient(opt.Cookie)
	resp, err := searchCli.SearchKeyword(ctx, opt.KeyWord)
	if err != nil {
		return errors.Wrapf(err, "search keyword:%s failed: %v", opt.KeyWord, err)
	}
	authorIDs := resp.CollectAuthorID()
	authorInfo := make([][]string, 0)
	authorInfo = append(authorInfo, structinfo.ToSliceName(model.UserInfo{}))
	for id := range authorIDs {
		userInfo, err := userCli.GetUserInfo(ctx, id)
		if err != nil {
			fmt.Printf("find author %s info fail.reason %v \n", id, err)
			continue
		}
		fmt.Printf("find author %s info Success\n ", id)
		slice := structinfo.ToSlice(userInfo)
		authorInfo = append(authorInfo, slice)
	}
	if len(authorInfo) < 2 {
		return errors.New("find nothing")
	}
	excelCli := excel.NewClient()
	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrapf(err, "get current directory failed: %v", err)
	}
	path := dir + "/tt_" + opt.KeyWord + ".xlsx"
	err = excelCli.Create(ctx, path, authorInfo)
	if err != nil {
		return errors.Errorf("create excel client failed: %v", path)
	}

	fmt.Printf("all success,save in path:%s \n", path)
	return nil
}
