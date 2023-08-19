package mailer

import "context"

type Mailer interface {
	Send(ctx context.Context, template EmailTemplate) error
}

type EmailTemplate struct {
	Body          string
	To            string
	From          string
	RecipientName string
	SenderName    string
	Subject       string
}
