package service_test

import (
	"context"
	"testing"

	"github.com/Chandler-WQ/tiktok_api/api/service"
	"github.com/Chandler-WQ/tiktok_api/util/log"
)

func TestGetUserInfo(t *testing.T) {
	cli := service.NewUserClient(cookie)
	ctx := context.Background()
	res, err := cli.GetUserInfo(ctx, "juleko_o")
	t.Logf("%v", err)
	t.Logf("%s", log.NewLogString(res))
}
