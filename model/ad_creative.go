package ad_struct_model

type CreativeDoc struct {
	Content map[string]map[string]map[string]interface{} `bson:"content,omitempty" json:"content"` //creative_type、creative_id、字段名、字段值

	Ext interface{} // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
}

type AdCreative struct {
	DocId       string `bson:"_id,omitempty" json:"_id"`
	CampaignID  int64  `bson:"campaignId,omitempty" json:"campaignId"`
	PackageName string `bson:"packageName,omitempty" json:"packageName"`
	Country     string `bson:"countryCode,omitempty" json:"countryCode"`

	CreativeLayer map[string]map[string][]CreativeLayer // key=ftype__subtype, key=creative resolution

	UpdateTime int64 `json:"update_time"`                    // 更新时间戳
	Status     int   `bson:"status,omitempty" json:"status"` // 模版启用状态
}

type CreativeLayer struct {
	CreativeDocId string
	CreativeKey   int64  `protobuf:"varint,1,opt,name=creative_key,json=creativeKey,proto3" json:"creative_key,omitempty"`
	CreativeId    int64  `protobuf:"varint,2,opt,name=creative_id,json=creativeId,proto3" json:"creative_id,omitempty"`
	AdvCreativeId int64  `protobuf:"varint,3,opt,name=adv_creative_id,json=advCreativeId,proto3" json:"adv_creative_id,omitempty"`
	UniqCid       int64  `protobuf:"varint,4,opt,name=uniq_cid,json=uniqCid,proto3" json:"uniq_cid,omitempty"`
	VideoLength   string `protobuf:"bytes,5,opt,name=video_length,json=videoLength,proto3" json:"video_length,omitempty"`
	VideoSize     string `protobuf:"bytes,6,opt,name=video_size,json=videoSize,proto3" json:"video_size,omitempty"`
	Bitrate       int32  `protobuf:"varint,7,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	Width         int32  `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32  `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	CreateTime    int64  `protobuf:"varint,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    int64  `protobuf:"varint,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Fmd5          string `protobuf:"bytes,12,opt,name=fmd5,proto3" json:"fmd5,omitempty"`

	CreativeType string
	FtType       string
	SubType      string

	Resolution string

	Platform  int32  `protobuf:"varint,13,opt,name=platform,proto3" json:"platform,omitempty"`
	MinOs     int64  `protobuf:"varint,14,opt,name=minOs,proto3" json:"minOs,omitempty"`
	Mime      string `protobuf:"bytes,15,opt,name=mime,proto3" json:"mime,omitempty"`
	Attribute string `protobuf:"bytes,16,opt,name=attribute,proto3" json:"attribute,omitempty"`
}
