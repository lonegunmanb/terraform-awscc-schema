package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2OriginEndpointPolicy = `{
  "block": {
    "attributes": {
      "channel_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_endpoint_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaPackageV2::OriginEndpointPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagev2OriginEndpointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2OriginEndpointPolicy), &result)
	return &result
}
