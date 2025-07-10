package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2ListenerRules = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description": "Set of Resource Identifiers.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Plural Data Source schema for AWS::ElasticLoadBalancingV2::ListenerRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticloadbalancingv2ListenerRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2ListenerRules), &result)
	return &result
}
