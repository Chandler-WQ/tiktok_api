package model

import "github.com/Chandler-WQ/tiktok_api/util/sets"

type SearchResp struct {
	StatusCode         int                `json:"status_code"`
	Data               []Data             `json:"data"`
	Qc                 string             `json:"qc"`
	Cursor             int                `json:"cursor"`
	HasMore            int                `json:"has_more"`
	AdInfo             AdInfo             `json:"ad_info"`
	Extra              Extra              `json:"extra"`
	LogPb              LogPb              `json:"log_pb"`
	GlobalDoodleConfig GlobalDoodleConfig `json:"global_doodle_config"`
	Backtrace          string             `json:"backtrace"`
}

func (s SearchResp) CollectAuthorID() sets.StringSets {
	set := sets.StringSets{}
	for _, data := range s.Data {
		if data.Item.Author.UniqueID != "" {
			set.Add(data.Item.Author.UniqueID)
		}
	}
	return set
}

type Video struct {
	ID            string   `json:"id"`
	Height        int      `json:"height"`
	Width         int      `json:"width"`
	Duration      int      `json:"duration"`
	Ratio         string   `json:"ratio"`
	Cover         string   `json:"cover"`
	OriginCover   string   `json:"originCover"`
	DynamicCover  string   `json:"dynamicCover"`
	PlayAddr      string   `json:"playAddr"`
	DownloadAddr  string   `json:"downloadAddr"`
	ShareCover    []string `json:"shareCover"`
	ReflowCover   string   `json:"reflowCover"`
	Bitrate       int      `json:"bitrate"`
	EncodedType   string   `json:"encodedType"`
	Format        string   `json:"format"`
	VideoQuality  string   `json:"videoQuality"`
	EncodeUserTag string   `json:"encodeUserTag"`
}

type Author struct {
	ID              string `json:"id"`
	UniqueID        string `json:"uniqueId"`
	Nickname        string `json:"nickname"`
	AvatarThumb     string `json:"avatarThumb"`
	AvatarMedium    string `json:"avatarMedium"`
	AvatarLarger    string `json:"avatarLarger"`
	Signature       string `json:"signature"`
	Verified        bool   `json:"verified"`
	SecUID          string `json:"secUid"`
	Secret          bool   `json:"secret"`
	Ftc             bool   `json:"ftc"`
	Relation        int    `json:"relation"`
	OpenFavorite    bool   `json:"openFavorite"`
	CommentSetting  int    `json:"commentSetting"`
	DuetSetting     int    `json:"duetSetting"`
	StitchSetting   int    `json:"stitchSetting"`
	PrivateAccount  bool   `json:"privateAccount"`
	DownloadSetting int    `json:"downloadSetting"`
}

type Music struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	PlayURL     string `json:"playUrl"`
	CoverThumb  string `json:"coverThumb"`
	CoverMedium string `json:"coverMedium"`
	CoverLarge  string `json:"coverLarge"`
	AuthorName  string `json:"authorName"`
	Original    bool   `json:"original"`
	Duration    int    `json:"duration"`
	Album       string `json:"album"`
}

type Challenges struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	ProfileThumb  string `json:"profileThumb"`
	ProfileMedium string `json:"profileMedium"`
	ProfileLarger string `json:"profileLarger"`
	CoverThumb    string `json:"coverThumb"`
	CoverMedium   string `json:"coverMedium"`
	CoverLarger   string `json:"coverLarger"`
	IsCommerce    bool   `json:"isCommerce"`
}

type Stats struct {
	DiggCount    int `json:"diggCount"`
	ShareCount   int `json:"shareCount"`
	CommentCount int `json:"commentCount"`
	PlayCount    int `json:"playCount"`
	CollectCount int `json:"collectCount"`
}

type DuetInfo struct {
	DuetFromID string `json:"duetFromId"`
}

