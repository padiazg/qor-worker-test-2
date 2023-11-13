package user

import (
	wrk_se "github.com/padiazg/qor-worker-test/app/workers/send-email"
	wrk_sn "github.com/padiazg/qor-worker-test/app/workers/send-newsletter"
	"github.com/padiazg/qor-worker-test/config/application"
	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	um "github.com/padiazg/qor-worker-test/model/user"
	"github.com/qor/admin"
	"github.com/qor/worker"
)

// Config home config struct
type Config struct{}

// App home app
type App struct {
	Config *Config
}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	application.Admin.AddMenu(&admin.Menu{Name: "Users", Priority: 1})

	user := application.Admin.AddResource(&um.User{}, userConfig)

	// user.Action(&admin.Action{
	// 	Name: "List jobs for SendNewsletter",
	// 	Handler: func(argument *admin.ActionArgument) error {
	// 		w := application.Workers["SendNewsletter"].(*worker.Worker)
	// 		c := argument.Context.NewResourceContext(w.JobResource)

	// 		result, err := c.FindMany()
	// 		if err != nil {
	// 			return err
	// 		}

	// 		r := result.(*[]*worker.QorJob)

	// 		for _, r0 := range *r {
	// 			fmt.Printf("result: %+v\n", r0)
	// 		}

	// 		return nil
	// 	},
	// 	Modes: []string{"menu_item", "edit"},
	// })

	user.Action(&admin.Action{
		Name: "Send Newsletter A",
		Handler: func(argument *admin.ActionArgument) error {
			var (
				w      = application.Workers["SendNewsletter"].(*wrk_sn.Worker).Worker
				job    = w.GetRegisteredJob("SendNewsletter")
				res    = w.JobResource
				ctx    = argument.Context.Context
				newJob = res.NewStruct().(worker.QorJobInterface)
			)

			newJob.SetSerializableArgumentValue(&wrk_sn.SendNewsletterArgument{
				Subject:      "Hello",
				Content:      "World",
				SendPassword: "123456",
			})

			newJob.SetJob(job)
			res.CallSave(newJob, ctx)
			w.AddJob(newJob)

			return nil
		},
		Modes: []string{"menu_item", "edit"},
	})

	user.Action(&admin.Action{
		Name: "Send Newsletter B",
		Handler: func(argument *admin.ActionArgument) error {
			var w = application.Workers["SendNewsletter"].(*wrk_sn.Worker)

			w.SendNewsletter(&wrk_sn.SendNewsletterArgument{
				Subject:      "Hello",
				Content:      "World",
				SendPassword: "123456",
			}, argument.Context)

			return nil
		},
		Modes: []string{"menu_item", "edit"},
	})

	user.Action(&admin.Action{
		Name: "Send email (sync)",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {
				var (
					user          = record.(*um.User)
					emailSettings = application.Config.Settings.Email
				)

				if err := emailSettings.SendEmail(user.Email, "Hello", &templatedata.TemplateData{
					"Name": user.First,
				}); err != nil {
					argument.Context.AddError(err)
					return err
				}
			}

			return nil
		},
		Modes: []string{"menu_item", "edit"},
	})

	user.Action(&admin.Action{
		Name: "Send email (async)",
		Handler: func(argument *admin.ActionArgument) error {
			var (
				w      = application.Workers["SendEmail"].(*wrk_se.Worker).Worker
				job    = w.GetRegisteredJob("SendEmail")
				res    = w.JobResource
				ctx    = argument.Context.Context
				newJob = res.NewStruct().(worker.QorJobInterface)
			)

			for _, record := range argument.FindSelectedRecords() {
				var (
					user = record.(*um.User)
				)
				newJob.SetSerializableArgumentValue(&wrk_se.SendEmailArgument{
					To:      user.Email,
					Subject: "Hello from Action",
					Data: &templatedata.TemplateData{
						"First": user.First,
						"Last":  user.Last,
					},
				})

				newJob.SetJob(job)
				res.CallSave(newJob, ctx)
				w.AddJob(newJob)

				// if err := emailSettings.SendEmail(user.Email, "Hello", &templatedata.TemplateData{
				// 	"Name": user.First,
				// }); err != nil {
				// 	argument.Context.AddError(err)
				// 	return err
				// }

			}

			return nil
		},
		Modes: []string{"menu_item", "edit"},
	})

}
