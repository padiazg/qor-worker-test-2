package sendemail

import (
	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	"github.com/qor/mailer"
	"github.com/qor/qor/resource"
)

type DataItem struct {
	Key   string
	Value string
	// Type  string // text, string, checkbox, number, float, datetime, select_one, select_many, collection_edit, single_edit
}

type SendEmailArgument struct {
	From        string
	To          string
	Subject     string
	Content     string `sql:"size:65532"`
	Data        *templatedata.TemplateData
	Attachments []mailer.Attachment
}

func (d *SendEmailArgument) ConfigureQorResource(res resource.Resourcer) {
	// if r, ok := res.(*admin.Resource); ok {
	// 	r.UseTheme("grid")
	// }
}

func (d *DataItem) ConfigureQorResource(res resource.Resourcer) {
	// if r, ok := res.(*admin.Resource); ok {
	// r.UseTheme("grid")

	// r.OverrideEditAttrs(func() {
	// 	r.EditAttrs(r.EditAttrs(), "-Key", "-Value")
	// })
	// }
}
