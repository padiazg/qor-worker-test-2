# QOR worker test
A test for QORv3 worker 

## Pre-requisites
Create a `.env` file in the root of this project with content below. I'm using Sendgrid for outgoing emails, but you can use any SMTP server you want. If you don't want to use email, just remove the `EMAIL_*` variables from the `.env` file.
```ini
# database
DATABASE_DIALECT=sqlite3
DATABASE_CONNECTIONSTRING=worker.db
DATABASE_ENABLELOGS=false

# static
STATIC_PATH=./static
STATIC_ROUTE=/static

# email
EMAIL_HOST=smtp.sendgrid.net
EMAIL_PORT=587
EMAIL_USERNAME=apikey
EMAIL_PASSWORD=<your-sendgrid-pass>
```

## How to run
1. Clone this repo
2. Run `go run main.go`
3. Open `http://localhost:9000/admin` in your browser
