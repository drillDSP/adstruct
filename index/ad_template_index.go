package ad_struce_index

import (
	"sync"

	ad_struct_model "thief/pkg_v2/adstruct/model"
)

func SetIndexTemplate(index *IndexTemplate) {
	index_template = index
}

func GetIndexTemplate() *IndexTemplate {
	return index_template
}

var index_template = NewIndexTemplate()

type IndexTemplate struct {
	sync.RWMutex

	lastUpdateTime int64

	m map[string]ad_struct_model.AdTemplate // key=template id
}

func NewIndexTemplate() *IndexTemplate {
	return &IndexTemplate{
		m: map[string]ad_struct_model.AdTemplate{},
	}
}

func (index *IndexTemplate) UpdateOne(temp ad_struct_model.AdTemplate) {
	if index == nil {
		return
	}
	if temp.UpdateTime > 0 && temp.UpdateTime < index.lastUpdateTime {
		return
	}

	index.Lock()
	if temp.Status != 1 {
		delete(index.m, temp.ID)
	}
	index.m[temp.ID] = temp
	index.Unlock()
}

func (index *IndexTemplate) UpdateAll(temp ...ad_struct_model.AdTemplate) {
	if index == nil {
		return
	}
	for i, _ := range temp {
		index.UpdateOne(temp[i])
	}
}

func (index *IndexTemplate) GetOne(id string) (temp ad_struct_model.AdTemplate) {
	if index == nil || len(index.m) == 0 {
		return
	}

	index.RLock()
	temp = index.m[id]
	index.RUnlock()

	return temp
}

func (index *IndexTemplate) GetAll() (temp []ad_struct_model.AdTemplate) {
	if index == nil || len(index.m) == 0 {
		return
	}

	index.RLock()
	temp = make([]ad_struct_model.AdTemplate, len(index.m))
	i := 0
	for _, v := range index.m {
		temp[i] = v
		i++
	}
	index.RUnlock()

	return temp
}
