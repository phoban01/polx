# IAM Policy eXporter

### About

polx is tool that can help in generating least-privilege IAM policies.

```
polx helps generate IAM policies

Usage:
  polx [command]

Available Commands:
  ct          Create IAM Policy from Cloudtrail audit trail
  help        Help about any command
  version     Polx helps generate IAM policies

Flags:
  -h, --help   help for polx

Use "polx [command] --help" for more information about a command.
```

### CloudTrail

**Note that events can take ~15 minutes before they appear in the CloudTrail event log**

The `polx ct` command will generate an AWS IAM profile from the CloudTrail event log. You can specify the lookback window and filter
using a `username` or `AccessKeyId`. The user which `polx` runs as can be specified using the `AWS_PROFILE` environment variable or `--profile` flag. This user should have the appropriate CloudTrail permissions assigned.

```
Credentials can be specified using $AWS_PROFILE environment variable or using the --profile flag

Usage:
  polx ct [flags]

Examples:

	# Build policy for all events in last hour
	polx ct --window 60 --profile aws-admin-profile

	# Build policy for events associated with access key
	polx ct --window 60 --profile aws-admin-profile --access-key-id XXXX-XXXXXX-XXX


Flags:
      --access-key-id string   Filter Events by AccessKeyId
  -h, --help                   help for ct
      --profile string         The profile which will make CloudTrail API calls
      --region string          AWS Region (default "us-east-1")
      --username string        Filter Events by Username
  -w, --window int             How far back in the audit trail to look for events (minutes) (default 30)
```

### Test

`make test`

### Build

`make build`

### Install

`make install`
