package service_test

import (
	"context"
	"testing"

	"github.com/Chandler-WQ/tiktok_api/api/service"
	"github.com/Chandler-WQ/tiktok_api/util/log"
)

func TestSearchKeyword(t *testing.T) {
	cli := service.NewSearchClient(msToken, cookie)
	ctx := context.Background()
	res, err := cli.SearchKeyword(ctx, "china")
	t.Logf("%v", err)
	t.Logf("%s", log.NewLogString(res))
}
