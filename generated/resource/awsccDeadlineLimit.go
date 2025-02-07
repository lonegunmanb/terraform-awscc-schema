package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineLimit = `{
  "block": {
    "attributes": {
      "amount_requirement_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "current_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "farm_id": {
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
      "limit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Definition of AWS::Deadline::Limit Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineLimit), &result)
	return &result
}
