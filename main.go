package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/api/service"
	"github.com/Chandler-WQ/tiktok_api/util/excel"
	"github.com/Chandler-WQ/tiktok_api/util/sets"
	"github.com/Chandler-WQ/tiktok_api/util/structinfo"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

var searchCli *service.SearchClient
var userCli *service.UserClient

type Option struct {
	KeyWord    string
	Cookie     string
	Cache      bool
	SaveCookie bool
	Offset     int64
	FindTimes  int64
	Debug      bool
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
	flag.Int64Var(&opt.Offset, "offset", 0, "the offset of keyword")
	flag.Int64Var(&opt.FindTimes, "find_times", 10, "How many times to find the end")
	flag.BoolVar(&opt.SaveCookie, "save_cookie", false, "whether just save cookie")
	flag.BoolVar(&opt.Debug, "debug", false, "whether open debug mode")
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
	searchCli = service.NewSearchClient(opt.Cookie)
	userCli = service.NewUserClient(opt.Cookie)

	if opt.Debug {
		searchCli.WithDebug(true)
		userCli.WithDebug(true)
	}
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
	offset := opt.Offset
	times := opt.FindTimes
	cookie := opt.Cookie
	keyword := opt.KeyWord
	authorInfo := make([][]string, 0)
	searchID := ""
	authorInfo = append(authorInfo, structinfo.ToSliceName(model.UserInfo{}))
	bar := progressbar.Default(times)
	userIDs := sets.StringSets{}
	for i := 0; i < int(times); i++ {
		fmt.Println()
		fmt.Printf("start search keyword:%s offset:%v searchID:%v \n", keyword, offset, searchID)
		authorInfoTmp, nextOffset, nextSearchID, err := SearchUserInfoByKeyword(ctx, cookie, keyword, searchID, offset, userIDs)
		if err != nil {
			_ = bar.Add(1)
			fmt.Printf("search keyword:%s offset:%v searchID:%v  error: %v\n", keyword, offset, searchID, err)
			continue
		}
		_ = bar.Add(1)
		fmt.Printf("search keyword:%s offset:%v searchID:%v  success\n", keyword, offset, searchID)
		authorInfo = append(authorInfo, authorInfoTmp...)
		offset = nextOffset
		searchID = nextSearchID
		// limit the speed of the search
		if i != int(times)-1 {
			time.Sleep(time.Second * 5)
		}
		fmt.Println()

	}

	if len(authorInfo) < 2 {
		return errors.New("find nothing")
	}
	excelCli := excel.NewClient()
	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrapf(err, "get current directory failed: %v", err)
	}
	timeStr := time.Now().Format("2006-01-02T15:04:05")
	path := dir + "/" + timeStr + "tt_" + opt.KeyWord + ".xlsx"
	err = excelCli.Create(ctx, path, authorInfo)
	if err != nil {
		return errors.Errorf("create excel client failed: %v", path)
	}

	fmt.Printf("all success,save in path:%s \n", path)
	return nil
}

func SearchUserInfoByKeyword(ctx context.Context, cookie, keyword, searchID string, offset int64, existUserIDs sets.StringSets) (info [][]string, nextOffset int64, nextSearchID string, err error) {
	resp, err := searchCli.SearchKeyword(ctx, keyword, searchID, offset)
	if err != nil {
		return nil, 0, "", errors.Wrapf(err, "search keyword:%s failed: %v", keyword, err)
	}
	authorInfo := make([][]string, 0)
	authorIDs := resp.CollectAuthorID()
	for id := range authorIDs {
		if existUserIDs.Contains(id) {
			fmt.Printf("author %s repeat,skip \n", id)
			continue
		}
		userInfo, err := userCli.GetUserInfo(ctx, id)
		if err != nil {
			fmt.Printf("find author %s info fail.reason %v \n", id, err)
			continue
		}
		fmt.Printf("find author %s info Success\n ", id)
		slice := structinfo.ToSlice(userInfo)
		authorInfo = append(authorInfo, slice)
		existUserIDs.Add(id)
	}
	return authorInfo, int64(resp.Cursor), resp.Extra.Logid, nil
}
