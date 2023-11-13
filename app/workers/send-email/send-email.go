package sendemail

import (
	"fmt"
	"log"

	"github.com/padiazg/qor-worker-test/config/application"
	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	"github.com/qor/admin"
	"github.com/qor/worker"
)

// Config home config struct
type Config struct {
	AdminConfig *admin.Config
}

// App home app
type Worker struct {
	Config *Config
	Worker *worker.Worker
}

// New new home app
func New(config *Config) *Worker {
	return &Worker{Config: config}
}

// ConfigureApplication configure application
func (w *Worker) ConfigureApplication(application *application.Application) {
	w.Worker = worker.New()

	r := application.Admin.NewResource(&SendEmailArgument{})
	r.Meta(&admin.Meta{
		Name:      "Data",
		FieldName: "Data",
		// Type:      "template_data",
		Config: &SendEmailDataMetaConfig{
			Fields: []DataItem{
				{
					Key:   "First",
					Value: "",
				},
				{
					Key:   "Last",
					Value: "",
				},
				{
					Key:   "Time",
					Value: "",
					// Type:  "datetime",
				},
			},
		},
	})

	w.Worker.RegisterJob(&worker.Job{
		Name: "SendEmail",
		Handler: func(argument interface{}, job worker.QorJobInterface) error {
			var (
				arg           = argument.(*SendEmailArgument)
				to            = arg.To
				subject       = arg.Subject
				content       = arg.Content
				data          = arg.Data
				emailSettings = application.Config.Settings.Email
			)

			job.AddLog("Started sending email")
			job.AddLog(fmt.Sprintf("To: %s", to))
			job.AddLog(fmt.Sprintf("Subject: %s", subject))
			job.AddLog(fmt.Sprintf("Content: %s", content))
			job.AddLog(fmt.Sprintf("argument: %+v", argument))
			job.AddLog(fmt.Sprintf("data: %+v", data))

			if data == nil {
				data = &templatedata.TemplateData{}
			}

			if err := emailSettings.SendEmail(to, subject, data); err != nil {
				job.AddLog(fmt.Sprintf("Error: %v", err))
				return fmt.Errorf("error sending email: %v", err)
			}

			job.SetProgress(100)
			return nil
		},
		// Arguments used to run a job
		Resource: r,
	})

	if w.Config.AdminConfig == nil {
		w.Config.AdminConfig = &admin.Config{}
	}

	// add worker to admin dashboard
	application.Admin.AddResource(w.Worker, w.Config.AdminConfig)

	// add worker to applications workers map
	application.SetWorker("SendEmail", w)
}

// SendEmail send email
func (wrk *Worker) SendEmail(argument *SendEmailArgument, ctx *admin.Context) error {
	var (
		w      = wrk.Worker
		job    = w.GetRegisteredJob("SendEmail")
		res    = w.JobResource
		newJob = res.NewStruct().(worker.QorJobInterface)
	)

	log.Printf("argument: %+v", argument)

	// apply argument to job
	newJob.SetSerializableArgumentValue(argument)

	// set job to be run
	newJob.SetJob(job)

	// save the job
	if err := res.CallSave(newJob, ctx.Context); err != nil {
		return fmt.Errorf("error saving job: %v", err)
	}

	// // add job to worker
	if err := w.AddJob(newJob); err != nil {
		return fmt.Errorf("error adding job to worker: %v", err)
	}

	return nil
}
