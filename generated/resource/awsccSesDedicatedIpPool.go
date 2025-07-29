package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesDedicatedIpPool = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_name": {
        "computed": true,
        "description": "The name of the dedicated IP pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_mode": {
        "computed": true,
        "description": "Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the dedicated IP pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SES::DedicatedIpPool",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesDedicatedIpPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesDedicatedIpPool), &result)
	return &result
}
