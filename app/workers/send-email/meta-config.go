package sendemail

import (
	"fmt"

	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
)

type SendEmailDataMetaConfig struct {
	Fields []DataItem
}

func (d *SendEmailDataMetaConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*admin.Meta); ok {
		var (
			baseResource = meta.GetBaseResource().(*admin.Resource)
			res          = baseResource.GetAdmin().NewResource(DataItem{}, &admin.Config{Name: meta.GetName()})
		)

		for _, f := range d.Fields {
			// if f.Type == "" {
			// 	f.Type = "string"
			// }
			res.Meta(&admin.Meta{
				Name:      f.Key,
				FieldName: f.Key,
				// Type:      f.Type,
				Type: "string",
				Valuer: func(r interface{}, context *qor.Context) interface{} {
					return r.(*DataItem).Value
				},
			})
			res.EditAttrs(res.EditAttrs(), f.Key)
		}

		res.OverrideEditAttrs(func() { res.EditAttrs(res.EditAttrs(), "-Key", "-Value") })
		// res.OverrideEditAttrs(func() { res.EditAttrs(res.EditAttrs(), "-type") })

		meta.Type = "template_data"
		meta.Resource = res

		// meta.FormattedValuer = func(record interface{}, context *qor.Context) (result interface{}) {
		// 	return d.Data
		// }

		// meta.Valuer = func(record interface{}, context *qor.Context) (result interface{}) {
		// 	return d.Data
		// }

		meta.SetSetter(func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			r := record.(*SendEmailArgument)
			if r.Data == nil {
				r.Data = &templatedata.TemplateData{}
			}
			for _, v := range metaValue.MetaValues.Values {
				v0 := v.Value.([]string)[0]
				r.Data.Set(v.Name, v0)
			}
		})
	}
}

func ToString(value interface{}) string {
	if v, ok := value.([]string); ok {
		for _, s := range v {
			if s != "" {
				return s
			}
		}
		return ""
	} else if v, ok := value.(string); ok {
		return v
	} else if v, ok := value.([]interface{}); ok {
		for _, s := range v {
			if fmt.Sprint(s) != "" {
				return fmt.Sprint(s)
			}
		}
		return ""
	}
	return fmt.Sprintf("%v", value)
}
