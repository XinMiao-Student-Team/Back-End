package models

type MailOptions struct {
	MailUser string
	MailPass string
	MailHost string
	MailPort int
	MailTo   string
	Subject  string
	Body     string
}
