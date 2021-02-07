package formatters

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

type IAMPolicy struct {
	Version            string                `json:"Version"`
	IAMPolicyStatement []*IAMPolicyStatement `json:"IAMPolicyStatement"`
}

type IAMPolicyStatement struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource []string `json:"Resource"`
}

func FormatAsIAMPolicy(events []*cloudtrail.Event) *IAMPolicy {
	var allSt []*IAMPolicyStatement
	for _, e := range events {
		var resources []string
		for _, r := range e.Resources {
			if strings.HasPrefix(*r.ResourceName, "arn") {
				resources = append(resources, *r.ResourceName)
			}
		}
		if len(resources) == 0 {
			resources = []string{"*"}
		}
		action := fmt.Sprintf("%s:%s", strings.TrimSuffix(*e.EventSource, ".amazonaws.com"), *e.EventName)
		allSt = append(allSt, &IAMPolicyStatement{
			Effect:   "Allow",
			Action:   []string{action},
			Resource: resources,
		})
	}
	uniqSt := uniq(allSt)
	group := group(uniqSt)
	return &IAMPolicy{
		Version:            "2012-10-17",
		IAMPolicyStatement: group,
	}
}

func (p *IAMPolicy) String() (string, error) {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func uniq(statements []*IAMPolicyStatement) (result []*IAMPolicyStatement) {
	hashList := make(map[string]struct{})
	for _, st := range statements {
		hash := sha256.New()
		hash.Write([]byte(fmt.Sprintf("%v", *st)))
		hKey := fmt.Sprintf("%x", hash.Sum(nil))
		if _, ok := hashList[hKey]; !ok {
			hashList[hKey] = struct{}{}
			result = append(result, st)
		}
	}
	return
}

func group(statements []*IAMPolicyStatement) (result []*IAMPolicyStatement) {
	serviceMap := make(map[string][]string)
	for _, st := range statements {
		if st.Resource[0] != "*" {
			result = append(result, st)
		} else {
			service := strings.Split(st.Action[0], ":")[0]
			serviceMap[service] = append(serviceMap[service], st.Action[0])
		}
	}
	for _, act := range serviceMap {
		result = append(result, &IAMPolicyStatement{
			Effect:   "Allow",
			Action:   act,
			Resource: []string{"*"},
		})
	}
	return
}