type TextExtra struct {
	AwemeID      string `json:"awemeId"`
	Start        int    `json:"start"`
	End          int    `json:"end"`
	HashtagName  string `json:"hashtagName"`
	HashtagID    string `json:"hashtagId"`
	Type         int    `json:"type"`
	UserID       string `json:"userId"`
	IsCommerce   bool   `json:"isCommerce"`
	UserUniqueID string `json:"userUniqueId"`
	SecUID       string `json:"secUid"`
	SubType      int    `json:"subType"`
}
type AuthorStats struct {
	FollowingCount int `json:"followingCount"`
	FollowerCount  int `json:"followerCount"`
	HeartCount     int `json:"heartCount"`
	VideoCount     int `json:"videoCount"`
	DiggCount      int `json:"diggCount"`
	Heart          int `json:"heart"`
}
type Item struct {
	ID                string           `json:"id"`
	Desc              string           `json:"desc"`
	CreateTime        int              `json:"createTime"`
	Video             Video            `json:"video"`
	Author            Author           `json:"author"`
	Music             Music            `json:"music"`
	Challenges        []Challenges     `json:"challenges"`
	Stats             Stats            `json:"stats"`
	DuetInfo          DuetInfo         `json:"duetInfo"`
	OriginalItem      bool             `json:"originalItem"`
	OfficalItem       bool             `json:"officalItem"`
	TextExtra         []TextExtra      `json:"textExtra"`
	Secret            bool             `json:"secret"`
	ForFriend         bool             `json:"forFriend"`
	Digged            bool             `json:"digged"`
	ItemCommentStatus int              `json:"itemCommentStatus"`
	ShowNotPass       bool             `json:"showNotPass"`
	Vl1               bool             `json:"vl1"`
	ItemMute          bool             `json:"itemMute"`
	AuthorStats       AuthorStats      `json:"authorStats"`
	PrivateItem       bool             `json:"privateItem"`
	DuetEnabled       bool             `json:"duetEnabled"`
	StitchEnabled     bool             `json:"stitchEnabled"`
	ShareEnabled      bool             `json:"shareEnabled"`
	IsAd              bool             `json:"isAd"`
	Collected         bool             `json:"collected"`
	EffectStickers    []EffectStickers `json:"effectStickers"`
}

type EffectStickers struct {
	Name string `json:"name"`
	ID   string `json:"ID"`
}

type Common struct {
	DocIDStr string `json:"doc_id_str"`
}

type Icon struct {
	URLList []string `json:"urlList"`
}
type Thumbnail struct {
	URLList []string `json:"urlList"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}
type ExtraInfo struct {
	Subtype string `json:"subtype"`
}
type Anchors struct {
	ID          string    `json:"id"`
	Type        int       `json:"type"`
	Keyword     string    `json:"keyword"`
	Icon        Icon      `json:"icon"`
	Schema      string    `json:"schema"`
	LogExtra    string    `json:"logExtra"`
	Description string    `json:"description"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	ExtraInfo   ExtraInfo `json:"extraInfo"`
}

type Words struct {
	Word   string `json:"word"`
	WordID string `json:"word_id"`
}
type VideoSuggestWordsStruct struct {
	Words    []Words `json:"words"`
	Scene    string  `json:"scene"`
	HintText string  `json:"hint_text"`
}

type VideoSuggestWordsList struct {
	VideoSuggestWordsStruct []VideoSuggestWordsStruct `json:"video_suggest_words_struct"`
}

type Data struct {
	Type   int    `json:"type"`
	Item   Item   `json:"item,omitempty"`
	Common Common `json:"common"`
}

type AdInfo struct {
}

type Extra struct {
	Now             int64  `json:"now"`
	Logid           string `json:"logid"`
	FatalItemIds    []any  `json:"fatal_item_ids"`
	SearchRequestID string `json:"search_request_id"`
	APIDebugInfo    any    `json:"api_debug_info"`
}

type LogPb struct {
	ImprID string `json:"impr_id"`
}
type GlobalDoodleConfig struct {
	FeedbackSurvey any `json:"feedback_survey"`
}
