package model_test

import (
	"testing"

	"github.com/Chandler-WQ/tiktok_api/api/model"
	"github.com/Chandler-WQ/tiktok_api/util/log"
	"github.com/Chandler-WQ/tiktok_api/util/structinfo"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	var str = `{"user":{"id":"6791929390982923269","uniqueId":"sdcostumeguy","nickname":"sdcostumeguy","avatarLarger":"https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/7134b210b4d510a6ebc4bd77c1a0dad6~c5_1080x1080.jpeg?lk3s=a5d48078&x-expires=1710140400&x-signature=PU6zJpsU1tXm12vkBYhw7T9kmaw%3D","avatarMedium":"https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/7134b210b4d510a6ebc4bd77c1a0dad6~c5_720x720.jpeg?lk3s=a5d48078&x-expires=1710140400&x-signature=F7xDrLfpnnXMArWj4Em4zS3N50Y%3D","avatarThumb":"https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/7134b210b4d510a6ebc4bd77c1a0dad6~c5_100x100.jpeg?lk3s=a5d48078&x-expires=1710140400&x-signature=8t4K1qIIMf5PBesUsePTYaD%2F95w%3D","signature":"sdcostumeguy.  San Diego CA","secUid":"MS4wLjABAAAAoNVJce0TLEnKKaBkZ6Tu-hej9HxISTLVpDm-dQnjA69n6NEA0lsm81UC1Q3Kq6yH","bioLink":{},"commerceUserInfo":{"commerceUser":false},"duetSetting":3,"stitchSetting":3,"region":"US","profileTab":{"showPlayListTab":false},"followingVisibility":2,"canExpPlaylist":true,"profileEmbedPermission":1,"language":"en"},"stats":{"diggCount":0,"shareCount":0,"commentCount":0,"playCount":0,"collectCount":0}}`
	info := &model.UserInfo{}
	err := sonic.UnmarshalString(str, info)
	assert.Nil(t, err)
	t.Logf("%s", log.NewLogString(info))
	t.Logf("ToSlice:%s", log.NewLogString(structinfo.ToSlice(info)))
	t.Logf("ToSliceName:%s", log.NewLogString(structinfo.ToSliceName(info)))
}
