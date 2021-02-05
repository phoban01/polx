package ct

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
)

type mockCloudTrailClient struct {
	cloudtrailiface.CloudTrailAPI
}

func (m *mockCloudTrailClient) LookupEvents(input *cloudtrail.LookupEventsInput) (resp *cloudtrail.LookupEventsOutput, err error) {
	resp = &cloudtrail.LookupEventsOutput{
		Events: []*cloudtrail.Event{},
	}
	return
}

func TestGetLogsForPeriod(t *testing.T) {
	t.Helper()

	client := &mockCloudTrailClient{}

	t.Run("It returns a slice of events", func(t *testing.T) {
		want := []*cloudtrail.Event{}
		got, err := GetLogsForPeriod(client, &CloudTrailOpts{})
		assert.IsType(t, want, got)
		assert.NoError(t, err)
	})
}
