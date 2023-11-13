package templatedata

type TemplateData map[string]string

func (d TemplateData) Set(key string, value string) {
	d[key] = value
}
