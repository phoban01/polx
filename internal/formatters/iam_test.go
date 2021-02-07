package formatters

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

const policyFixture = `{
  "Version": "2012-10-17",
  "IAMPolicyStatement": [
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

func TestFormatAsIAMPolicy(t *testing.T) {
	t.Helper()
	events := []*cloudtrail.Event{
		&cloudtrail.Event{
			EventName:   aws.String("LookupEvents"),
			EventSource: aws.String("cloudtrail.amazonaws.com"),
		},
	}

	mockIAMPolicy := &IAMPolicy{
		Version: "2012-10-17",
		IAMPolicyStatement: []*IAMPolicyStatement{&IAMPolicyStatement{
			Effect:   "Allow",
			Action:   []string{"cloudtrail:LookupEvents"},
			Resource: []string{"*"},
		}},
	}

	t.Run("Returns an IAM IAMPolicy from a slice of events", func(t *testing.T) {
		got := FormatAsIAMIAMPolicy(events)
		assert.Equal(t, got, mockIAMPolicy)
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

	t.Run("Prints a IAMPolicy as a correctly formatted string", func(t *testing.T) {
		p := FormatAsIAMIAMPolicy(events)
		got, err := p.String()
		assert.NoError(t, err)
		assert.Equal(t, got, policyFixture)
	})

}
