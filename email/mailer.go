package email

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/aws-sdk-go/aws"
)

type SESMailer struct {
	svc *ses.Client
}

func NewSESMailer(ctx context.Context) (*SESMailer, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	svc := ses.NewFromConfig(cfg)

	return &SESMailer{svc: svc}, nil
}

func (m *SESMailer) SendEmail(
	ctx context.Context, email string, subject string, body string,
) error {
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{email},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data:    aws.String(body),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &types.Content{
				Data:    aws.String(subject),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(email),
	}

	if _, err = m.svc.SendEmail(ctx, input); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
