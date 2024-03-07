package model

import (
	"reflect"

	"github.com/spf13/cast"
)

type UserInfoResp struct {
	Extra       Extra     `json:"extra"`
	LogPb       LogPb     `json:"log_pb"`
	ShareMeta   ShareMeta `json:"shareMeta"`
	StatusCode  int       `json:"statusCode"`
	StatusCode0 int       `json:"status_code"`
	StatusMsg   string    `json:"status_msg"`
	UserInfo    UserInfo  `json:"userInfo"`
}

func (u UserInfoResp) ToSliceName() []string {
	v := reflect.ValueOf(u.UserInfo)
	t := v.Type()
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			fields = walkStructName(field.Type, fields)
		} else {
			fields = append(fields, cast.ToString(field.Name))
		}
	}
	return fields
}

func walkStructName(t reflect.Type, fields []string) []string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			fields = walkStructName(field.Type, fields)
		} else {
			fields = append(fields, cast.ToString(field.Name))
		}
	}
	return fields
}

func (u UserInfoResp) ToSlice() []string {
	v := reflect.ValueOf(u.UserInfo)
	t := v.Type()
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i)
		if value.Kind() == reflect.Struct {
			fields = walkStruct(value, fields)
		} else {
			fields = append(fields, cast.ToString(value.Interface()))
		}
	}
	return fields
}

func walkStruct(v reflect.Value, fields []string) []string {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)

		if value.Kind() == reflect.Struct {
			fields = walkStruct(value, fields)
		} else {
			fields = append(fields, cast.ToString(value.Interface()))
		}
	}
	return fields
}

type ShareMeta struct {
	Desc  string `json:"desc"`
	Title string `json:"title"`
}

type CommerceUserInfo struct {
	CommerceUser bool `json:"commerceUser"`
}

type ProfileTab struct {
	ShowPlayListTab bool `json:"showPlayListTab"`
}

type User struct {
	AvatarLarger           string           `json:"avatarLarger",CE:"AvatarLarger1"`
	AvatarMedium           string           `json:"avatarMedium"`
	AvatarThumb            string           `json:"avatarThumb"`
	CanExpPlaylist         bool             `json:"canExpPlaylist"`
	CommentSetting         int              `json:"commentSetting"`
	CommerceUserInfo       CommerceUserInfo `json:"commerceUserInfo"`
	DuetSetting            int              `json:"duetSetting"`
	FollowingVisibility    int              `json:"followingVisibility"`
	Ftc                    bool             `json:"ftc"`
	ID                     string           `json:"id"`
	IsADVirtual            bool             `json:"isADVirtual"`
	IsEmbedBanned          bool             `json:"isEmbedBanned"`
	NickNameModifyTime     int              `json:"nickNameModifyTime"`
	Nickname               string           `json:"nickname"`
	OpenFavorite           bool             `json:"openFavorite"`
	PrivateAccount         bool             `json:"privateAccount"`
	ProfileEmbedPermission int              `json:"profileEmbedPermission"`
	ProfileTab             ProfileTab       `json:"profileTab"`
	Relation               int              `json:"relation"`
	SecUID                 string           `json:"secUid"`
	Secret                 bool             `json:"secret"`
	Signature              string           `json:"signature"`
	StitchSetting          int              `json:"stitchSetting"`
	TtSeller               bool             `json:"ttSeller"`
	UniqueID               string           `json:"uniqueId"`
	Verified               bool             `json:"verified"`
}

type UserInfo struct {
	Stats1 UserInfoStats `json:"stats"`
	User   User          `json:"user"`
}

type UserInfoStats struct {
	DiggCount      int `json:"diggCount"`
	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
	FriendCount    int `json:"friendCount"`
	Heart          int `json:"heart"`
	HeartCount     int `json:"heartCount"`
	VideoCount     int `json:"videoCount"`
}
