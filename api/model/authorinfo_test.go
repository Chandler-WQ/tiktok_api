package model_test

import (
	"testing"

	"github.com/Chandler-WQ/tiktok_api/api/model"
)

func TestToSlice(t *testing.T) {
	userInfo := model.UserInfoResp{
		Extra:       model.Extra{},
		LogPb:       model.LogPb{},
		ShareMeta:   model.ShareMeta{},
		StatusCode:  0,
		StatusCode0: 0,
		StatusMsg:   "",
		UserInfo: model.UserInfo{
			Stats1: model.UserInfoStats{
				DiggCount:      200,
				FollowerCount:  300,
				FollowingCount: 0,
				FriendCount:    0,
				Heart:          0,
				HeartCount:     0,
				VideoCount:     0,
			},
			User: model.User{
				AvatarLarger:           "a",
				AvatarMedium:           "sds",
				AvatarThumb:            "vcv",
				CanExpPlaylist:         false,
				CommentSetting:         1000,
				CommerceUserInfo:       model.CommerceUserInfo{},
				DuetSetting:            0,
				FollowingVisibility:    0,
				Ftc:                    false,
				ID:                     "",
				IsADVirtual:            false,
				IsEmbedBanned:          false,
				NickNameModifyTime:     0,
				Nickname:               "",
				OpenFavorite:           false,
				PrivateAccount:         false,
				ProfileEmbedPermission: 0,
				ProfileTab:             model.ProfileTab{},
				Relation:               0,
				SecUID:                 "",
				Secret:                 false,
				Signature:              "",
				StitchSetting:          0,
				TtSeller:               false,
				UniqueID:               "",
				Verified:               false,
			},
		},
	}
	slice := userInfo.ToSlice()
	sliceName := userInfo.ToSliceName()
	t.Logf("%v %v", slice, len(slice))
	t.Logf("%v %v", sliceName, len(sliceName))
}
