package templatedata

type TemplateData map[string]interface{}

func (d TemplateData) Set(key string, value interface{}) {
	d[key] = value
}
