package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsPrivacyBudgetTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_refresh": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "collaboration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "membership_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "membership_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "epsilon": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "users_noise_per_query": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "privacy_budget_template_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "privacy_budget_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Represents a privacy budget within a collaboration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsPrivacyBudgetTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsPrivacyBudgetTemplate), &result)
	return &result
}
