package ct

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
)

type CloudTrailOpts struct {
	Start       *time.Time
	End         *time.Time
	AccessKeyID string
	Username    string
	Window      int
}

func GetLogsForPeriod(client cloudtrailiface.CloudTrailAPI, opts *CloudTrailOpts) (events []*cloudtrail.Event, err error) {
	input := &cloudtrail.LookupEventsInput{
		StartTime: aws.Time(time.Now().Add(time.Minute * (-1 * time.Duration(opts.Window)))),
		EndTime:   aws.Time(time.Now()),
	}
	if opts.Start != nil {
		input.StartTime = opts.Start
	}
	if opts.End != nil {
		input.EndTime = opts.End
	}
	if opts.AccessKeyID != "" {
		setLookupAttribute(input, "AccessKeyId", opts.AccessKeyID)
	}
	if opts.Username != "" {
		setLookupAttribute(input, "Username", opts.Username)
	}
	return getEvents(client, input)
}

func getEvents(ct cloudtrailiface.CloudTrailAPI, input *cloudtrail.LookupEventsInput) (events []*cloudtrail.Event, err error) {
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

func setLookupAttribute(input *cloudtrail.LookupEventsInput, key, value string) {
	input.SetLookupAttributes(append(input.LookupAttributes, &cloudtrail.LookupAttribute{
		AttributeKey:   aws.String(key),
		AttributeValue: aws.String(value),
	}))
}
