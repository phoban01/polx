package formatter

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

const policyFixture = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "cloudtrail:LookupEvents"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}`

func TestParser(t *testing.T) {
	t.Helper()
	events := []*cloudtrail.Event{
		&cloudtrail.Event{
			EventName:   aws.String("LookupEvents"),
			EventSource: aws.String("cloudtrail.amazonaws.com"),
		},
	}

	mockPolicy := &Policy{
		Version: "2012-10-17",
		Statement: []*Statement{&Statement{
			Effect:   "Allow",
			Action:   []string{"cloudtrail:LookupEvents"},
			Resource: []string{"*"},
		}},
	}

	t.Run("Returns an IAM Policy from a slice of events", func(t *testing.T) {
		got := FormatAsIAMPolicy(events)
		assert.Equal(t, got, mockPolicy)
	})

}

func TestString(t *testing.T) {
	t.Helper()
	events := []*cloudtrail.Event{
		&cloudtrail.Event{
			EventName:   aws.String("LookupEvents"),
			EventSource: aws.String("cloudtrail.amazonaws.com"),
		},
	}

	t.Run("Prints a Policy as a correctly formatted string", func(t *testing.T) {
		p := FormatAsIAMPolicy(events)
		got, err := p.String()
		assert.NoError(t, err)
		assert.Equal(t, got, policyFixture)
	})

}
