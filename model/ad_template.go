package ad_struct_model

import "math/rand"

// TemplateDoc
type AdTemplate struct {
	ID           string   `json:"id"`            // 模版逻辑id
	HtmlID       string   `json:"html_id"`       // 模版 html代码id关联
	FeatureID    string   `json:"feature_id"`    // 模版 特征id关联
	ParentID     []string `json:"parent_id"`     // 父 模版逻辑id。若存在则优先召回
	TemplateDesc string   `json:"template_desc"` // 模版 report id关联

	APIFramework int   `json:"api_framework"` // 模版 render 所需框架
	ClickArea    []int `json:"click_area"`    // 模版 特征关联
	TemplateType int   `json:"template_type"` // 模版 特征关联

	ElementLayer map[string][]ElementLayer `json:"element_layer"` // 模版表现所需全部素材关联。key=element layer，非空有效。同层 element 召回任一

	FlowDirection TemplateFlowDirection `json:"flow_direction"` // 模版定向流量相关配置

	UpdateTime int64 `json:"update_time"` // 更新时间戳
	Status     int   `json:"status"`      // 模版启用状态
}

type ElementLayer struct {
	Name              string              `json:"name"`              // 模版中 element id
	Priority          int                 `json:"priority"`          // 模版中 element 组织同层级之间描述优先级，数字小的优先级高
	FrontPlaceholder  string              `json:"front_placeholder"` // element html代码 占位符关联
	Optional          bool                `json:"optional"`          // element 可选标记，true：非必要召回项
	ElementRule       []ElementRule       `json:"element_rule"`      // element 定向规则。key=element name
	ElementResolution []ElementResolution `json:"element_resolution,omitempty"`
}

type TemplateFlowDirection struct {
	Adx          []string                 `json:"adx"`          // 定向渠道
	Adtype       []string                 `json:"adtype"`       // 定向渠道
	Businesstype []string                 `json:"businesstype"` // 定向业务线
	FilterRule   map[string]AllFilterRule `json:"filter_rule"`  // 定向规则。key=【default\adxname】，value1=多组定向规则，任一满足即可召回，value2=多组定向规则组合，满足全部定向才可召回
}

type ElementRule struct {
	FtType     string                   `json:"ft_type"`     // element index
	SubType    string                   `json:"sub_type"`    // element index
	FilterRule map[string]AllFilterRule `json:"filter_rule"` // 定向规则。key=【default\adxname】，value1=多组定向规则，任一满足即可召回，value2=多组定向规则组合，满足全部定向才可召回
}

type ElementResolution struct {
	Operator   string              `json:"operator"`
	Resolution map[string][]string `json:"resolution"`
}

type AllFilterRule [][]FilterRule

type FilterRule struct {
	Attribute     string   `json:"attribute"`      // 定向字段
	AttributeType string   `json:"attribute_type"` // 【req\cfg\cr】
	Operator      string   `json:"operator"`       // 定向逻辑
	Values        []string `json:"values"`         // 定向字段需要满足的值
	ValuesType    string   `json:"values_type"`    // 【req\cfg\cr】
}

func (temp *AdTemplate) GetTemplateParentId() (parentid []string) {
	if temp == nil {
		return nil
	}
	return temp.ParentID
}

func (temp *AdTemplate) GetTemplateAllElementPlaceholder() (placeholder map[string]string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	for _, elements := range temp.ElementLayer {
		for _, elem := range elements {
			if len(elem.FrontPlaceholder) == 0 {
				continue
			}
			if placeholder == nil {
				placeholder = map[string]string{}
			}
			placeholder[elem.Name] = elem.FrontPlaceholder
		}
	}
	return placeholder
}

func (temp *AdTemplate) GetTemplateElementRule(element string) []ElementRule {
	if temp == nil || len(element) == 0 {
		return nil
	}
	for _, elements := range temp.ElementLayer {
		for _, elem := range elements {
			if elem.Name == element {
				return elem.ElementRule
			}
		}
	}
	return nil
}

func (temp *AdTemplate) GetTemplateAllElementList() (elementlist []string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	elementlist = append(elementlist, temp.GetTemplateDefaultElementList()...)
	elementlist = append(elementlist, temp.GetTemplateOptionalElementList()...)
	for _, elems := range temp.GetTemplatePriorityElementList() {
		elementlist = append(elementlist, elems...)
	}
	for _, elems := range temp.GetTemplateRandomElementList() {
		elementlist = append(elementlist, elems...)
	}
	return elementlist
}

func (temp *AdTemplate) GetTemplateDefaultElementList() (elementlist []string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	for _, elem := range temp.ElementLayer["default"] {
		elementlist = append(elementlist, elem.Name)
	}
	return elementlist
}

func (temp *AdTemplate) GetTemplateOptionalElementList() (elementlist []string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	for _, elem := range temp.ElementLayer["optional"] {
		elementlist = append(elementlist, elem.Name)
	}
	return elementlist
}

func (temp *AdTemplate) GetTemplatePriorityElementList() (elementlist map[string][]string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	for layer, elements := range temp.ElementLayer {
		if layer == "default" {
			continue
		}
		if elementlist == nil {
			elementlist = map[string][]string{}
		}
		priority := map[int]string{} // sort by priority
		for _, elem := range elements {
			if elem.Priority == 0 {
				continue
			}
			priority[elem.Priority] = elem.Name
		}
		for i := 1; ; i++ {
			if len(priority) == len(elementlist[layer]) {
				break
			}
			if _, ok := priority[i]; !ok {
				continue
			}
			elementlist[layer] = append(elementlist[layer], priority[i])
		}
	}
	return elementlist
}

func (temp *AdTemplate) GetTemplateRandomElementList() (elementlist map[string][]string) {
	if temp == nil || len(temp.ElementLayer) == 0 {
		return nil
	}
	for layer, elements := range temp.ElementLayer {
		if layer == "default" {
			continue
		}
		for _, elem := range elements {
			if elem.Priority != 0 {
				continue
			}
			if elementlist == nil {
				elementlist = map[string][]string{}
			}
			elementlist[layer] = append(elementlist[layer], elem.Name)
		}
		rand.Shuffle(len(elementlist[layer]), func(i, j int) {
			elementlist[layer][i], elementlist[layer][j] = elementlist[layer][j], elementlist[layer][i]
		})
	}
	return elementlist
}
