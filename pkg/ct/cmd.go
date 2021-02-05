package ct

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/phoban01/rolex/pkg/parser"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	opts := CloudTrailOpts{
		Start: aws.Time(time.Now().Add(time.Minute * -120)),
		End:   aws.Time(time.Now()),
		// AccessKeyID: aws.String("AKIA4OIXKKYIDGP364BO"),
		// AccessKeyID: aws.String("AKIA4OIXKKYIFPEXEAV7"),
		Username: aws.String("iamadmin"),
	}
	c := &cobra.Command{
		Use:   "policy",
		Short: "Create IAM Policy from Cloudtrail audit trail",
		Run: func(cmd *cobra.Command, args []string) {
			events, err := GetLogsForPeriod(&opts)
			if err != nil {
				panic(err)
			}
			policy, err := parser.FormatAsIAMPolicy(events)
			if err != nil {
				panic(err)
			}
			fmt.Println(policy)
		},
	}

	c.Flags().IntVarP(opts.Lookback, "lookback", "l", 30, "How far back in the audit trail to look for events (minutes)")
	c.Flags().StringVar(opts.AccessKeyID, "access-key-id", "", "Filter Events by Access Key ID")
	c.Flags().StringVar(opts.Username, "username", "", "Filter Events by Username")
	c.Flags().StringVar(opts.AwsProfile, "profile", "", "The profile which will make CloudTrail API calls")

	return c
}
