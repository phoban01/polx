package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSSessionOpts struct {
	Region     string
	AwsProfile string
}

func NewAWSSession(opts *AWSSessionOpts) *session.Session {
	config := aws.Config{
		Region: aws.String(opts.Region),
	}
	if opts.AwsProfile != "" {
		config.Credentials = credentials.NewSharedCredentials("", opts.AwsProfile)
	}
	return session.Must(session.NewSession(&config))
}
