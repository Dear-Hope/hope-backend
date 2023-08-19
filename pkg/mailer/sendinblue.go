package mailer

import (
	"HOPE-backend/config"
	"fmt"
	"github.com/sendinblue/APIv3-go-library/lib"
	"golang.org/x/net/context"
)

type sendInBlue struct {
	sib *lib.APIClient
}

func New(cfg config.MailerConfig) Mailer {
	sibConfig := lib.NewConfiguration()
	sibConfig.AddDefaultHeader("api-key", cfg.SendInBlue.ApiKey)
	sibConfig.AddDefaultHeader("partner-key", cfg.SendInBlue.PartnerKey)

	return &sendInBlue{sib: lib.NewAPIClient(sibConfig)}
}

func (s sendInBlue) Send(ctx context.Context, template EmailTemplate) error {
	_, _, err := s.sib.TransactionalEmailsApi.SendTransacEmail(
		context.Background(),
		lib.SendSmtpEmail{
			Sender: &lib.SendSmtpEmailSender{
				Name:  template.SenderName,
				Email: template.From,
			},
			To: []lib.SendSmtpEmailTo{
				{
					Email: template.To,
				},
			},
			Subject:     template.Subject,
			HtmlContent: template.Body,
		},
	)
	if err != nil {
		return fmt.Errorf("[Mailer.Send][990001] Failed to send email: %v", err)
	}

	return nil
}
