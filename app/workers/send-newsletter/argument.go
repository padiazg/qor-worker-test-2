package sendnewsletter

// Arguments used to run a job
type SendNewsletterArgument struct {
	Subject      string
	Content      string `sql:"size:65532"`
	SendPassword string
}
