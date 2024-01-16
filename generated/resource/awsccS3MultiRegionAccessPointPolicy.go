package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3MultiRegionAccessPointPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mrap_name": {
        "description": "The name of the Multi Region Access Point to apply policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "description": "Policy document to apply to a Multi Region Access Point",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_status": {
        "computed": true,
        "description": "The Policy Status associated with this Multi Region Access Point",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "is_public": {
              "computed": true,
              "description": "Specifies whether the policy is public or not.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "The policy to be attached to a Multi Region Access Point",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3MultiRegionAccessPointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3MultiRegionAccessPointPolicy), &result)
	return &result
}
