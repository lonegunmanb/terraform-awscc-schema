package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreCodeInterpreterCustom = `{
  "block": {
    "attributes": {
      "code_interpreter_arn": {
        "computed": true,
        "description": "The ARN of a CodeInterpreter resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "code_interpreter_id": {
        "computed": true,
        "description": "The id of the code interpreter.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the code interpreter was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the code interpreter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that the code interpreter uses to access resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "Timestamp when the code interpreter was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the code interpreter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_configuration": {
        "description": "Network configuration for code interpreter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "computed": true,
              "description": "Network modes supported by code interpreter",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "Status of code interpreter.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource definition for AWS::BedrockAgentCore::CodeInterpreterCustom",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreCodeInterpreterCustomSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreCodeInterpreterCustom), &result)
	return &result
}
