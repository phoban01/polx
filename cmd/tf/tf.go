package tf

import (
	"github.com/hashicorp/terraform/configs"
	"github.com/phoban01/polx/internal/formatters"
	"github.com/spf13/afero"
)

var resourceServiceLookup = map[string]string{
	"aws_instance":   "ec2",
	"aws_s3_bucket":  "s3",
	"aws_iam_policy": "iam",
}

type Terraform struct {
	Resources []*configs.Resource
	Policy    formatters.IAMPolicy
}

func Parser(path string) *Terraform {
	appFs := afero.NewOsFs()
	parser := configs.NewParser(appFs)
	values, _ := parser.LoadConfigFile(path)
	return &Terraform{
		Resources: values.ManagedResources,
	}
}

// for _, r := range values.ManagedResources {
//     switch r.ProviderConfigAddr().LocalName {
//     case "aws":
//         fmt.Println(resourceServiceLookup[r.Type])
//     }
// }
