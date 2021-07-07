package ad_struce_index

import (
	"sync"

	ad_struct_model "thief/pkg_v2/adstruct/model"
)

func SetIndexCampaign(index *IndexCampaign) {
	index_Campaign = index
}

func GetIndexCampaign() *IndexCampaign {
	return index_Campaign
}

var index_Campaign = NewIndexCampaign()

type IndexCampaign struct {
	sync.RWMutex

	lastUpdateTime int64

	m map[int64]ad_struct_model.AdCampaign // key=campaign id, key=country, key=ft type + sub type
}

func NewIndexCampaign() *IndexCampaign {
	return &IndexCampaign{
		m: map[int64]ad_struct_model.AdCampaign{},
	}
}

func (index *IndexCampaign) UpdateOne(camp ad_struct_model.AdCampaign) {
	if index == nil {
		return
	}
	if camp.UpdateTime > 0 && camp.UpdateTime < index.lastUpdateTime {
		return
	}

	index.Lock()
	if camp.Status != 1 {
		delete(index.m, camp.CampaignId)
	}
	index.m[camp.CampaignId] = camp
	index.Unlock()
}

func (index *IndexCampaign) UpdateAll(cr ...ad_struct_model.AdCampaign) {
	if index == nil {
		return
	}
	for i, _ := range cr {
		index.UpdateOne(cr[i])
	}
}

func (index *IndexCampaign) GetOne(campaignID int64) (camp ad_struct_model.AdCampaign) {
	if index == nil || len(index.m) == 0 {
		return
	}

	index.RLock()
	camp = index.m[campaignID]
	index.RUnlock()

	return camp
}
