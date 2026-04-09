package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockEnforcedGuardrailConfiguration = `{
  "block": {
    "attributes": {
      "config_id": {
        "computed": true,
        "description": "Unique ID for the account enforced configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the configuration was created",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The ARN of the role used to create the configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_arn": {
        "computed": true,
        "description": "ARN representation for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_id": {
        "computed": true,
        "description": "Unique ID for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_identifier": {
        "computed": true,
        "description": "Identifier for the guardrail, could be the ID or the ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_version": {
        "computed": true,
        "description": "Numerical guardrail version (not DRAFT)",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_enforcement": {
        "computed": true,
        "description": "Model-specific information for the enforced guardrail configuration. If not present, the configuration is enforced on all models",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "excluded_models": {
              "computed": true,
              "description": "Models to exclude from enforcement. If a model is in both lists, it is excluded",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "included_models": {
              "computed": true,
              "description": "Models to enforce the guardrail on",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "owner": {
        "computed": true,
        "description": "Configuration owner type",
        "description_kind": "plain",
        "type": "string"
      },
      "selective_content_guarding": {
        "computed": true,
        "description": "Selective content guarding controls for enforced guardrails",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "messages": {
              "computed": true,
              "description": "Selective guarding mode for user messages",
              "description_kind": "plain",
              "type": "string"
            },
            "system": {
              "computed": true,
              "description": "Selective guarding mode for system prompts",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "Timestamp when the configuration was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_by": {
        "computed": true,
        "description": "The ARN of the role used to update the configuration",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::EnforcedGuardrailConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockEnforcedGuardrailConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockEnforcedGuardrailConfiguration), &result)
	return &result
}
