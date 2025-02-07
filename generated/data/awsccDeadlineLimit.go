package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineLimit = `{
  "block": {
    "attributes": {
      "amount_requirement_name": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "farm_id": {
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
      "limit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Deadline::Limit",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineLimit), &result)
	return &result
}
