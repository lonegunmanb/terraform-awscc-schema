package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubActionTarget = `{
  "block": {
    "attributes": {
      "action_target_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom action target.",
        "description_kind": "plain",
        "type": "string"
      },
      "action_target_id": {
        "description": "The ID for the custom action target. Can contain up to 20 alphanumeric characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description for the custom action target.",
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
      "name": {
        "description": "The name of the custom action target. Can contain up to 20 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Creates a custom action target in Security Hub.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubActionTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubActionTarget), &result)
	return &result
}
