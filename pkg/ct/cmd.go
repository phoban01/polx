package ct

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/phoban01/polx/pkg/formatter"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	sessOpts := new(SessionOpts)
	opts := new(CloudTrailOpts)
	c := &cobra.Command{
		Use:   "ct",
		Short: "Create IAM Policy from Cloudtrail audit trail",
		Long:  "Credentials can be specified using $AWS_PROFILE environment variable or using the --profile flag",
		Example: `
	# Build policy for all events in last hour
	polx ct --window 60 --profile aws-admin-profile

	# Build policy for events associated with access key
	polx ct --window 60 --profile aws-admin-profile --access-key-id XXXX-XXXXXX-XXX
`,
		Run: func(cmd *cobra.Command, args []string) {
			sess := NewSession(sessOpts)
			client := cloudtrail.New(sess)
			events, err := GetLogsForPeriod(client, opts)
			if err != nil {
				fmt.Printf("ERROR: %s", err)
				os.Exit(1)
			}
			if len(events) == 0 {
				fmt.Printf("Warning: No events found for time period\n")
				os.Exit(0)
			}
			policy := formatter.FormatAsIAMPolicy(events)
			response, err := policy.String()
			if err != nil {
				fmt.Printf("ERROR: %s", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", response)
		},
	}

	c.Flags().StringVar(&sessOpts.AwsProfile, "profile", "", "The profile which will make CloudTrail API calls")
	c.Flags().StringVar(&sessOpts.Region, "region", "us-east-1", "AWS Region")
	c.Flags().StringVar(&opts.AccessKeyID, "access-key-id", "", "Filter Events by AccessKeyId")
	c.Flags().IntVarP(&opts.Window, "window", "w", 30, "How far back in the audit log to look for events (minutes)")
	c.Flags().StringVar(&opts.Username, "username", "", "Filter Events by Username")

	return c
}
