package ad_struce_index

import (
	"sync"

	ad_struct_model "thief/pkg_v2/adstruct/model"
)

func SetIndexCreative(index *IndexCreative) {
	index_creative = index
}

func GetIndexCreative() *IndexCreative {
	return index_creative
}

var index_creative = NewIndexCreative()

type IndexCreative struct {
	sync.RWMutex

	lastUpdateTime int64

	m map[int64]map[string]ad_struct_model.AdCreative // key=campaign id, key=country, key=ft type + sub type
}

func NewIndexCreative() *IndexCreative {
	return &IndexCreative{
		m: map[int64]map[string]ad_struct_model.AdCreative{},
	}
}

func (index *IndexCreative) UpdateOne(cr ad_struct_model.AdCreative) {
	if index == nil {
		return
	}
	if cr.UpdateTime > 0 && cr.UpdateTime < index.lastUpdateTime {
		return
	}

	index.Lock()
	if cr.Status != 1 {
		delete(index.m[cr.CampaignID], cr.Country)
	}
	if index.m[cr.CampaignID] == nil {
		index.m[cr.CampaignID] = map[string]ad_struct_model.AdCreative{}
	}
	index.m[cr.CampaignID][cr.Country] = cr
	index.Unlock()
}

func (index *IndexCreative) UpdateAll(cr ...ad_struct_model.AdCreative) {
	if index == nil {
		return
	}
	for i, _ := range cr {
		index.UpdateOne(cr[i])
	}
}

func (index *IndexCreative) GetOne(campaignID int64, country string) (cr ad_struct_model.AdCreative) {
	if index == nil || len(index.m) == 0 {
		return
	}

	index.RLock()
	if len(index.m[campaignID]) > 0 {
		if len(index.m[campaignID][country].CreativeLayer) > 0 {
			cr = index.m[campaignID][country]
		}
	}
	index.RUnlock()

	return cr
}
