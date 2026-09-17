package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53TrafficPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the traffic policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "computed": true,
        "description": "Any comments to include about the traffic policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "document": {
        "computed": true,
        "description": "The definition of the traffic policy in JSON format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the traffic policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_policy_id": {
        "computed": true,
        "description": "The ID that Amazon Route 53 assigned to the traffic policy when it was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The DNS type of the resource record sets that Amazon Route 53 creates when the traffic policy is used to create a traffic policy instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number that Amazon Route 53 assigned to the traffic policy.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Route53::TrafficPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53TrafficPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53TrafficPolicy), &result)
	return &result
}
