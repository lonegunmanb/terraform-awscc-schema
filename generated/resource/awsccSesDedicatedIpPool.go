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
