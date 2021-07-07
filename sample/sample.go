package sample

import _ "embed"

//go:embed DSP.iv.video+动态通用endcard_191_template.json
var sample []byte

func NewTemplateSample() []byte {
	return sample
}
