package ct

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionOpts struct {
	Region     string
	AwsProfile string
}

func NewSession(opts *SessionOpts) *session.Session {
	config := aws.Config{
		Region: aws.String(opts.Region),
	}
	if opts.AwsProfile != "" {
		config.Credentials = credentials.NewSharedCredentials("", opts.AwsProfile)
	}
	return session.Must(session.NewSession(&config))
}
