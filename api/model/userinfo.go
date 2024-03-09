package model

type UserInfo struct {
	User  User  `json:"user,omitempty"`
	Stats Stats `json:"stats,omitempty"`
	// ItemList []any `json:"itemList,omitempty"`
}

type User struct {
	ID                     string           `json:"id,omitempty"`
	ShortID                string           `json:"shortId,omitempty"`
	UniqueID               string           `json:"uniqueId,omitempty"`
	Nickname               string           `json:"nickname,omitempty"`
	AvatarLarger           string           `json:"avatarLarger,omitempty"`
	AvatarMedium           string           `json:"avatarMedium,omitempty"`
	AvatarThumb            string           `json:"avatarThumb,omitempty"`
	Signature              string           `json:"signature,omitempty"`
	CreateTime             int              `json:"createTime,omitempty"`
	Verified               bool             `json:"verified,omitempty"`
	SecUID                 string           `json:"secUid,omitempty"`
	Ftc                    bool             `json:"ftc,omitempty"`
	Relation               int              `json:"relation,omitempty"`
	OpenFavorite           bool             `json:"openFavorite,omitempty"`
	BioLink                BioLink          `json:"bioLink,omitempty"`
	CommentSetting         int              `json:"commentSetting,omitempty"`
	CommerceUserInfo       CommerceUserInfo `json:"commerceUserInfo,omitempty"`
	DuetSetting            int              `json:"duetSetting,omitempty"`
	StitchSetting          int              `json:"stitchSetting,omitempty"`
	PrivateAccount         bool             `json:"privateAccount,omitempty"`
	Secret                 bool             `json:"secret,omitempty"`
	IsADVirtual            bool             `json:"isADVirtual,omitempty"`
	RoomID                 string           `json:"roomId,omitempty"`
	UniqueIDModifyTime     int              `json:"uniqueIdModifyTime,omitempty"`
	TtSeller               bool             `json:"ttSeller,omitempty"`
	Region                 string           `json:"region"`
	ProfileTab             ProfileTab       `json:"profileTab,omitempty"`
	FollowingVisibility    int              `json:"followingVisibility,omitempty"`
	RecommendReason        string           `json:"recommendReason,omitempty"`
	NowInvitationCardURL   string           `json:"nowInvitationCardUrl,omitempty"`
	NickNameModifyTime     int              `json:"nickNameModifyTime,omitempty"`
	IsEmbedBanned          bool             `json:"isEmbedBanned,omitempty"`
	CanExpPlaylist         bool             `json:"canExpPlaylist,omitempty"`
	ProfileEmbedPermission int              `json:"profileEmbedPermission,omitempty"`
	Language               string           `json:"language,omitempty"`
	EventList              []interface{}    `json:"eventList,omitempty"`
	SuggestAccountBind     bool             `json:"suggestAccountBind,omitempty"`
}

type BioLink struct {
	Link string `json:"link,omitempty"`
	Risk int    `json:"risk,omitempty"`
}

type CommerceUserInfo struct {
	CommerceUser bool `json:"commerceUser"`
}

type ProfileTab struct {
	ShowPlayListTab bool `json:"showPlayListTab"`
}
