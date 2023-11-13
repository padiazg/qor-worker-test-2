package sendnewsletter

import (
	"fmt"
	"time"

	"github.com/padiazg/qor-worker-test/config/application"
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
	w.Worker.RegisterJob(&worker.Job{
		Name: "SendNewsletter",
		Handler: func(argument interface{}, job worker.QorJobInterface) error {
			// `AddLog` add job log
			job.AddLog("Started sending newsletters...")
			job.AddLog(fmt.Sprintf("Argument: %+v", argument.(*SendNewsletterArgument)))

			for i := 1; i <= 3; i++ {
				time.Sleep(100 * time.Millisecond)
				job.AddLog(fmt.Sprintf("Sending newsletter %v...", i))
				// `SetProgress` set job progress percent, from 0 - 100
				job.SetProgress(uint(i))
			}

			job.AddLog("Finished send newsletters")
			return nil
		},
		// Arguments used to run a job
		Resource: application.Admin.NewResource(&SendNewsletterArgument{}),
	})

	if w.Config.AdminConfig == nil {
		w.Config.AdminConfig = &admin.Config{}
	}

	// add worker to admin dashboard
	application.Admin.AddResource(w.Worker, w.Config.AdminConfig)

	// add worker to applications workers map
	application.SetWorker("SendNewsletter", w)
}

// SendNewsLetter add a job to worker
func (wrk *Worker) SendNewsletter(argument *SendNewsletterArgument, ctx *admin.Context) error {
	var (
		w      = wrk.Worker
		job    = w.GetRegisteredJob("SendNewsletter")
		res    = w.JobResource
		newJob = res.NewStruct().(worker.QorJobInterface)
	)

	// apply argument to job
	newJob.SetSerializableArgumentValue(argument)

	// set job to be run
	newJob.SetJob(job)

	// save the job
	if err := res.CallSave(newJob, ctx.Context); err != nil {
		return fmt.Errorf("error saving job: %v", err)
	}

	// add job to worker
	if err := w.AddJob(newJob); err != nil {
		return fmt.Errorf("error adding job to worker: %v", err)
	}

	return nil
}
