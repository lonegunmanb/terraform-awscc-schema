package data

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
        "computed": true,
        "description": "The ID for the custom action target. Can contain up to 20 alphanumeric characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for the custom action target.",
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
        "description": "The name of the custom action target. Can contain up to 20 characters.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::ActionTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubActionTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubActionTarget), &result)
	return &result
}
