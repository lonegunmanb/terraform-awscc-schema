package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2OriginEndpointPolicy = `{
  "block": {
    "attributes": {
      "channel_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "channel_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin_endpoint_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "\u003cp\u003eRepresents a resource policy that allows or denies access to an origin endpoint.\u003c/p\u003e",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediapackagev2OriginEndpointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2OriginEndpointPolicy), &result)
	return &result
}
