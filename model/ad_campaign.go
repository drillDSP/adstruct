package ad_struct_model

// CampaignDoc
type AdCampaign struct {
	CampaignId   int64  `bson:"campaignId,omitempty" json:"campaignId"`
	PackageName  string `bson:"packageName,omitempty" json:"packageName"`
	AdvertiserId int64  `bson:"advertiserId,omitempty" json:"advertiserId"`
	Name         string `bson:"name,omitempty" json:"name"`
	CampaignType int8   `bson:"campaignType,omitempty" json:"campaignType"`
	Category     uint8  `bson:"category,omitempty" json:"category"`

	Price        float64          `bson:"price,omitempty" json:"price"`
	FrequencyCap int              `bson:"frequencyCap,omitempty" json:"frequencyCap"`
	AdSchedule   map[string][]int `bson:"adSchedule,omitempty" json:"adSchedule"`
	StartTime    int64            `bson:"startTime,omitempty" json:"startTime"`
	EndTime      int64            `bson:"endTime,omitempty" json:"endTime"`
	Status       int8             `bson:"status,omitempty" json:"status"`
	Updated      int64            `bson:"updated,omitempty" json:"updated"`
	CostType     int              `bson:"costType,omitempty" json:"costType"`

	PreviewUrl     string              `bson:"previewUrl,omitempty" json:"previewUrl"` //storeUrl
	LandingPageUrl string              `bson:"landingPageUrl,omitempty" json:"landingPageUrl"`
	Domain         string              `bson:"appDomain,omitempty" json:"appDomain"`
	Developer      string              `bson:"developer,omitempty" json:"developer"`
	IabCategory    map[string][]string `bson:"iabCategory,omitempty" json:"iabCategory"`
	ApkUrl         string              `bson:"apkUrl,omitempty" json:"apkUrl"`
	TrackingUrl    string              `bson:"trackingUrl,omitempty" json:"trackingurl"`
	DeepLink       string              `bson:"deepLink,omitempty" json:"deepLink"`

	CountryCode []string         `bson:"countryCode,omitempty" json:"countryCode"`
	CityCode    map[string][]int `bson:"cityCode,omitempty" json:"citycode"` //城市

	AppName       string              `bson:"appName,omitempty" json:"appName"` // [render]Advertiser
	PublisherId   int64               `bson:"publisherId,omitempty" json:"publisherId"`
	Platform      uint8               `bson:"platform,omitempty" json:"platform"`
	OsVersionMin  int                 `bson:"osVersionMinV2,omitempty" json:"osVersionMinV2"`
	OsVersionMax  int                 `bson:"osVersionMaxV2,omitempty" json:"osVersionMaxV2"`
	NetworkTypeV2 []uint8             `bson:"networkTypeV2,omitempty" json:"networkTypeV2"` //0 all 9 wifi
	MobileCode    []string            `bson:"mobileCode,omitempty" json:"mobileCode"`       //运营商all
	DeviceModelV3 map[string][]string `bson:"deviceModelV3,omitempty" json:"deviceModelV3"`

	TryNew         map[string]TryNew          `bson:"tryNew,omitempty" json:"tryNew"`
	TemplateConf   map[string]map[string]int  `bson:"dspTemplateConf,omitempty" json:"dspTemplateConf"` //模板定向
	TryNewTemplate map[string]*TryNewTemplate `bson:"tryNewTpl,omitempty" json:"tryNewTpl"`

	UpdateTime int64
}

type TryNew struct {
	CampId   int64    `bson:"campaignId,omitempty" json:"campaignId"`
	Status   int64    `bson:"status,omitempty" json:"status"`
	Stime    int64    `bson:"stime,omitempty" json:"stime"`
	Etime    int64    `bson:"etime,omitempty" json:"etime"`
	Rate     float32  `bson:"rate,omitempty" json:"rate"`
	IsM      int64    `bson:"is_m,omitempty" json:"is_m"`
	Area     string   `bson:"area,omitempty" json:"area"`
	IsDsp    int64    `bson:"is_dsp,omitempty" json:"is_dsp"`
	Adx      []string `bson:"adx,omitempty" json:"adx"`
	AdType   []int32  `bson:"adType,omitempty" json:"adType"`
	AppId    []string `bson:"appId,omitempty" json:"appId"`
	TagId    []string `bson:"tagId,omitempty" json:"tagId"`
	BidRaise float32  `bson:"bidRaise,omitempty" json:"bidRaise"`
	MTime    int64    `bson:"mTime,omitempty" json:"mTime"`
	Priority int32    `bson:"priority,omitempty" json:"priority"`
}

// 模版试新配置，多个配置之间试新时间无交集
type (
	TryNewTemplate struct {
		Id   int     `bson:"id,omitempty" json:"id,omitempty"`
		Rate float64 `bson:"rate,omitempty" json:"rate,omitempty"` // 试新比例，百分位。如配置为 1，则表示 1% 比例

		Is_m          int      `bson:"is_m,omitempty" json:"is_m,omitempty"`     // 1:is m 2:non m
		Is_dsp        int      `bson:"is_dsp,omitempty" json:"is_dsp,omitempty"` // 1:is dsp 2: non dsp
		Area          []string `bson:"area,omitempty" json:"area,omitempty"`     // 投放地区
		Adx           []string `bson:"adx,omitempty" json:"adx,omitempty"`
		Is_m_adtype   []int    `bson:"adType,omitempty" json:"adType,omitempty"`       //若isM = 2 ，则为空数组
		Is_dsp_adtype []string `bson:"dspAdType,omitempty" json:"dspAdType,omitempty"` //若isDsp = 2 ，则为空数组

		TplSet []TplTryNewSet `bson:"tplSet,omitempty" json:"tplSet,omitempty"`

		StartTime  int64 `bson:"stime,omitempty" json:"stime,omitempty"`   // 开始试新，时间戳
		EndTime    int64 `bson:"etime,omitempty" json:"etime,omitempty"`   // 结束试新，时间戳
		UpdateTime int64 `bson:"mTime,omitempty" json:"mTime,omitempty"`   // 配置更新，时间戳
		Status     int   `bson:"status,omitempty" json:"status,omitempty"` // 1:active 2:paused
	}

	TplTryNewSet struct {
		GroupId int32   `bson:"groupId,omitempty" json:"groupId,omitempty"` // 模版组合id
		TplIds  []int32 `bson:"tplIds,omitempty" json:"tplIds,omitempty"`   // 模版id
		Rate    float64 `bson:"rate,omitempty" json:"rate,omitempty"`       // 试新比例，百分位。如配置为 1，则表示 1% 比例
	}
)
