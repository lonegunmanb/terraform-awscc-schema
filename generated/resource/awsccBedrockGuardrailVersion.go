package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockGuardrailVersion = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the Guardrail version",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "guardrail_arn": {
        "computed": true,
        "description": "Arn representation for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_id": {
        "computed": true,
        "description": "Unique id for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_identifier": {
        "description": "Identifier (GuardrailId or GuardrailArn) for the guardrail",
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
      "version": {
        "computed": true,
        "description": "Guardrail version",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::GuardrailVersion Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockGuardrailVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockGuardrailVersion), &result)
	return &result
}
