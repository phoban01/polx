package ct

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

type CloudTrailOpts struct {
	Region      *string
	AwsProfile  *string
	Start       *time.Time
	End         *time.Time
	AccessKeyID *string
	Username    *string
	Lookback    *int
}

func GetLogsForPeriod(opts *CloudTrailOpts) (events []*cloudtrail.Event, err error) {
	sess := newSession(opts)
	svc := cloudtrail.New(sess)

	input := &cloudtrail.LookupEventsInput{
		StartTime: aws.Time(time.Now()),
		EndTime:   aws.Time(time.Now().Add(time.Minute * (-1 * time.Duration(*opts.Lookback)))),
	}
	if opts.Start != nil {
		input.StartTime = opts.Start
	}
	if opts.End != nil {
		input.EndTime = opts.End
	}
	if opts.AccessKeyID != nil {
		setLookupAttribute(input, aws.String("AccessKeyId"), opts.AccessKeyID)
	}
	if opts.Username != nil {
		setLookupAttribute(input, aws.String("Username"), opts.Username)
	}
	return getEvents(svc, input)
}

func newSession(opts *CloudTrailOpts) *session.Session {
	config := aws.Config{
		Region: opts.Region,
	}
	if opts.AwsProfile != nil {
		config.Credentials = credentials.NewSharedCredentials("", *opts.AwsProfile)
	}
	return session.Must(session.NewSession(&config))
}

func getEvents(ct *cloudtrail.CloudTrail, input *cloudtrail.LookupEventsInput) (events []*cloudtrail.Event, err error) {
	result, err := ct.LookupEvents(input)
	for _, e := range result.Events {
		events = append(events, e)
	}
	for result.NextToken != nil {
		input.SetNextToken(*result.NextToken)
		result, err = ct.LookupEvents(input)
		if err != nil {
			return nil, err
		}
		for _, e := range result.Events {
			events = append(events, e)
		}
	}
	return
}

func setLookupAttribute(input *cloudtrail.LookupEventsInput, key, value *string) {
	input.SetLookupAttributes(append(input.LookupAttributes, &cloudtrail.LookupAttribute{
		AttributeKey:   key,
		AttributeValue: value,
	}))
}
