package sendemail

type EmailArgument struct {
	From    string
	To      string
	Subject string
	Content string `sql:"size:65532"`
}
