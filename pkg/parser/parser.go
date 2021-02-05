package parser

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

type policy struct {
	Version   string       `json:"Version"`
	Statement []*statement `json:"Statement"`
}

type statement struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource []string `json:"Resource"`
}

func FormatAsIAMPolicy(events []*cloudtrail.Event) ([]byte, error) {
	var allSt []*statement
	for _, e := range events {
		act := fmt.Sprintf("%s:%s", strings.TrimSuffix(*e.EventSource, ".amazonaws.com"), *e.EventName)
		st := &statement{
			Effect: "Allow",
			Action: []string{act},
		}
		var hasResource bool
		for _, r := range e.Resources {
			if strings.HasPrefix(*r.ResourceName, "arn") {
				st.Resource = append(st.Resource, *r.ResourceName)
				hasResource = true
			}
		}
		if hasResource != true {
			st.Resource = []string{"*"}
		}
		allSt = append(allSt, st)
	}
	uniqSt := uniq(allSt)
	group := group(uniqSt)
	p := &policy{
		Version:   "2012-10-17",
		Statement: group,
	}
	return json.MarshalIndent(p, "", "  ")
}

func uniq(statements []*statement) (result []*statement) {
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

func group(statements []*statement) (result []*statement) {
	serviceMap := make(map[string][]string)
	for _, st := range statements {
		if st.Resource[0] != "*" {
			result = append(result, st)
		} else {
			service := strings.Split(st.Action[0], ":")[0]
			serviceMap[service] = append(serviceMap[service], st.Action[0])
		}
	}
	for _, v := range serviceMap {
		result = append(result, &statement{
			Effect:   "Allow",
			Action:   v,
			Resource: []string{"*"},
		})
	}
	return
}
