package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomAiGuardrailVersion = `{
  "block": {
    "attributes": {
      "ai_guardrail_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ai_guardrail_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ai_guardrail_version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_id": {
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
      "modified_time_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Wisdom::AIGuardrailVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWisdomAiGuardrailVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomAiGuardrailVersion), &result)
	return &result
}
